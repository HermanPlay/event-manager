package database

import (
	"testing"

	"github.com/HermanPlay/web-app-backend/internal/config"
)

func TestNewPostgresDatabase(t *testing.T) {
	cfg_error := getCfg(false)
	cfg := getCfg(true)
	t.Run("correct db config", func(t *testing.T) {
		_, err := NewPostgresDatabase(cfg)
		if err != nil {
			t.Errorf("could not establish db connection %s", err)
		}
	})
	t.Run("wrong db config", func(t *testing.T) {
		_, err := NewPostgresDatabase(cfg_error)
		assertError(t, err)
	})
}
func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}
func getCfg(correct bool) *config.Config {
	if correct {
		return &config.Config{
			Db: config.Db{
				User:     "postgres",
				Password: "postgres",
				Host:     "localhost",
				Port:     5432,
				DBName:   "test_db",
			},
		}
	} else {
		return &config.Config{
			Db: config.Db{
				User:     "testuser",
				Password: "testpassword",
				Host:     "localhost",
				Port:     5432,
				DBName:   "wrong",
			},
		}
	}
}

func TestConnect(t *testing.T) {
	cfg := getCfg(true)
	db, _ := NewPostgresDatabase(cfg)
	got := db.Connect()
	if got == nil {
		t.Errorf("recieved null when connecting to db")
	}

}
