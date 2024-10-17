package config

import (
	"errors"
	"os"
	"strconv"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port      int
		ApiSecret string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

var errApiPort = errors.New("error parsing env variable port")
var errApiPortMissing = errors.New("error api port is not present in env")
var errApiSecret = errors.New("error parsing env variable api_secret")
var errApiSecretMissing = errors.New("error api secret is not present in env")
var errDbHost = errors.New("error parsing env variable db_host")
var errDbHostMissing = errors.New("error db host is not present in env")
var errDbPort = errors.New("error parsing env variable db_port")
var errDbPortMissing = errors.New("error db port is not present in env")
var errDbUser = errors.New("error parsing env variable db_user")
var errDbUserMissing = errors.New("error db user is not present in env")
var errDbPassword = errors.New("error parsing env variable db_password")
var errDbPasswordMissing = errors.New("error db password is not present in env")
var errDbName = errors.New("error parsing env variable db_name")
var errDbNameMissing = errors.New("error db name is not present in env")

// Loads config from ENVIRONEMT
func GetConfig() (*Config, error) {
	port, ok := os.LookupEnv("port")
	if !ok {
		return nil, errApiPortMissing
	}
	app_port, err := strconv.Atoi(port)
	if err != nil {
		return nil, errApiPort
	}

	api_secret, ok := os.LookupEnv("api_secret")
	if !ok {
		return nil, errApiSecretMissing
	}
	if len(api_secret) == 0 {
		return nil, errApiSecret
	}

	app := App{
		Port:      app_port,
		ApiSecret: api_secret,
	}

	db_host, ok := os.LookupEnv("db_host")
	if !ok {
		return nil, errDbHostMissing
	}
	if len(db_host) == 0 {
		return nil, errDbHost
	}

	db_port_env, ok := os.LookupEnv("db_port")
	if !ok {
		return nil, errDbPortMissing
	}
	db_port, err := strconv.Atoi(db_port_env)
	if err != nil {
		return nil, errDbPort
	}

	db_user, ok := os.LookupEnv("db_user")
	if !ok {
		return nil, errDbUserMissing
	}
	if len(db_user) == 0 {
		return nil, errDbUser
	}

	db_password, ok := os.LookupEnv("db_password")
	if !ok {
		return nil, errDbPasswordMissing
	}
	if len(db_password) == 0 {
		return nil, errDbPassword
	}

	db_name, ok := os.LookupEnv("db_name")
	if !ok {
		return nil, errDbNameMissing
	}
	if len(db_name) == 0 {
		return nil, errDbName
	}

	db := Db{
		Host:     db_host,
		Port:     db_port,
		User:     db_user,
		Password: db_password,
		DBName:   db_name,
	}

	return &Config{
		App: app,
		Db:  db,
	}, nil
}
