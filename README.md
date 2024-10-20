# Weather API Service

This project is a weather API service built using Golang. The service fetches weather data from a third-party API (e.g., Visual Crossing API) and implements caching using Redis with a 12-hour expiration. The API allows clients to retrieve current weather information for a specified city and improves performance by caching frequently requested data.

## Table of Contents

- [Features](#features)
- [System Architecture](#system-architecture)
- [Installation and Setup](#installation-and-setup)
- [API Endpoints](#api-endpoints)
- [API Tests](#api-tests)
- [Usage](#usage)
- [System Diagrams](#system-diagrams)
- [Contributing](#contributing)
- [License](#license)

## Features

- Retrieves real-time weather data from a third-party API
- Implements Redis-based caching for performance optimization
- Custom date format handling for proper time unmarshalling
- Graceful error handling for external API failures
- Extensible caching layer for future optimizations

## System Architecture

The system architecture is designed as follows:

1. **WeatherService**:
   - Responsible for fetching weather data from the third-party API.
   - Implements caching using Redis to reduce API calls.
   - Handles unmarshalling of JSON responses, including custom date formats.

2. **Redis Cache**:
   - Caches weather data for each city with a 12-hour expiration to optimize repeated requests.

3. **Third-Party Weather API**:
   - This service fetches weather data from an external API like Visual Crossing.

4. **Error Handling**:
   - If the data is unavailable or if the external API returns an error, the service will return an appropriate message.

## Installation and Setup

### Prerequisites

- Golang (v1.18 or later)
- Redis for caching
- Docker (optional but recommended for running Redis and tests)
- Visual Crossing API key (or any compatible weather API)

### Step-by-Step Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/weather-api.git
   cd weather-api
   ```

2. **Set Up Environment Variables**:
   You need to set the following environment variables for the API key and Redis connection:
   ```bash
   export WEATHER_API_KEY="your-weather-api-key"
   export REDIS_URL="redis://localhost:6379"
   ```

3. **Install Dependencies**:
   If you use Go modules, ensure you install all the required dependencies:
   ```bash
   go mod tidy
   ```

4. **Run the Service**:
   To start the weather service, run the following command:
   ```bash
   go run main.go
   ```

5. **Run Redis**:
   If you're using Docker, run Redis with:
   ```bash
   docker run --name redis -d -p 6379:6379 redis
   ```

## API Endpoints

### `GET /weather/{city}`

Fetches the weather data for a specified city. If the weather data is cached, it will return the cached response, otherwise, it will fetch fresh data from the third-party API.

- **URL**: `/weather/{city}`
- **Method**: `GET`
- **URL Params**: 
   - `city=[string]` - Name of the city
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: JSON containing weather data.
  
  Example Response:
  ```json
  {
    "queryCost": 1,
    "latitude": 37.7749,
    "longitude": -122.4194,
    "resolvedAddress": "San Francisco, CA",
    "days": [
      {
        "datetime": "2024-10-20",
        "tempmax": 21.2,
        "tempmin": 13.5,
        "temp": 17.8,
        "humidity": 68.0,
        "precipprob": 5.0
      }
    ]
  }
  ```
  
- **Error Response**:
  - **Code**: 400 Bad Request
  - **Message**: `City is required`
  
  - **Code**: 500 Internal Server Error
  - **Message**: `Failed to fetch weather data`

## API Tests

You can run API tests using `curl` or an API testing tool like Postman.

### Example CURL Requests

1. **Get Weather Data for a City**:
   ```bash
   curl -X GET "http://localhost:8080/weather/London"
   ```

2. **Handling Errors**:
   If the city name is missing:
   ```bash
   curl -X GET "http://localhost:8080/weather/"
   ```

You can also use Postman to simulate the requests and validate responses.

## Usage

Once the service is running, you can query the weather data for a specific city by sending an HTTP `GET` request to the `/weather/{city}` endpoint. For example:

```bash
curl http://localhost:8080/weather/London
```

The service will first check if the data is available in the Redis cache. If not, it will fetch the data from the third-party weather API and store it in Redis with a 12-hour expiration.

## System Diagrams
![System Diagram](./system_diagram.png)

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any bugs or want to request new features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
