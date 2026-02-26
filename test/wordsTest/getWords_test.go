package wordstest

import (
	"fuckCet/backend/words"
	"testing"
)

const testDbPath = "../../backend/source/vocabulary.db"

func TestWordsList(t *testing.T) {
	db := words.NewDatabase()
	err := db.Init(testDbPath)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	if connected, err := db.Connected(); err != nil || !connected {
		t.Fatal(err)
		t.Fatal("Database is not connected")
	}

	wordList := words.NewWordList()
	wordList.GetWords(db, "SELECT * FROM words WHERE source GLOB '3*'")

	if len(wordList.Words) == 0 {
		t.Fatal("Expected to retrieve words, but got an empty list")
	}

	for _, word := range wordList.Words {
		t.Logf("Retrieved word: ID=%d, Word=%s, Src=%s", word.Id, word.Word, word.Src)
	}

}
