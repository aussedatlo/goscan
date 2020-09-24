package main

import (
	app "../../internal/app"
)

func main() {
	a := app.CreateApp()
	app.RunApp(a)
}
