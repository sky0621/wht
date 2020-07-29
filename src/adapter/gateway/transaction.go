package gateway

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type TxFunc func(ctx context.Context, txx *sqlx.Tx) (*TxResponse, error)

type TxResponse struct {
	CreatedID    int64 // 登録時のID
	AffectedRows int64 // 更新時の結果件数
}

// 渡されたロジックをトランザクション内で実行するためのデコレータ
func Tx(ctx context.Context, db *sqlx.DB, fn TxFunc) (*TxResponse, error) {
	var tx *sqlx.Tx
	{
		var err error
		tx, err = db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: 0, // RDBの設定に任せる
			ReadOnly:  false,
		})
		if err != nil {
			return nil, xerrors.Errorf("failed to begin transaction: %w", err)
		}
		if tx == nil {
			return nil, xerrors.New("failed at begin transaction: tx is nil")
		}
	}

	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic occurred: %+v", p)
		}
		if err := tx.Rollback(); err != nil {
			log.Printf("failed to commit transaction: %+v", err)
		}
	}()

	res, err := fn(ctx, tx)
	if err != nil {
		return nil, xerrors.Errorf("failed to exec function: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("failed to commit transaction: %w", err)
	}

	return res, err
}
