package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type User struct {
	Id         string
	Username   string
	Created_at time.Time
}

type Chat struct {
	Id         string
	Name       string
	Users      []User
	Created_at time.Time
}

type Message struct {
	Id         string
	Chat       Chat
	Author     User
	Text       string
	Created_at time.Time
}

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func userCreator(name string) (User, error) {
	u := User{}
	u.Id = idGenerator()
	u.Username = name
	u.Created_at = time.Now()
	// TODO: Add a user to the db
	return u, nil
}

func idGenerator() string {
	var id string
	// TODO: Generate id
	return id
}

func main() {
	server()
}
