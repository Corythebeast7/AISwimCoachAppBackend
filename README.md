# AI Swim Coach App Backend

A RESTful API backend for the AI Swim Coach application.

## Project Structure

```
AISwimCoachAppBackend/
├── handlers/       # Request handlers
├── models/         # Data models
├── routes/         # API route definitions
├── go.mod          # Go module file
├── main.go         # Application entry point
└── README.md       # This file
```

## Getting Started

### Prerequisites

- Go 1.24 or later

### Running the Application

1. Clone the repository
2. Navigate to the project directory
3. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

- `GET /api/health` - Health check endpoint to verify the API is running

## Adding New Endpoints

To add new endpoints to the API:

1. Create handler functions in the `handlers` package
2. Define data models in the `models` package if needed
3. Add routes in the `routes` package

## Future Enhancements

- Authentication and authorization
- Database integration
- Logging and monitoring
- API documentation with Swagger