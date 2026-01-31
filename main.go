package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const DATAPATH = "data.json"

func loadData() (buckets map[string][]string) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		// File missing, create buckets manually:
		return make(map[string][]string)
	}

	err = json.Unmarshal(data, &buckets)
	if err != nil {
		fmt.Printf("Error unmarshalling data from disk-read-file. %v", err)
		os.Exit(0)
	}

	return buckets
}

func saveData(buckets map[string][]string) error {
	// Serializes and saves data to file path.
	data, err := json.MarshalIndent(buckets, "", " ")
	if err != nil {
		fmt.Printf("Error trying to Marshal buckets in saveData! Error:%v", err)
	}

	fmt.Printf("Here's how the data looks like thus far:\n\n%s\n\n", string(data))

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
	if len(os.Args) >= 3 {
		// NB: This is only the happy path, definitely Fix it up :)
		key := os.Args[1]
		val := strings.Join(os.Args[2:], " ")

		buckets[key] = append(buckets[key], val)
	}

	return errors.New("Too few args king, supply at least one! :)")

}

func main() {
	buckets := loadData()

	err := handleInput(buckets)
	if err != nil {
		fmt.Printf("Ooof... %v \n", err)
	}

	saveData(buckets)

	fmt.Println("Nice, enjoy!")
}
