package configs

import (
	"context"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func InitDB() *pgx.Conn {
	_ = godotenv.Load()

	config, err := pgx.ParseConfig("postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("PostgreSQL connected successfully ✅")

	DB = conn
	return conn
}

func GetDB() *pgx.Conn {
	return DB
}

func CloseDB(ctx context.Context) {
	if DB != nil {
		if err := DB.Close(ctx); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}
}
