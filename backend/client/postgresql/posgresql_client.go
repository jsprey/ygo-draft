package postgresql

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/config"
)

// PosgresqlClient handles the communication with the postgresql database.
type PosgresqlClient struct {
	Connection *pgx.Conn
}

// NewPosgresqlClient creates a new postgresql client with the given database information.
func NewPosgresqlClient(dbCtx config.DbContext) (*PosgresqlClient, error) {
	logrus.Debugf("PostgresSQL-Client -> Create new client for %s", dbCtx.DatabaseUrl)

	conn, err := pgx.Connect(context.Background(), dbCtx.GetConnectionUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to connect do database: %w", err)
	}

	return &PosgresqlClient{
		Connection: conn,
	}, nil
}

// Close closes the connection to the internal database.
func (pc *PosgresqlClient) Close() error {
	return pc.Connection.Close(context.Background())
}

// Query perform a query against the database.
func (pc *PosgresqlClient) Query(query string, args ...any) (pgx.Rows, error) {
	if len(args) == 0 {
		rows, err := pc.Connection.Query(context.Background(), query)
		if err != nil {
			return nil, err
		}
		return rows, nil
	} else {
		rows, err := pc.Connection.Query(context.Background(), query, args)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
}

// QueryRow perform a row query against the database.
func (pc *PosgresqlClient) QueryRow(query string, args ...any) pgx.Row {
	if len(args) == 0 {
		return pc.Connection.QueryRow(context.Background(), query)
	} else {
		return pc.Connection.QueryRow(context.Background(), query, args)
	}
}

// Exec performs an exec against the database.
func (pc *PosgresqlClient) Exec(query string, args ...any) error {
	logrus.Tracef("PostgreSQL-Client -> performing query [%s]", query)

	if len(args) == 0 {
		_, err := pc.Connection.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	} else {
		_, err := pc.Connection.Exec(context.Background(), query, args)
		if err != nil {
			return err
		}
	}

	return nil
}

// Select performs an exec against the database.
func (pc *PosgresqlClient) Select(dst interface{}, query string, args ...interface{}) error {
	logrus.Tracef("PostgreSQL-Client -> performing select query [%s]", query)

	if len(args) == 0 {
		err := pgxscan.Select(context.Background(), pc.Connection, dst, `SELECT id, name, email, age FROM users`)
		if err != nil {
			return err
		}
	} else {
		err := pgxscan.Select(context.Background(), pc.Connection, dst, `SELECT id, name, email, age FROM users`, args)
		if err != nil {
			return err
		}
	}

	return nil
}
