package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"fmt"


	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func InitDB() *sql.DB {
	connStr := "root:Fkwini2002@tcp(localhost:3306)/BLOG?parseTime=true"
	var err error
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal ping database:", err)
	}

	return db
}

func LoggerMiddleware(c *fiber.Ctx) error {
	fmt.Printf("Request:", c.Method(), c.Path())
	return c.Next()
}
