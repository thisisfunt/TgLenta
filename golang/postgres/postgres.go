package postgres

import (
	"database/sql"
	"fmt"
	"lenta/models"

	_ "github.com/lib/pq"
)

var con *sql.DB

func InitConnection() {
	var err error
	con, err = sql.Open("postgres", "user=postgres password=postgres dbname=lenta host=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetAllCategories() []models.Category {
	rows, err := con.Query("SELECT id, tag, name from categories")
	if err != nil {
		fmt.Println("Problem with gettinng categories:", err)
		return []models.Category{}
	}
	categories := make([]models.Category, 0)
	for rows.Next() {
		category := models.Category{}
		err := rows.Scan(&category.Id, &category.Tag, &category.Name)
		if err != nil {
			fmt.Println("Problem with getting categories:", err)
			return []models.Category{}
		}
		categories = append(categories, category)
	}
	return categories
}

func GetPostsByLastWeek() []models.Post {
	rows, err := con.Query(`
	SELECT p.id, p.text, p.ref, p.publication_date, c.logo_ref, c.name, c.subscribers_count, c.category_id, p.views, p.redirected, c.premium
	FROM posts as p
	LEFT JOIN channels as c
	ON p.channel_id=c.id
	WHERE p.publication_date >= CURRENT_DATE - INTERVAL '7 days';`)
	if err != nil {
		fmt.Println("Problem with getting posts:", err)
		return []models.Post{}
	}
	posts := make([]models.Post, 0)
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.Text, &post.Ref, &post.PublicationDate, &post.LogoRef, &post.ChannelName, &post.SubscribersCount, &post.CategoryId, &post.Views, &post.Redirected, &post.IsChannelPremium)
		if err != nil {
			fmt.Println("Problem with getting posts:", err)
		}
		posts = append(posts, post)
	}
	return posts
}
