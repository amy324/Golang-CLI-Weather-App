

# Golang Weather CLI App

A simple command-line tool to retrieve weather information.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)


## Overview

This CLI app is a command-line tool that allows users to quickly fetch and display weather information for a given location.

## Features

* Retrieve current weather information based on location.
* Supports both automatic detection of the user's location and manual entry.

## Getting Started

### Prerequisites

* Be sure to have G9 installed on your machine
* OpenWeatherMap API key (instructions on obtaining one are [here](https://openweathermap.org/appid))
* Install any dependencies 

### Installation

1\. Clone the repository:

```bash
https://github.com/amy324/Golang-CLI-Weather-App
cd Golang-CLI-Weather-App
```

2\. Set up your OpenWeatherMap API key:

Create a .env file, and format your API key as `YOUR_API_KEY=your actual API key`

3\. Build and install the CLI tool:

```bash\
go build -o Golang-CLI-Weather-App  main.go
```

## Usage

To get weather information for your location or a city of your choice, run:

```bash\
.go run main.go
```
or alternatively
```bash\
.go run .
```

### Terminal Output Example

Output after running program:

```
Choose an option:
1. Get weather for your location
2. Enter a specific location
Enter your choice (1 or 2): 
```
Example output after selecting option `1`,  "Your City" will be replaced with your actual city after obtaining your geolocation:

```
Your City
Weather in Your City:
Description: overcast clouds
Temperature: 10.27°C
Feels Like: 9.23°C
Min Temperature: 9.20°C
Max Temperature: 10.99°C
Pressure: 1027 hPa
Humidity: 72%
Wind Speed: 5.66 m/s
Cloudiness: 100% 

```
Example output after selecting option `2`,  and searching for a city, in this case Paris:

```
Choose an option:
1. Get weather for your location
2. Enter a specific location
Enter your choice (1 or 2): 
2
Enter the location: 
Paris
Weather in Paris, FR:
Description: clear sky
Temperature: 12.71°C
Feels Like: 12.05°C
Min Temperature: 11.11°C
Max Temperature: 13.47°C
Pressure: 1031 hPa
Humidity: 77%
Wind Speed: 3.09 m/s
Cloudiness: 0%

```
## Code Explanation

### `func main()`

- The `main` function initializes the application, loads environment variables, and presents a menu to the user.
- It uses command-line flags to allow the user to specify a location.
- The user's location is obtained through the `getUserLocation` function, which makes a request to the ipapi.co API based on the user's IP address.
- The `fetchAndDisplayWeather` function is responsible for fetching and displaying weather information based on the user's choice.

### `getWeatherData`, `getWeatherData`, `displayWeatherInfo`

- The `getWeatherData` function constructs the OpenWeatherMap API URL, makes a GET request, and parses the JSON response into the `WeatherData` struct.
- The `displayWeatherInfo` function takes a `WeatherData` struct and prints relevant weather information to the console.

### `types`

- Defines the data structures (`WeatherData`, `Coord`, `Weather`, `Main`, `Wind`, `Clouds`, `Sys`) used to represent the JSON response from the OpenWeatherMap API.

## Configuration

The CLI tool uses environment variables for configuration. You can set your OpenWeatherMap API key in an `.env` file.

## Contributing
Feel free to contribute to the project by opening issues, providing feedback, or submitting pull requests. Your input is valuable!

