// Package database contains database connection logic and data-access repositories.
package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/seanvm/serverless-golang-starter-backend/app"
	"github.com/seanvm/serverless-golang-starter-backend/app/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Datastore struct {
	UserRepository app.UserRepository
}

type Connection struct {
}

type DbCredentials struct {
	Username     string
	Password     string
	Host         string
	Port         int32
	DatabaseName string
	TimeZone     string
}

// NewDb initializes a new database connection and injects it into the repositories
func Db() *Datastore {
	log.Printf("Initializing DB Connections \n")

	db := connect()

	ds := Datastore{}
	ds.UserRepository = &UserRepository{db}

	return &ds
}

func connect() *gorm.DB {
	dsn := connectionString(getDbCredentials())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println(err)
		errorString := fmt.Sprintf("Unable to connect to database: %v \n", err)
		panic(errorString)
	}

	return db
}

func connectionString(creds DbCredentials) string {
	values := url.Values{}
	values.Add("sslmode", "disable")
	values.Add("TimeZone", "Canada/Pacific")

	dsn := url.URL{
		User:     url.UserPassword(creds.Username, creds.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", creds.Host, creds.Port),
		Path:     creds.DatabaseName,
		RawQuery: (&values).Encode(),
	}

	return dsn.String()
}

// getDbCredentials retrieves the database credentials depending on the app's current environment.
func getDbCredentials() DbCredentials {
	creds := DbCredentials{}

	if helpers.IsOffline() {
		// Serverless offline runs inside a docker container
		creds.Host = "host.docker.internal"
	}

	if helpers.IsTestEnv() {
		creds.Host = "localhost"
	}

	if helpers.IsOffline() || helpers.IsTestEnv() {
		log.Println("Offline: Using Localhost")
		creds.Port = 5432
		creds.Username = "postgres"
		creds.Password = "postgres"
		creds.DatabaseName = "influencer_platform"
	} else {
		// Production Credentials go here
		// Using panic here to ensure that a new container gets started in case secrets/config are outdated
		panic("Error retrieving DB creds")
	}

	return creds
}
