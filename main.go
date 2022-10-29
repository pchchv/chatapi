package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
)

var (
	usersCollection *mongo.Collection
	seededRand      *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

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
	// TODO: Checking user existence
	u := User{}
	var err error
	u.Id, err = idGenerator("u")
	if err != nil {
		return u, err
	}
	u.Username = name
	u.Created_at = time.Now()
	b, err := bson.Marshal(u)
	if err != nil {
		return u, err
	}
	_, err = usersCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return u, err
	}
	return u, nil
}

func chatCreator(json_map map[string]interface{}) (Chat, error) {
	var chat Chat
	return chat, nil
}

func idGenerator(mode string) (string, error) {
	id := strGenerator(numbers, 1+rand.Intn(5)) + strGenerator(letters+numbers, 1+rand.Intn(5))
	switch mode {
	case "u":
		id = "uuu" + id
	case "c":
		id = "chh" + id
	case "m":
		id = "sss" + id
	default:
		return id, errors.New("Mode error")
	}
	return id, nil
}

func strGenerator(charset string, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	db()
	server()
}
