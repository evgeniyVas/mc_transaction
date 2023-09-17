package psql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mc_transaction/internal/logger"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Storages struct {
	TransactionStorage *TransactionStorage
	BalanceStorage     *BalanceStorage
}

func New(dsn string, timeout time.Duration) (*Storages, error) {
	conn, err := Connection(dsn, timeout)
	if err != nil {
		return nil, err
	}

	return &Storages{
		TransactionStorage: NewTransactionStorage(conn),
		BalanceStorage:     NewBalanceStorage(conn),
	}, nil
}

func (s *Storages) Close() error {
	err := s.TransactionStorage.conn.Close()
	if err != nil {
		return err
	}
	err = s.BalanceStorage.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func rollBackTx(tx *sqlx.Tx, err error) {
	if err != nil {
		if tx == nil {
			logger.Error(fmt.Errorf("tx error: %w", err).Error())
		} else if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error(fmt.Errorf("rollback error: %w", err).Error())
		}
	}
}

func Connection(dsn string, timeout time.Duration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// почитать про пулы соединений в pq
	ctxBg := context.Background()
	pingTry := 0
	for {
		ctx, cancel := context.WithTimeout(ctxBg, timeout)
		err = db.PingContext(ctx)
		cancel()
		if err != nil {
			log.Print(fmt.Errorf("pg: try to ping. ERR => %s\n", err.Error()))
			if time.Duration(pingTry)*time.Second > 1*time.Minute {
				pingTry = 0
			}
			time.Sleep(time.Duration(pingTry) * time.Second)
			pingTry++
			continue
		}
		break
	}
	log.Print("pg: connected")
	return db, nil
}
