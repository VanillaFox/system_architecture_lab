package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()

	// TODO: make with ENVs
	pgConn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?target_session_attrs=read-write",
		"conf",
		"pass",
		"postgres",
		5432,
		"conf",
	)

	cfg, err := pgxpool.ParseConfig(pgConn)
	if err != nil {
		panic(err)
	}

	cfg.ConnConfig.PreferSimpleProtocol = true

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(pool)
	fmt.Println("Hello, world!")
}
