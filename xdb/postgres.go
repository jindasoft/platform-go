package xdb

import (
	"context"
	"fmt"
	"time"

	"github.com/jindasoft/jinda-platform/xlogger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresService interface {
	GetDB() *gorm.DB
	Health(ctx context.Context) error
	Close() error
	WithTx(tx *gorm.DB) PostgresService
	AutoMigrate(models ...interface{}) error
	Operations() *PostgresOperations
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string // disable, allow, prefer, require, verify-ca, verify-full
	IsDebug  bool
}

type postgresService struct {
	db *gorm.DB
}

func NewPostgresService(ctx context.Context, cfg *PostgresConfig) (*postgresService, error) {
	dsn := buildPostgresDSN(cfg)

	if cfg.IsDebug {
		xlogger.SysInfof("Connecting to PostgreSQL: %s:%d", cfg.Host, cfg.Port)
	}

	// Create logger for GORM
	gormLogger := logger.Default
	if !cfg.IsDebug {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	xlogger.SysInfof("PostgreSQL initialized successfully")
	return &postgresService{db: db}, nil
}

// AutoMigrate automatically migrates the database schema
func (s *postgresService) AutoMigrate(models ...interface{}) error {
	return s.db.AutoMigrate(models...)
}

func buildPostgresDSN(cfg *PostgresConfig) string {
	sslMode := cfg.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		sslMode,
	)
	return dsn
}

// GetDB returns the GORM database instance
func (s *postgresService) GetDB() *gorm.DB {
	return s.db
}

// Health checks the connection health
func (s *postgresService) Health(ctx context.Context) error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

// Close closes the database connection
func (s *postgresService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// WithTx returns a new service instance with a transaction
func (s *postgresService) WithTx(tx *gorm.DB) PostgresService {
	return &postgresService{db: tx}
}

// Operations returns a PostgresOperations instance for common CRUD operations
func (s *postgresService) Operations() *PostgresOperations {
	return NewPostgresOperations(s)
}
