package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type albums struct {
	// gorm.Model
	AlbumId int64
	Title   string
	ArtisId int64
}

func get_all_albums() []albums {
	db, err := gorm.Open(sqlite.Open("./chinook.db"), &gorm.Config{})
	fmt.Println(db)
	if err != nil {
		panic("failed to connect database")
	}
	var datas []albums
	err = db.Find(&datas).Error
	if err != nil {
		panic(err)
	}
	// // defer db.Close()
	// migrator := db.Migrator()
	// tables := migrator.GetTables()

	// // Print the table names
	// for _, table := range tables {
	// 	fmt.Println(table)
	// }
	return datas
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var datas []albums = get_all_albums()
		c.JSON(http.StatusOK, gin.H{
			"data": datas,
		})
	})
	r.Run()
}
