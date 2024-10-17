package config

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestGetConfig(t *testing.T) {
	correct := &Config{
		App{Port: 8080, ApiSecret: "secret"},
		Db{Port: 5432, Host: "localhost", User: "postgres", Password: "postgres", DBName: "backend"},
	}
	t.Run("correct config", func(t *testing.T) {
		generateConfig(true, true, true, true, true, true, true)
		result, _ := GetConfig()
		assert.DeepEqual(t, result, correct)
		resetConfig()
	})
	t.Run("missing port", func(t *testing.T) {
		generateConfig(false, false, false, false, false, false, false)
		t.Log(os.Getenv("port"))
		_, err := GetConfig()
		assertError(t, err, errApiPortMissing)
		resetConfig()
	})
	t.Run("missing api_secret", func(t *testing.T) {
		generateConfig(true, false, false, false, false, false, false)
		_, err := GetConfig()
		assertError(t, err, errApiSecretMissing)
		resetConfig()
	})
	t.Run("missing db_host", func(t *testing.T) {
		generateConfig(true, true, false, false, false, false, false)
		_, err := GetConfig()
		assertError(t, err, errDbHostMissing)
		resetConfig()
	})
	t.Run("missing db_port", func(t *testing.T) {
		generateConfig(true, true, true, false, false, false, false)
		_, err := GetConfig()
		assertError(t, err, errDbPortMissing)
		resetConfig()
	})
	t.Run("missing db_user", func(t *testing.T) {
		generateConfig(true, true, true, true, false, false, false)
		_, err := GetConfig()
		assertError(t, err, errDbUserMissing)
		resetConfig()
	})
	t.Run("missing db_password", func(t *testing.T) {
		generateConfig(true, true, true, true, true, false, false)
		_, err := GetConfig()
		assertError(t, err, errDbPasswordMissing)
		resetConfig()
	})
	t.Run("missing db_name", func(t *testing.T) {
		generateConfig(true, true, true, true, true, true, false)
		_, err := GetConfig()
		assertError(t, err, errDbNameMissing)
		resetConfig()
	})

	t.Run("invalid port", func(t *testing.T) {
		os.Setenv("port", "invalid")
		_, err := GetConfig()
		assertError(t, err, errApiPort)
		resetConfig()
	})
	t.Run("invalid api_secret", func(t *testing.T) {
		generateConfig(true, false, false, false, false, false, false)
		os.Setenv("api_secret", "")
		_, err := GetConfig()
		assertError(t, err, errApiSecret)
		resetConfig()
	})
	t.Run("invalid db_host", func(t *testing.T) {
		generateConfig(true, true, false, false, false, false, false)
		os.Setenv("db_host", "")
		_, err := GetConfig()
		assertError(t, err, errDbHost)
		resetConfig()
	})
	t.Run("invalid db_port", func(t *testing.T) {
		generateConfig(true, true, true, false, false, false, false)
		os.Setenv("db_port", "invalid")
		_, err := GetConfig()
		assertError(t, err, errDbPort)
		resetConfig()
	})
	t.Run("invalid db_user", func(t *testing.T) {
		generateConfig(true, true, true, true, false, false, false)
		os.Setenv("db_user", "")
		_, err := GetConfig()
		assertError(t, err, errDbUser)
		resetConfig()
	})
	t.Run("invalid db_password", func(t *testing.T) {
		generateConfig(true, true, true, true, true, false, false)
		os.Setenv("db_password", "")
		_, err := GetConfig()
		assertError(t, err, errDbPassword)
		resetConfig()
	})
	t.Run("invalid db_name", func(t *testing.T) {
		generateConfig(true, true, true, true, true, true, false)
		os.Setenv("db_name", "")
		_, err := GetConfig()
		assertError(t, err, errDbName)
		resetConfig()
	})
}

func generateConfig(port, api_secret, db_host, db_port, db_user, db_password, db_name bool) {
	if port {
		os.Setenv("port", "8080")
	}
	if api_secret {
		os.Setenv("api_secret", "secret")
	}
	if db_host {
		os.Setenv("db_host", "localhost")
	}
	if db_port {
		os.Setenv("db_port", "5432")
	}
	if db_user {
		os.Setenv("db_user", "postgres")
	}
	if db_password {
		os.Setenv("db_password", "postgres")
	}
	if db_name {
		os.Setenv("db_name", "backend")
	}
}

func resetConfig() {
	os.Unsetenv("port")
	os.Unsetenv("api_secret")
	os.Unsetenv("db_host")
	os.Unsetenv("db_port")
	os.Unsetenv("db_user")
	os.Unsetenv("db_password")
	os.Unsetenv("db_name")
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
