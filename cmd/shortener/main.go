package main

import (
	"go-short/internal/app"
)

func main() {
	a := app.NewApp()
	a.Run(":8080")
}
