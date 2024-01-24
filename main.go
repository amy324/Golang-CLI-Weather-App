package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

const weatherAPIURL = "https://api.openweathermap.org/data/2.5/weather"
const ipLocationAPIURL = "https://ipapi.co/json"


// WeatherData represents the structure of the JSON response from the weather API
type WeatherData struct {
	Coord      Coord      `json:"coord"`
	Weather    []Weather  `json:"weather"`
	Base       string     `json:"base"`
	Main       Main       `json:"main"`
	Visibility int        `json:"visibility"`
	Wind       Wind       `json:"wind"`
	Clouds     Clouds     `json:"clouds"`
	Dt         int        `json:"dt"`
	Sys        Sys        `json:"sys"`
	Timezone   int        `json:"timezone"`
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Cod        int        `json:"cod"`
}

// Coord represents the coordinates of the location
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Weather represents weather information
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main represents main weather information
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// Wind represents wind information
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

// Clouds represents cloud information
type Clouds struct {
	All int `json:"all"`
}

// Sys represents system information
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Define command-line flags
	locationFlag := flag.String("location", "", "Specify the location for weather information")
	flag.Parse()

	// Retrieve the API key from environment variable
	apiKey, apiKeyExists := os.LookupEnv("OPENWEATHERMAP_API_KEY")

	// Check if the API key is set
	if !apiKeyExists {
		fmt.Println("OpenWeatherMap API key not set. Exiting.")
		os.Exit(1)
	}

	// Display menu for the user to choose an option
	color.Green("Choose an option:")
	color.Cyan("1. Get weather for your location")
	color.Cyan("2. Enter a specific location")

	var option int
	color.Green("Enter your choice (1 or 2): ")
	_, readErr := fmt.Scan(&option)
	if readErr != nil {
		log.Fatal("Error reading input:", readErr)
	}

	switch option {
	case 1:
		// Get weather for the user's location
		userLocation, locErr := getUserLocation()
		if locErr != nil {
			log.Fatal("Error getting user location:", locErr)
		}
		fetchAndDisplayWeather(userLocation, apiKey)
	case 2:
		// Get weather for a specific location
		if *locationFlag == "" {
			color.Green("Enter the location: ")
			fmt.Scan(locationFlag)
		}

		// Check if the location is not empty
		if *locationFlag != "" {
			fetchAndDisplayWeather(*locationFlag, apiKey)
		} else {
			fmt.Println("Invalid location. Exiting.")
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid option. Exiting.")
		os.Exit(1)
	}
}


func getUserLocation() (string, error) {
	// Make a request to get user's approximate geolocation based on IP
	resp, err := http.Get(ipLocationAPIURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for errors in the API response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching user location. Please enter a location manually.")
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var locationData map[string]interface{}
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return "", err
	}

	// Extract city from the location data
	city, ok := locationData["city"].(string)
	if !ok {
		return "", fmt.Errorf("City not found in user location data")
	}

	return city, nil
}

func fetchAndDisplayWeather(location, apiKey string) {
	// Fetch weather data from the API
	weatherData, err := getWeatherData(location, apiKey)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}

	// Display weather information
	displayWeatherInfo(weatherData)
}


func getWeatherData(location, apiKey string) (*WeatherData, error) {
	// Construct the API URL with query parameters
	apiURL := fmt.Sprintf("%s?q=%s&appid=%s", weatherAPIURL, url.QueryEscape(location), apiKey)

	// Make a GET request to the weather API
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for errors in the API response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status code", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
func displayWeatherInfo(weatherData *WeatherData) {
	// Display relevant weather information
	color.Cyan("Weather in %s, %s:\n", weatherData.Name, weatherData.Sys.Country)
	color.Cyan("Description: %s\n", weatherData.Weather[0].Description)
	color.Cyan("Temperature: %.2f°C\n", weatherData.Main.Temp-273.15)  // Converts Kelvin to Celsius
	color.Cyan("Feels Like: %.2f°C\n", weatherData.Main.FeelsLike-273.15) 
	color.Cyan("Min Temperature: %.2f°C\n", weatherData.Main.TempMin-273.15) 
	color.Cyan("Max Temperature: %.2f°C\n", weatherData.Main.TempMax-273.15) 
	color.Cyan("Pressure: %d hPa\n", weatherData.Main.Pressure)
	color.Cyan("Humidity: %d%%\n", weatherData.Main.Humidity)
	color.Cyan("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
	color.Cyan("Cloudiness: %d%%\n", weatherData.Clouds.All)
}


