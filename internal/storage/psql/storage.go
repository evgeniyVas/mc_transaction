package psql

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Storage struct {
	conn *sqlx.DB
}

func New(dsn string, timeout time.Duration) (*Storage, error) {
	conn, err := Connection(dsn, timeout)
	if err != nil {
		return nil, err
	}

	return &Storage{
		conn: conn,
	}, nil
}

func (s *Storage) Close() error {
	err := s.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
