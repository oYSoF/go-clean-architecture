package postgres

import (
	"lms/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgres_Connect(t *testing.T) {
	// To insure that a proper postgres config exists
	cfg, cfgErr := config.Parse("../config/config.json")
	assert.Empty(t, cfgErr)
	// To insure that your config leads to a connection.
	db, err := Connect(cfg.Postgres)
	assert.Empty(t, err)
	assert.NotEmpty(t, db)
}

func TestPostgres_ConnectionError(t *testing.T) {
	cfg, cfgErr := config.Parse("../config/example.config.json")
	assert.Empty(t, cfgErr)
	// To insure that a wrong config leads to an connection error.
	db, err := Connect(cfg.Postgres)
	assert.NotEmpty(t, err)
	assert.Empty(t, db)
}
