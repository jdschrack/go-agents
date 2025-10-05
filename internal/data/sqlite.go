package data

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DbContextKey struct{}

func WithDatabase(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, DbContextKey{}, db)
}

func FromContext(ctx context.Context) *sql.DB {
	db, ok := ctx.Value(DbContextKey{}).(*sql.DB)
	if !ok {
		return nil
	}
	return db
}

func GetConnection(ctx context.Context, dbPath string, readonly bool) (*sql.DB, error) {
	dsn := dbPath + "?_journal=WAL&_synchronous=NORMAL&_fk=true"

	if readonly {
		dsn += "&mode=ro"
	}

	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	pragmas := []string{
		"PRAGMA cache_size = -64000",
		"PRAGMA temp_store = MEMORY",
		"PRAGMA mmap_size = 268435456",
	}

	for _, pragma := range pragmas {
		if _, err := conn.Exec(pragma); err != nil {
			return nil, err
		}
	}

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(1)
	conn.SetConnMaxLifetime(0)

	return conn, nil
}
