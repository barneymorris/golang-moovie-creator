package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// TODO Move to config
var DATABASE_URL string = "postgres://barney:barney@localhost:5432/barney"

type PostgreseClient struct {
	Connection *pgx.Conn
}

func GetNewPostgresClient() *PostgreseClient {
	conn := GetPostgresConnection()

	return &PostgreseClient{
		Connection: conn,
	}
}

func GetPostgresConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.Close(context.Background())

	return conn
}