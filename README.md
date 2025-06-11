# Weathercast Backend

A Go-based backend service that aggregates weather news from multiple Indian news sources including Times of India and Indian Express.

## Features

-   Aggregates weather news from multiple sources
-   RESTful API endpoints
-   Fast and efficient news fetching
-   Built with Go and Gin framework

## Prerequisites

-   Go 1.19 or higher
-   Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/weathercast-backend.git
cd weathercast-backend
```

2. Install dependencies:

```bash
go mod download
```

## Usage

To start the server:

```bash
go run main.go
```

The server will start on port 8080 by default.

## API Endpoints

### Get Weather News

```
GET /api/weather
```

Returns aggregated weather news from all configured sources.

Response format:

```json
{
    "news": [
        {
            "title": "Weather news title",
            "description": "Weather news description",
            "source": "Source name",
            "timestamp": "2025-06-11T12:00:00Z"
        }
    ]
}
```

## Project Structure

```
├── main.go           # Entry point of the application
├── handlers/         # HTTP request handlers
├── routes/          # Route definitions
├── ie/              # Indian Express news scraper
├── toi/             # Times of India news scraper
└── utils/           # Utility functions
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

-   Times of India for weather news data
-   Indian Express for weather news data
