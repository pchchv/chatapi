package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "ChatAPI service. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

func addUserHandler(c echo.Context) error {
	username := c.QueryParam("name")
	user, err := userCreator(username)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, user)
}

func addChatHandler(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	chat, err := chatCreator(json_map)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, chat)
}

func addMessageHandler(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	message, err := messageCreator(json_map)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, message)
}

func usersChatsHandler(c echo.Context) error {
	userId := c.QueryParam("user")
	chats, err := chatsFinder(userId)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, chats)
}

func chatMessagesHandler(c echo.Context) error {
	chatId := c.QueryParam("chat")
	messages, err := messageFinder(chatId)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, messages)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.GET("/user/chats", usersChatsHandler)
	e.GET("/chat/messages", chatMessagesHandler)
	e.POST("/user", addUserHandler)
	e.POST("/chat", addChatHandler)
	e.POST("/message", addMessageHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
