package postgres

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	// db driver
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

type Database struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

type Config struct {
	DSN             string        `koanf:"dsn" validate:"required"`
	MaxOpenConns    int           `koanf:"maxOpenConns" validate:"required"`
	ConnMaxLifetime time.Duration `koanf:"connMaxLifetime" validate:"required"`
	MaxIdleConns    int           `koanf:"maxIdleConns" validate:"required"`
	ConnMaxIdleTime time.Duration `koanf:"connMaxIdleTime" validate:"required"`
}

func New(logger *zap.SugaredLogger, cfg Config) (*Database, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Database{
		db:     db,
		logger: logger,
	}, nil
}

// connect tries to connect to postgres with given config
func connect(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		return &sqlx.DB{}, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime * time.Second)

	err = db.Ping()
	if err != nil {
		return &sqlx.DB{}, err
	}

	return db, nil
}

func (d *Database) Disconnect() {
	err := d.db.Close()
	if err != nil {
		d.logger.Errorf("error while closing postgres connection: %s", err)
	}

	d.logger.Info("postgres connection closed successfully")
}

func (d *Database) DB() *sqlx.DB {
	return d.db
}
