package persistence

import (
	"server/config"
	"server/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Database struct {
	Instance *gorm.DB
}

func NewPersistentDB(dbPath string) *Database {
	db, err := newDB(dbPath)
	if err != nil {
		log.Fatalf("Could not open database %s. Error: %v", dbPath, err)
	}
	if config.GetEnvironment().ENVIRONMENT_NAME == config.DevelopmentEnvironment {
		model.LoadSeedData(db.Instance)
	}
	return db
}

func NewInMemoryDB() *Database {
	db, err := newDB(IN_MEMORY_DB)
	if err != nil {
		log.Fatalf("Could not open in-memory database. Error: %v", err)
	}
	return db
}

func newDB(dbPath string, migrationModels ...interface{}) (*Database, error) {
	db, err := newSqliteDB(dbPath, &model.Step{}, &model.Execution{}, &model.Output{}, &model.Status{}, &model.SessionConfig{})
	if err != nil {
		return nil, err
	}
	return &Database{
		Instance: db,
	}, nil
}
