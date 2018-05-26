package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/kyokomi/emoji.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const ApiUri = "http://api.openweathermap.org/data/2.5/weather"

type Result struct {
	Name    string
	Main    Main
	Sys     Sys
	Weather [1]Weather
}

type Main struct {
	Temp     float64
	Humidity float64
}

type Sys struct {
	Country string
}

// "weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}]
type Weather struct {
	Description string
	Icon        string
}

var weatherEmojiMap = map[string]string{
	"11d": ":cloud_with_lightning:",
	"09d": ":cloud_with_rain:",
	"10d": ":umbrella: :cloud_with_rain:",
	"13d": ":snowflake:",
	// Lots of cloudy weather variations
	"02d": ":cloud:",
	"02n": ":cloud:",
	"03d": ":cloud:",
	"03n": ":cloud:",
	"04d": ":cloud:",
	"04n": ":cloud:",
}

func weatherEmoji(icon string) string {
	weatherEmoji := weatherEmojiMap[icon]

	if weatherEmoji != "" {
		weatherEmoji = weatherEmoji + " "
	}

	return weatherEmoji
}

func HumanizedWeatherMessage(input string, apiKey string) string {
	client := &http.Client{}

	if input == "" {
		fmt.Println("Please, provide the name of the city.")
		os.Exit(1)
	}

	request, err := http.NewRequest("GET", ApiUri, nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Building query params
	query := request.URL.Query()
	query.Add("appid", apiKey)
	query.Add("q", input)
	query.Add("units", "metric")

	// Encoding silly data coming from user input as well
	request.URL.RawQuery = query.Encode()

	// Explicitly building headers
	request.Header.Add("Accept", "application/json")

	// Making the actual request and "scheduling" body closure
	response, err := client.Do(request)
	defer response.Body.Close()

	responseStruct := Result{}

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)

		// Changes &responseBody struct with body data
		jsonErr := json.Unmarshal(body, &responseStruct)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}

	weatherEmoji := weatherEmoji(responseStruct.Weather[0].Icon)

	return emoji.Sprintf("%v, %v is under %v (%v%v°C)",
		responseStruct.Name,
		responseStruct.Sys.Country,
		responseStruct.Weather[0].Description,
		weatherEmoji,
		responseStruct.Main.Temp)
}

func main() {
	apiKey := flag.String("key", "", "Provide a valid https://openweathermap.org/api API key")
	flag.Parse()

	// Validating inputs presence. Present -h if required argument is not given.
	if *apiKey == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	input := strings.Join(flag.Args(), " ")
	message := HumanizedWeatherMessage(input, *apiKey)

	fmt.Println(message)
}
