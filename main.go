package main

import (
	"fmt"
	"net/http"

	"github.com/kevinliao852/goapi-example/utils"
)

func main() {
	app := utils.App{}

	app.Create()

	app.Register("GET", "/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "123")
	})

	app.Register("GET", "/greeting", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "hello")
	})

	app.Start(":3600")

	fmt.Println("end")
}
