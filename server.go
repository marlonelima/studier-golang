package main

type App struct{}

func main() {
	app := App{}

	app.InitRoutes()
	app.Run("localhost:8080")
}