package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const DATAPATH = "data.json"

func loadData() map[string][]string {
	_, err := os.ReadFile("data.json")
	if err != nil {
		// File missing, create buckets manually:
		return make(map[string][]string)
	}

	return nil // TODO: fix - add actual correct return
}

func saveData(buckets map[string][]string) error {
	// Serializes and saves data to file path.
	data, err := json.MarshalIndent(buckets, "", " ")
	if err != nil {
		fmt.Printf("Error trying to Marshal buckets in saveData! Error:%v", err)
	}
	os.WriteFile(DATAPATH, data, 0644)

	return nil
}

func handleInput(buckets map[string][]string) error {
	/*
		Organize input into buckets as follows:
			left arg is the bucket name
			right arg is the content.
		In case there is only one arg, idea goes in the default bucket.

		Happy path implementation:
		bi vid: talk about how awesome Go is
	*/
	if len(os.Args) > 2 {
		if len(os.Args) > 3 {
			// NB: This is only the happy path, definitely Fix it up :)
			key := os.Args[1]
			val := os.Args[2:]

			buckets[key] = val
		}
		return nil
	}
	return errors.New("Too few args king, supply at least one! :)")
}

func main() {
	buckets := loadData()

	err := handleInput(buckets)
	if err != nil {
		fmt.Println(fmt.Errorf("Oops, remember happy path is <bucket:> <yapa yapa yapa>: %w", err))
	}

	saveData(buckets)

	fmt.Println("Nice, enjoy!")
}
