
# Firebase Push Notification Microservice

This microservice allows you to send push notifications to specific device tokens using Firebase Cloud Messaging (FCM). It is built using **Go** with the **Gin** web framework and supports **Swagger** documentation.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [License](#license)

## Features
- Send push notifications using Firebase Cloud Messaging (FCM)
- Input multiple tokens and send personalized data, titles, and bodies
- Swagger documentation for API
- Custom 404 handling
- JSON response formatting

## Requirements
- Go 1.20+
- Firebase Admin SDK JSON key (credentials file)
- Gin framework
- Slog for logging
- Swag for Swagger documentation

## Installation

### Run from Docker Hub

```sh 
docker run \
  -u "$UID" \
  -p 8080:8080 \
  -v $(pwd)/keys.json:/app/keys.json \
  -e MODE=production \
  -e PORT=8080 \
  owles/fbpushd:latest
```

### Build from source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/fbpushd.git
   ```

2. Navigate to the project directory:
   ```bash
   cd fbpushd
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Install Swag to generate Swagger documentation:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

5. Generate Swagger documentation:
   ```bash
   swag init
   ```

## Configuration

You need to provide your Firebase Admin SDK credentials to send push notifications. Add your `keys.json` file to the root of the project.

## Running the Application

1. Set the `PORT` environment variable or default to `8080`:
   ```bash
   export PORT=8080
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. The application will be available at:
   ```
   http://localhost:8080
   ```

## API Documentation

The API documentation is available via Swagger UI:

1. Access the Swagger documentation at:
   ```
   http://localhost:8080/swagger/index.html
   ```

2. Swagger will display detailed information about the available endpoints, request/response formats, and error codes.

## Project Structure

```
fbpushd/
├── main.go                  # Entry point of the application
├── go.mod                   # Go module definition
├── docs/                    # Swagger documentation (generated via swag init)
├── keys.json                # Firebase credentials (not in version control)
└── internal/
    ├── api/
    │   └── web.go           # API route and handler definitions
    └── push/
        ├── config.go        # Firebase configuration and setup
        └── notify.go        # Notification sending logic
```

## Endpoints

### POST `/push`

- **Description**: Send push notifications to specified device tokens using Firebase.
- **Request Body**:
  ```json
  [
    {
      "token": "your-firebase-token",
      "push": {
        "data": {
          "key": "value"
        },
        "title": "Notification Title",
        "body": "Notification Body"
      }
    }
  ]
  ```
- **Response**:
  ```json
  [
    {
      "token": "your-firebase-token",
      "success": true
    }
  ]
  ```

### Custom 404 Route

- **Response**:
  ```json
  {
    "result": "error",
    "message": "invalid API method"
  }
  ```

## License

This project is licensed under the MIT License.
