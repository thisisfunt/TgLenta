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
