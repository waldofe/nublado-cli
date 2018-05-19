package main

import (
	"os"
	"regexp"
	"testing"
)

// TODO: This should partially mock the request response
func TestIntegration(t *testing.T) {
	result := HumanizedWeatherMessage("sao paulo", os.Getenv("OPENWEATHER_API_KEY"))
	match, _ := regexp.MatchString(`It's (\d*\.?\d*)°C right now in Sao Paulo, BR!`, result)

	if !match {
		t.Error("Expected result to match with expected response")
	}
}
