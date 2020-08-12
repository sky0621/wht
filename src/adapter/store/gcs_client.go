package store

import (
	"context"
	"io"
	"time"

	"github.com/rs/zerolog/log"
	"gocloud.dev/blob"
	"golang.org/x/xerrors"
)

type CloudStorageClient interface {
	ExecSignedURL(ctx context.Context, bucketPurpose BucketPurpose, object string, expire time.Duration) (url string, err error)
	ExecUploadObject(ctx context.Context, bucketPurpose BucketPurpose, object string, reader io.Reader) error
}

// GCSバケットの使用目的
type BucketPurpose int

const (
	// 未定義
	BucketPurposeUndefined = iota // 異常扱いと同義なので +1 しない。
	// 画像コンテンツ置き場
	ImageContentsBucket
)

type cloudStorageClient struct {
	bucketNameMap    map[BucketPurpose]string
	signedURLFunc    SignedURLFunc
	uploadObjectFunc UploadObjectFunc
}

func NewCloudStorageClient(ctx context.Context, bucketNameMap map[BucketPurpose]string) (CloudStorageClient, error) {
	log.Debug().Msg("NewCloudStorageClient___START")

	//var credentialJSON []byte
	//{
	//	log.Debug().Msg("cfg.appCredentials is blank")
	//	credential, err := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	//	if err != nil {
	//		return nil, xerrors.Errorf("failed to google.FindDefaultCredentials: %w", err)
	//	}
	//	if credential == nil || credential.JSON == nil {
	//		return nil, xerrors.New("defaultCredentials is nil")
	//	}
	//	credentialJSON = credential.JSON
	//}
	//log.Debug().Msgf("credentialJSON: %s", string(credentialJSON))
	//
	//var options storage.SignedURLOptions
	//{
	//	conf, err := google.JWTConfigFromJSON(credentialJSON, storage.ScopeReadOnly)
	//	if err != nil {
	//		return nil, xerrors.Errorf("failed to google.JWTConfigFromJSON: %w", err)
	//	}
	//	if conf == nil {
	//		return nil, xerrors.New("config is nil")
	//	}
	//
	//	options = storage.SignedURLOptions{
	//		GoogleAccessID: conf.Email,
	//		PrivateKey:     conf.PrivateKey,
	//		Method:         http.MethodGet,
	//	}
	//}

	return &cloudStorageClient{
		bucketNameMap: bucketNameMap,
		signedURLFunc: func(ctx context.Context, bucket, object string, expire time.Duration) (string, error) {
			b, cleaner, err := getBucket(ctx, bucket)
			if err != nil {
				return "", xerrors.Errorf("Failed to setup bucket: %w", err)
			}
			defer cleaner()

			var url string
			{
				var err error
				url, err = b.SignedURL(ctx, object, &blob.SignedURLOptions{
					Expiry: expire,
				})
				if err != nil {
					return "", xerrors.Errorf("failed to storage.SignedURL(bucket:%s, object:%s): %w", bucket, object, err)
				}
			}
			return url, nil
		},
		uploadObjectFunc: func(ctx context.Context, bucket, object string, reader io.Reader) error {
			b, cleaner, err := getBucket(ctx, bucket)
			if err != nil {
				return xerrors.Errorf("Failed to setup bucket: %w", err)
			}
			defer cleaner()

			var writer *blob.Writer
			{
				var err error
				writer, err = b.NewWriter(ctx, object, &blob.WriterOptions{
					ContentType:  "application/octet-stream",
					CacheControl: "no-cache",
				})
				if err != nil {
					return xerrors.Errorf("Failed to NewWriter: %w", err)
				}
				defer func() {
					if err := writer.Close(); err != nil {
						log.Err(err).Send()
					}
				}()

				_, err = writer.ReadFrom(reader)
				if err != nil {
					return xerrors.Errorf("Failed to ReadFrom: %w", err)
				}
			}
			return nil
		},
	}, nil
}

func (c *cloudStorageClient) ExecSignedURL(ctx context.Context, bucketPurpose BucketPurpose, object string, expire time.Duration) (string, error) {
	if c == nil || object == "" {
		return "", xerrors.Errorf("does not meet the preconditions: [object:%s]", object)
	}
	return c.signedURLFunc(ctx, c.bucketNameMap[bucketPurpose], object, expire)
}

func (c *cloudStorageClient) ExecUploadObject(ctx context.Context, bucketPurpose BucketPurpose, object string, reader io.Reader) error {
	if c == nil || object == "" || reader == nil {
		return xerrors.Errorf("does not meet the preconditions: [object:%s]", object)
	}
	return c.uploadObjectFunc(ctx, c.bucketNameMap[bucketPurpose], object, reader)
}

type SignedURLFunc func(ctx context.Context, bucket, object string, expire time.Duration) (url string, err error)
type UploadObjectFunc func(ctx context.Context, bucket, object string, reader io.Reader) error

func getBucket(ctx context.Context, bucket string) (*blob.Bucket, func(), error) {
	b, err := blob.OpenBucket(ctx, bucket)
	if err != nil {
		return nil, nil, xerrors.Errorf("Failed to setup bucket: %w", err)
	}
	return b, func() {
		if err := b.Close(); err != nil {
			log.Err(err).Send()
		}
	}, nil
}
