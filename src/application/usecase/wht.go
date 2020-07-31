package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/sky0621/wht/application/domain"
	"github.com/sky0621/wht/application/repository"
	"golang.org/x/xerrors"
)

type Wht interface {
	CreateWht(ctx context.Context, in *domain.WhtForCreate) (int64, error)
	CreateTextContents(ctx context.Context, recordDate time.Time, inputs []*domain.TextContentForCreate) error
	ReadWhts(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error)
	ReadTextContents(ctx context.Context, whtIDs []int64) (map[int64][]*domain.TextContent, error)
	GetWhtByRecordDate(ctx context.Context, recordDate time.Time) (*domain.Wht, error)
	UpsertWht(ctx context.Context, in *domain.WhtForCreate) (int64, error)
}

func NewWht(whtRepo repository.Wht, contentRepo repository.Content) Wht {
	return &wht{whtRepo: whtRepo, contentRepo: contentRepo}
}

type wht struct {
	whtRepo     repository.Wht
	contentRepo repository.Content
}

/*
 * CreateWht 「今日こと」を作成し、作成した「今日こと」のユニークIDを返す。
 * ただし、該当日の「今日こと」が作成済みの場合は作成せず、既存の「今日こと」のユニークIDを返す。
 */
func (w wht) CreateWht(ctx context.Context, in *domain.WhtForCreate) (int64, error) {
	already, err := w.GetWhtByRecordDate(ctx, in.RecordDate)
	if err != nil {
		return 0, xerrors.Errorf("failed to GetWhtByRecordDate[recordDate:%#+v]: %w", in.RecordDate, err)
	}
	if already != nil {
		fmt.Println("already exists") // TODO: use custom logger
		return already.ID, nil
	}
	return w.whtRepo.Create(ctx, in)
}

/*
 * CreateTextContent 「今日こと」のテキストコンテンツを作成し、作成済みの「今日こと」と紐付ける。
 * ただし、該当日の「今日こと」が未作成の場合は、「今日こと」を新規作成してから紐付ける。
 */
func (w wht) CreateTextContents(ctx context.Context, recordDate time.Time, inputs []*domain.TextContentForCreate) error {
	already, err := w.GetWhtByRecordDate(ctx, recordDate)
	if err != nil {
		return xerrors.Errorf("failed to GetWhtByRecordDate[recordDate:%#+v]: %w", recordDate, err)
	}

	var whtID int64
	if already == nil {
		var err error
		// まず、該当日の分の「今日こと」を作成
		whtID, err = w.whtRepo.Create(ctx, &domain.WhtForCreate{
			RecordDate: recordDate,
		})
		if err != nil {
			return xerrors.Errorf("failed to Create: %w", err)
		}
	} else {
		whtID = already.ID
	}

	// テキストコンテンツを作成
	if err := w.contentRepo.CreateTextContents(ctx, whtID, inputs); err != nil {
		return xerrors.Errorf("failed to CreateTextContent[whtID:%d][inputs:%#+v]: %w", whtID, inputs, err)
	}
	return nil
}

func (w wht) ReadWhts(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error) {
	return w.whtRepo.Read(ctx, condition)
}

func (w wht) ReadTextContents(ctx context.Context, whtIDs []int64) (map[int64][]*domain.TextContent, error) {
	contents, err := w.contentRepo.ReadTextContents(ctx, &domain.TextContentCondition{WhtIDs: whtIDs})
	if err != nil {
		return nil, xerrors.Errorf("failed to ReadTextContents[WhtIDs:%#+v]: %w", whtIDs, err)
	}
	results := make(map[int64][]*domain.TextContent)
	for _, content := range contents {
		for _, whtID := range whtIDs {
			if content.WhtID == whtID {
				results[whtID] = append(results[whtID], content)
			}
		}
	}
	return results, nil
}

func (w wht) GetWhtByRecordDate(ctx context.Context, recordDate time.Time) (*domain.Wht, error) {
	records, err := w.whtRepo.Read(ctx, &domain.WhtCondition{
		RecordDate: &recordDate,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to repository.Read[recordDate:%#+v]: %w", recordDate, err)
	}
	switch len(records) {
	case 0:
		return nil, nil
	case 1:
		return records[0], nil
	default:
		return nil, xerrors.New(fmt.Sprintf("failed to repository.Read[recordDate:%#+v]: not 1", recordDate))
	}
}

func (w wht) UpsertWht(ctx context.Context, in *domain.WhtForCreate) (int64, error) {
	// TODO: ?
	return w.whtRepo.Create(ctx, in)
}
