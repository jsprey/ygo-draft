package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/config"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrorNotConnected = errors.New("postgresql client is not connected")

// PostgresClient is responsible for the communication between our backend server and our postgresql database
type PostgresClient struct {
	PoolConnection *pgxpool.Pool
}

func NewPostgresClient(dbCtx config.DbContext) (*PostgresClient, error) {
	client := PostgresClient{}
	err := client.Connect(dbCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database [%s]: %w", dbCtx.DatabaseUrl, err)
	}

	return &client, nil
}

// Connect connects the server to the instance of the database
func (p *PostgresClient) Connect(dbCtx config.DbContext) error {
	if p.PoolConnection != nil {
		p.PoolConnection.Close()
		p.PoolConnection = nil
	}

	logrus.Debugf("Connecting to the database: %s", dbCtx.DatabaseUrl)

	dbPool, err := pgxpool.Connect(context.Background(), dbCtx.GetConnectionUrl())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	logrus.Debugf("Connection successful!")
	p.PoolConnection = dbPool

	return nil
}

// Disconnect disconnects the connection to the database
func (p *PostgresClient) Disconnect() {
	p.PoolConnection.Close()
	p.PoolConnection = nil
}

// IsConnected returns true when the connection to the database is currently established
func (p *PostgresClient) IsConnected() bool {
	return p.PoolConnection != nil
}

// Query executes a query against postgresql
func (p *PostgresClient) Query(query string, args ...any) (pgx.Rows, error) {
	logrus.Tracef("PostgreSQL-Client -> Query: [%s]", query)

	if !p.IsConnected() {
		return nil, fmt.Errorf("%s", ErrorNotConnected)
	}

	if len(args) == 0 {
		return p.PoolConnection.Query(context.Background(), query)
	}

	return p.PoolConnection.Query(context.Background(), query, args)
}

// QueryRow executes a query against postgresql and delivers only one row.
func (p *PostgresClient) QueryRow(query string, args ...any) (pgx.Row, error) {
	logrus.Tracef("PostgreSQL-Client -> QueryRow: [%s]", query)

	if !p.IsConnected() {
		return nil, fmt.Errorf("%s", ErrorNotConnected)
	}

	if len(args) == 0 {
		return p.PoolConnection.QueryRow(context.Background(), query), nil
	}

	return p.PoolConnection.QueryRow(context.Background(), query, args), nil
}

// Exec executes a statement against postgresql.
func (p *PostgresClient) Exec(query string, args ...any) (pgconn.CommandTag, error) {
	logrus.Tracef("PostgreSQL-Client -> Exec: [%s]", query)

	if !p.IsConnected() {
		return nil, fmt.Errorf("%s", ErrorNotConnected)
	}

	if len(args) == 0 {
		return p.PoolConnection.Exec(context.Background(), query)
	}

	return p.PoolConnection.Exec(context.Background(), query, args)
}

// Select performs a select and scans the data into a struct.
func (p *PostgresClient) Select(query string, targetObject any, args ...any) error {
	logrus.Tracef("PostgreSQL-Client -> Select: [%s]", query)

	if !p.IsConnected() {
		return fmt.Errorf("%s", ErrorNotConnected)
	}

	if len(args) == 0 {
		return pgxscan.Select(context.Background(), p.PoolConnection, targetObject, query)
	}

	return pgxscan.Select(context.Background(), p.PoolConnection, targetObject, query, args)
}
