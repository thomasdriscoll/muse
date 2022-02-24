package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func ConnectPostgreSQLDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/muse_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v \n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn
}
