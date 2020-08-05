package store

import (
	"context"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
)

type CloudStorageClient interface {
	ExecSignedURL(object string, expire time.Time) (url string, err error)
	ExecUploadObject(ctx context.Context, object string, reader io.Reader) error
}

type cloudStorageClient struct {
	bucket           string
	signedURLFunc    SignedURLFunc
	uploadObjectFunc UploadObjectFunc
}

func NewCloudStorageClient(ctx context.Context, bucket string) (CloudStorageClient, error) {
	var credentialJSON []byte
	{
		credential, err := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
		if err != nil {
			return nil, xerrors.Errorf("failed to google.FindDefaultCredentials: %w", err)
		}
		if credential == nil || credential.JSON == nil {
			return nil, xerrors.New("defaultCredentials is nil")
		}
		credentialJSON = credential.JSON
	}

	conf, err := google.JWTConfigFromJSON(credentialJSON, storage.ScopeReadOnly)
	if err != nil {
		return nil, xerrors.Errorf("failed to google.JWTConfigFromJSON: %w", err)
	}
	if conf == nil {
		return nil, xerrors.New("config is nil")
	}

	var options = storage.SignedURLOptions{
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Method:         http.MethodGet,
	}
	return &cloudStorageClient{
		bucket: bucket,
		signedURLFunc: func(bucket, object string, expire time.Time) (string, error) {
			options.Expires = expire
			url, err := storage.SignedURL(bucket, object, &options)
			if err != nil {
				return "", xerrors.Errorf("failed to storage.SignedURL(bucket:%s, object:%s): %w", bucket, object, err)
			}
			return url, nil
		},
		uploadObjectFunc: func(ctx context.Context, bucket, object string, reader io.Reader) error {
			client, err := storage.NewClient(ctx)
			if err != nil {
				return xerrors.Errorf("failed to storage.NewClient: %w", err)
			}
			wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
			wc.ObjectAttrs.ContentType = "application/octet-stream"
			wc.ObjectAttrs.CacheControl = "no-cache"
			if _, err = io.Copy(wc, reader); err != nil {
				return xerrors.Errorf("failed to storage.NewClient.Bucket(%s).Object(%s).NewWriter: %w", bucket, object, err)
			}
			defer func() {
				if wc == nil {
					return
				}
				err = wc.Close()
			}()
			return nil
		},
	}, nil
}

func (c *cloudStorageClient) ExecSignedURL(object string, expire time.Time) (string, error) {
	if c == nil || object == "" {
		return "", xerrors.Errorf("does not meet the preconditions: [object:%s]", object)
	}
	return c.signedURLFunc(c.bucket, object, expire)
}

func (c *cloudStorageClient) ExecUploadObject(ctx context.Context, object string, reader io.Reader) error {
	if c == nil || object == "" || reader == nil {
		return xerrors.Errorf("does not meet the preconditions: [object:%s]", object)
	}
	return c.uploadObjectFunc(ctx, c.bucket, object, reader)
}

type SignedURLFunc func(bucket, object string, expire time.Time) (url string, err error)
type UploadObjectFunc func(ctx context.Context, bucket, object string, reader io.Reader) error
