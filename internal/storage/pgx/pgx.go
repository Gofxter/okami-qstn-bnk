package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PGX struct {
	logger *zap.Logger
	pool   *pgxpool.Pool
}

func NewPGX(logger *zap.Logger, path string) *PGX {
	pool, err := pgxpool.New(context.Background(), path)

	if err != nil {
		logger.Debug("unable to create connection pool", zap.Error(err))
		return nil
	}

	logger.Info("database initialized and successfully connected")
	return &PGX{
		pool:   pool,
		logger: logger,
	}
}

func (conn *PGX) Ping(ctx context.Context) error {
	return conn.pool.Ping(ctx)
}

func (conn *PGX) Close() {
	conn.pool.Close()
}
