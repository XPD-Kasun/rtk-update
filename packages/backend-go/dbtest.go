package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDb2(dbName string) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dsn := "host=localhost user=postgres password=xpd dbname=test"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return db
}

func main() {

	type Ping struct {
		ItemCode sql.NullString `gorm:"size:20;type:text;not null;"`
	}

	type Product struct {
		gorm.Model
		Ping
		Title string    `gorm:"size:20;unique;non null;default:test;->:false;<-:create"`
		Price float64   `gorm:"precision:12;scale:2"`
		Colx  time.Time `gorm:"comment:this is another db"`
	}

	db := OpenDb2("test")
	db.AutoMigrate(&Product{})
	s := &Product{}
	db.First(s)
	fmt.Println(s)
	s.Title = "dsfdsf342345"
	s.Price = 349343.34
	s.Colx = time.Now()
	db.Save(s)
	fmt.Println(s.ItemCode.String)
}
