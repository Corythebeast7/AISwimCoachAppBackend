# AI Swim Coach App Backend

A RESTful API backend for the AI Swim Coach application.

## Project Structure

```
AISwimCoachAppBackend/src/
├── database/       # All data handling and storage
├── rest/           # API route definitions and request handling
├── service/        # Actionable services that perform work
├── vo/             # Data models
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

## Future Enhancements

- Authentication and authorization
- Database integration
- Logging and monitoring
- API documentation with Swagger