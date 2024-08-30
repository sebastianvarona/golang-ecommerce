package main

import (
	"ecommerce_template/internal/models"
	"fmt"
	"os"
)

func main() {
	db := models.NewDB()

	entries, err := os.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile("migrations/" + entry.Name())
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(string(data))
		if err != nil {
			fmt.Printf("❌ Could not apply migration: %s\n", entry.Name())
		} else {
			fmt.Printf("✅ Applied migration: %s\n", entry.Name())
		}
	}
}
