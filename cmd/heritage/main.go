package main

import (
	"fmt"
	"heritage/internal/app"
	"heritage/internal/store"
)

func main() {
	db, err := store.Open("heritage.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	service := app.New(db)
	rec, err := service.Register("R001", "昆曲口述史", "李梅", "江苏", 1985, 1200)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(service.Summary(rec.ID))
}
