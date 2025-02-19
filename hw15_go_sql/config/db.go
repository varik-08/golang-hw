package config

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func NewDB(ctx context.Context, dbCfg DB) (*pgxpool.Pool, error) {
	connConfig, err := pgx.ParseConfig(
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?TimeZone=Europe/Moscow",
			dbCfg.User,
			dbCfg.Password,
			net.JoinHostPort(dbCfg.Host, strconv.Itoa(dbCfg.Port)),
			dbCfg.Database,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create DSN for DB connection: %w", err)
	}
	dbc, err := pgxpool.New(ctx, connConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB : %w", err)
	}

	err = dbc.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return dbc, nil
}

func InitDB(ctx context.Context, dbCfg DB) {
	once.Do(func() {
		conn, err := NewDB(ctx, dbCfg)
		if err != nil {
			log.Fatalf("Ошибка инициализации БД: %v", err)
		}
		pool = conn

		log.Println("Connected to database")
	})
}

func GetDB() *pgxpool.Pool {
	if pool == nil {
		log.Fatal("БД не инициализирована! Вызовите InitDB() перед использованием.")
	}
	return pool
}
