package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
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
	chatsCollection *mongo.Collection
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
	usr, err := userGetter("name", name)
	if err == nil {
		return usr, errors.New("User already exists")
	}
	usr.Id, err = idGenerator("u")
	if err != nil {
		return usr, err
	}
	usr.Username = name
	usr.Created_at = time.Now()
	b, err := bson.Marshal(usr)
	if err != nil {
		return usr, err
	}
	_, err = usersCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return usr, err
	}
	return usr, nil
}

func chatCreator(json_map map[string]interface{}) (Chat, error) {
	var chat Chat
	var usr User
	var err error
	chat.Id, err = idGenerator("c")
	if err != nil {
		return chat, err
	}
	for k, v := range json_map {
		switch k {
		case "name":
			chat.Name = fmt.Sprint(v)
		case "users":
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					usr, err = userGetter("id", fmt.Sprint(s.Index(i)))
					if err != nil {
						return chat, err
					}
					chat.Users = append(chat.Users, usr)
				}
			}
		}
	}
	chat.Created_at = time.Now()
	b, err := bson.Marshal(usr)
	if err != nil {
		return chat, err
	}
	_, err = chatsCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func messageCreator(json_map map[string]interface{}) (Message, error) {
	var message Message
	var err error
	message.Id, err = idGenerator("m")
	if err != nil {
		return message, err
	}
	for k, v := range json_map {
		switch k {
		case "text":
			message.Text = fmt.Sprint(v)
		case "chat":
			_ = fmt.Sprint(v)
			message.Chat, err = chatGetter("id", fmt.Sprint(v))
			if err != nil {
				return message, err
			}
		case "author":
			message.Author, err = userGetter("id", fmt.Sprint(v))
			if err != nil {
				return message, err
			}
		}
	}
	message.Created_at = time.Now()
	return message, err
}

func userGetter(title string, value string) (User, error) {
	var result User
	res := usersCollection.FindOne(context.TODO(), bson.M{title: value})
	err := res.Decode(result)
	if err != nil {
		return result, errors.New("User not found")
	}
	return result, nil
}

func chatGetter(title string, value string) (Chat, error) {
	var result Chat
	res := chatsCollection.FindOne(context.TODO(), bson.M{title: value})
	err := res.Decode(result)
	if err != nil {
		return result, errors.New("Chat not found")
	}
	return result, nil
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
