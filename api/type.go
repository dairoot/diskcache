package api

import (
	"context"
	"database/sql"
)

type DiskCache struct {
	Ctx  context.Context
	Conn *sql.Conn
}
