package main

import (
	"lenta/models"
	"lenta/postgres"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var categories []models.Category
var posts []models.Post

func main() {
	postgres.InitConnection()

	go UpdateAllCacheByRepeat(time.Second * 10)

	r := gin.Default()
	r.GET("/api/ping", GetPong)
	r.GET("/api/categories", GetAllCategories)
	r.GET("/api/posts", GetAllPostsByLastWeek)
	r.Run()
}

func GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}

func GetAllCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func GetAllPostsByLastWeek(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func UpdateAllCacheByRepeat(duration time.Duration) {
	for range time.Tick(duration) {
		UpdateAllCache()
	}
}

func UpdateAllCache() {
	categories = postgres.GetAllCategories()
	posts = postgres.GetPostsByLastWeek()
}
