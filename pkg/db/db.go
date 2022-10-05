package db

import (
	"fmt"
	"rest-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbPassword,
	)

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: cfg.DbPreferSimpleProtocol,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbSchema(cfg),
		},
	})
}

func dbSchema(cfg *config.Config) string {
	var dbSchema string

	if len(cfg.DbSchema) > 0 {
		dbSchema = cfg.DbSchema + "."
	}

	return dbSchema
}
