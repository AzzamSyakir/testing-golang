package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
	"github.com/rs/zerolog/log"
)

// InitDB digunakan untuk menghubungkan ke database.
func InitDB() (*sql.DB, error) {

	//baca env nya
	sqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Membuka koneksi ke database
	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal membuka koneksi database")
	}

	// Memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal melakukan ping ke database")
	}

	log.Info().Msg("Terhubung ke database!")

	return db, nil
}
func InitDBTest() (*sql.DB, error) {

	//baca env nya
	sqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME_test"),
	)

	// Membuka koneksi ke database
	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal membuka koneksi database")
	}

	// Memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal melakukan ping ke database")
	}

	log.Info().Msg("Terhubung ke database!")

	return db, nil
}
