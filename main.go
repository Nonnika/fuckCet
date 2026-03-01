package main

import (
	"embed"
	aianalyse "fuckCet/backend/utils/aiAnalyse"
	"fuckCet/backend/words"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	app := NewApp()
	wordsList := words.NewWordList()
	db := words.NewDatabase()
	agent := aianalyse.NewAgent(os.Getenv("MODEL_NAME"), os.Getenv("MODEL_API_KEY"), os.Getenv("BASE_URL"))
	err = db.Init("./backend/source/vocabulary.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if connected, err := db.Connected(); err != nil || !connected {
		panic(err)
	}

	log.Println("Database connected successfully")

	wordsList.GetWords(db, "SELECT * FROM words WHERE source GLOB '3*'")
	log.Printf("Retrieved %d words from the database", len(wordsList.Words))

	wordsList.GetWordsDiscription(db)
	log.Printf("Retrieved descriptions for %d words", len(wordsList.Words))

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "fuckCet",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			wordsList,
			agent,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
