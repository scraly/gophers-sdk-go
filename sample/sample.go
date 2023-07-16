package main

import (
	"context"
	"fmt"
	"os"

	gopherssdk "github.com/scraly/gophers-sdk-go"
)

func main() {
	config := gopherssdk.NewConfiguration()
	client := gopherssdk.NewAPIClient(config)

	// Check Health
	// When we call GophersApi.CheckHealth method, it return a string
	// equals to OK if the Gophers API is running and healthy

	health, healthRes, healthErr := client.GophersApi.CheckHealth(context.Background()).Execute()
	if healthErr != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersApi.CheckHealth``: %v\n", healthErr)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", healthRes)
	}
	// response from `CheckHealth`: string
	fmt.Fprintf(os.Stdout, "Response from `GophersApi.CheckHealth`: %v\n", health)

	// Get Gophers
	gophers, gophersRes, GophersErr := client.GophersApi.GophersGet(context.Background()).Execute()
	if GophersErr != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersApi.GophersGet``: %v\n", GophersErr)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", gophersRes)
	}

	// response from `GophersGet`: []Gopher
	if gophersRes.StatusCode == 200 {

		// Get and display all existing Gophers
		fmt.Println("Response from `GophersApi.GophersGet`:")

		fmt.Println("Number of Gophers:", len(gophers))

		for _, myGopher := range gophers {
			fmt.Println("DisplayName: ", *myGopher.Displayname)
			fmt.Println("Name:", *myGopher.Name)
			fmt.Println("URL:", *myGopher.Url)
		}

	}
}
