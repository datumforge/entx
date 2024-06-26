package entx

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

const (
	DefaultCacheTTL = 1 * time.Second
)

// Config Settings for the ent database client
type Config struct {
	// Debug to print debug database logs
	Debug bool `json:"debug" koanf:"debug" jsonschema:"description=debug enables printing the debug database logs" default:"false"`
	// DatabaseName is the name of the database to use with otel tracing
	DatabaseName string `json:"databaseName" koanf:"databaseName" jsonschema:"description=the name of the database to use with otel tracing" default:"datum"`
	// DriverName name from dialect.Driver
	DriverName string `json:"driverName" koanf:"driverName" jsonschema:"description=sql driver name, supported drivers include sqlite, libsql, and psql" default:"libsql"`
	// MultiWrite enabled writing to two databases simultaneously
	MultiWrite bool `json:"multiWrite" koanf:"multiWrite" jsonschema:"description=enables writing to two databases simultaneously" default:"false"`
	// PrimaryDBSource is the primary database source for all read and write operations
	PrimaryDBSource string `json:"primaryDbSource" koanf:"primaryDbSource" jsonschema:"description=dsn of the primary database,required" default:"file:datum.db"`
	// SecondaryDBSource for when multi write is enabled
	SecondaryDBSource string `json:"secondaryDbSource" koanf:"secondaryDbSource" jsonschema:"description=dsn of the secondary database if multi-write is enabled" default:"file:backup.db"`
	// CacheTTL to have results cached for subsequent requests
	CacheTTL time.Duration `json:"cacheTTL" koanf:"cacheTTL" jsonschema:"description=cache results for subsequent requests, defaults to 1s" default:"1s"`
	// RunMigrations to run migrations on startup
	RunMigrations bool `json:"runMigrations" koanf:"runMigrations" jsonschema:"description=run migrations on startup" default:"true"`
	// MigrationProvider to use for running migrations
	MigrationProvider string `json:"migrationProvider" koanf:"migrationProvider" jsonschema:"description=migration provider to use for running migrations" default:"atlas"`
	// EnableHistory to enable history data to be logged to the database
	EnableHistory bool `json:"enableHistory" koanf:"enableHistory" jsonschema:"description=enable history data to be logged to the database" default:"false"`
}

// EntClientConfig configures the entsql drivers
type EntClientConfig struct {
	// config contains the base database settings
	config Config
	// primaryDB contains the primary db connection
	primaryDB *entsql.Driver
	// secondaryDB contains the secondary db connection, if set
	secondaryDB *entsql.Driver
	// logger contains the zap logger
	logger *zap.SugaredLogger
}

// DBOption allows users to optionally supply configuration to the ent connection
type DBOption func(opts *EntClientConfig)

// NewDBConfig returns a new ent database configuration
func NewDBConfig(c Config, opts ...DBOption) (*EntClientConfig, error) {
	ec := &EntClientConfig{
		config: c,
		logger: zap.NewNop().Sugar(), // set a no-op logger by default
	}

	// setup primary db connection
	var err error

	// apply options
	for _, opt := range opts {
		opt(ec)
	}

	ec.primaryDB, err = ec.NewEntDB(c.PrimaryDBSource)
	if err != nil {
		ec.logger.Errorw("failed to create primary db connection", "error", err)

		return nil, err
	}

	return ec, nil
}

// GetPrimaryDB returns the primary database configuration
func (c *EntClientConfig) GetPrimaryDB() *entsql.Driver {
	return c.primaryDB
}

// GetSecondaryDB returns the secondary db connection
func (c *EntClientConfig) GetSecondaryDB() *entsql.Driver {
	return c.secondaryDB
}

// WithLogger sets the logger for the ent client
func WithLogger(l *zap.SugaredLogger) DBOption {
	return func(c *EntClientConfig) {
		c.logger = l
	}
}

// WithSecondaryDB sets the secondary db connection if the driver supports multiwrite
func WithSecondaryDB() DBOption {
	return func(c *EntClientConfig) {
		if !CheckMultiwriteSupport(c.config.DriverName) {
			c.logger.Fatalw("unsupported multiwrite driver", "driver", c.config.DriverName)
		}

		var err error

		c.secondaryDB, err = c.NewEntDB(c.config.SecondaryDBSource)
		if err != nil {
			c.logger.Fatalw("failed to create primary db connection", "error", err)
		}
	}
}

// NewEntDB creates a new ent database connection
func (c *EntClientConfig) NewEntDB(dataSource string) (*entsql.Driver, error) {
	entDialect, err := CheckEntDialect(c.config.DriverName)
	if err != nil {
		return nil, fmt.Errorf("failed checking dialect: %w", err)
	}

	// setup db connection
	db, err := otelsql.Open(c.config.DriverName, dataSource,
		otelsql.WithAttributes(attribute.String("db.system", c.config.DriverName)))
	if err != nil {
		return nil, fmt.Errorf("failed connecting to database: %w", err)
	}

	if err = otelsql.RegisterDBStatsMetrics(db,
		otelsql.WithAttributes(attribute.String("db.system", c.config.DriverName)),
	); err != nil {
		return nil, fmt.Errorf("failed registering database metrics for otelsql: %w", err)
	}

	// enable foreign keys for libsql
	if c.config.DriverName == "libsql" {
		if _, err := db.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to enable enable foreign keys: %w", err)
		}
	}

	// verify db connection using ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed verifying database connection: %w", err)
	}

	return entsql.OpenDB(entDialect, db), nil
}

// Healthcheck pings the DB to check if the connection is working
func Healthcheck(client *entsql.Driver) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		if err := client.DB().Ping(); err != nil {
			return fmt.Errorf("db connection failed: %w", err)
		}

		return nil
	}
}

// CheckEntDialect checks if the dialect is supported and returns the ent dialect
// corresponding to the given dialect
func CheckEntDialect(d string) (string, error) {
	switch d {
	case "sqlite3":
		return dialect.SQLite, nil
	case "libsql":
		return dialect.SQLite, nil
	case "postgres":
		return dialect.Postgres, nil
	default:
		return "", newDialectError(d)
	}
}

// CheckMultiwriteSupport checks if the dialect supports multiwrite
func CheckMultiwriteSupport(d string) bool {
	switch d {
	case "sqlite3":
		return true
	case "libsql":
		return true
	case "postgres":
		return false
	default:
		return false
	}
}
