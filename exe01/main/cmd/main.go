package main

import (
	"fmt"
	"main/internal/application"
)

func main() {

	app := application.NewDefaultHTTP(":8080", "1234")

	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
