package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/pedrosantosbr/proto-hornex/domain"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	// project path
	dbInstance.DsnTest = "/home/hert/Documents/study/golang/proto-hornex/framework/database/database.db"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	connection.AutoMigrate(&domain.User{})

	connection.Create(domain.User{

		FirstName:   "name",
		LastName:    "last",
		DateOfBirth: time.Now(),
		Active:      true,
		Email:       "name@email.com",
	})

	connection.Create(domain.User{
		FirstName:   "name",
		LastName:    "last",
		DateOfBirth: time.Now(),
		Active:      true,
		Email:       "name2@email.com",
	})

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {

	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		// d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		// d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil

}
