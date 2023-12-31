package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func writeFromReader(reader io.Reader, writer io.Writer, bufSize int, item any) {

	var buffer []byte = make([]byte, bufSize)
	for {
		//fmt.Print(item)
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}
		writer.Write(buffer[0:n])
	}
	//wg.

}

type Activity struct {
	ID        int32
	Name      string
	TargetUrl string
}

func (Activity) TableName() string {
	return "Activity"
}

func OpenDb(dbName string) *gorm.DB {

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

	userName := os.Getenv("user_name")
	password := os.Getenv("password")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s", userName, password, dbName)
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

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	var server = gin.Default()

	server.POST("/person", func(ctx *gin.Context) {

		var wg *sync.WaitGroup = &sync.WaitGroup{}

		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 1<<30)
		ctx.Request.ParseMultipartForm(8 << 10)

		for k, v := range ctx.Request.Form {
			fmt.Println(">>>", k, v)
		}

		file, fileheader, _ := ctx.Request.FormFile("uploadfile")
		file2, fileheader2, _ := ctx.Request.FormFile("uploadfile2")

		go func() {

			wg.Add(1)
			outputFile, err := os.CreateTemp(".", "*"+path.Ext(fileheader.Filename))
			if err != nil {
				fmt.Println(err)
			}
			writeFromReader(file, outputFile, 32, "1")
			outputFile.Close()
			wg.Done()

		}()

		go func() {

			wg.Add(1)
			outputFile1, err := os.CreateTemp(".", "*"+path.Ext(fileheader2.Filename))
			if err != nil {
				fmt.Println(err)
			}

			writeFromReader(file2, outputFile1, 32, "2")
			outputFile1.Close()
			wg.Done()

		}()

		fmt.Println(file, fileheader)
		wg.Wait()

	})

	server.GET("/test", func(ctx *gin.Context) {

<<<<<<< HEAD
		db := OpenDb("kludemy_test")
=======
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

		var dsn = "host=localhost user=postgres password=*** dbname=***"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			fmt.Println(err)
		}
>>>>>>> d0ea3285a19b6ecaa0e10593269955d9150339e4

		var activity []Activity
		db.Where("id < 10").Find(&activity)
		fmt.Println(activity)
		ctx.JSON(200, activity)

	})

	server.GET("/test2", func(ctx *gin.Context) {
		type Product struct {
			gorm.Model
			price float64
		}

		db := OpenDb("test")
		db.AutoMigrate(&Product{})
	})

	fmt.Println("http://localhost:8082")
	server.Run("localhost:8082")
}
