# Hub Management Service

## Overview

The **Hub Management Service** is a backend service responsible for managing hubs, teams, and users. It provides various operations for creating, retrieving, and searching hubs, teams, and users within specific teams and hubs. The service is built in Go and uses Postgres (with support for SQLite for testing).

## Features

- **Hub Management**: Create and manage hubs.
- **Team Management**: Create and manage teams, linking them to hubs.
- **User Management**: Create users and associate them with teams.
- **Authentication**: Hardcoded authentication with JWT token support.
- **API Endpoints**: RESTful APIs for hubs, teams, and users.
- **Dockerized**: The service is set up with Docker Compose for easy local development.
- 
## Technologies
- **Go (Golang)**: Backend development language.
- **GORM**: Object-Relational Mapping (ORM) library for database interactions.
- **Postgres**: Database for storing hub, team, and user data.
- **SQLite**: In-memory database for testing.
- **Gin**: Web framework for building REST APIs.
- **JWT**: JSON Web Token for API authentication.
- **Docker**: Containerization for easy setup and deployment.

## Setup and Installation

### Prerequisites

1. **Docker**: Ensure Docker and Docker Compose are installed. You can download them from [here](https://www.docker.com/get-started).
2. **Go**: Ensure Go is installed on your system for development purposes.

### Steps to Run Locally with Docker Compose

1. **Clone the repository**:

   ```
   git clone https://github.com/yourusername/hub_management_service.git
   cd hub_management_service
   ```
   
2. **Run Docker Compose**:
  To start the app with Postgres, Swagger UI, and the application itself, run:

    ```
    docker-compose up --build
    ```

3.  **Verify the application**:
      Once the services are running, the application will be accessible. You can interact with the API using Swagger UI or Postman.
   
4. ** Run Unit Test**:
``` 
 cd hub_management_service 
 go test ./...
```
## Project Structure
    
    ```
    hub_management_service/
    ├── cmd/
    │   └── app/                # Main entry point for the application.
    ├── internal/
    │   ├── entity/             # Contains data models (Hub, Team, User).
    │   ├── repository/         # Interfaces and implementations for database operations.
    │   ├── service/            # Business logic and service interfaces.
    │   ├── handler/            # HTTP handlers for routing and request processing.
    │   ├── middleware/         # Middleware for logging, authentication, etc.
    │   └── router/             # Route definitions and API setup.
    ├── migrations/             # Database migration files (e.g., for Postgres).
    │── pkg/                    # Utility functions and shared components.
    ├── docs/                   # Documentation for the project.
    ├── .env                    # Environment variables for local development.
    ├── Dockerfile              # Dockerfile for building the application container.
    ├── docker-compose.yml      # Docker Compose file for setting up services.
    ├── go.mod                  # Go module file for managing dependencies.
    ```
## Explanation of Project Structure

### `cmd/app/`
This directory contains the entry point for the application. The `main.go` file initializes the server, sets up routing, and starts the application.

- **main.go**: This file is the main entry point of the application. It initializes the router, connects to the database, and starts the Gin HTTP server.

### `internal/`

This directory contains all the core logic of the application, divided into several components:

#### `entity/`

Contains data models that represent the application's entities. These include:

- **Hub**: Represents a hub entity.
- **Team**: Represents a team entity.
- **User**: Represents a user entity.

#### `repository/`

Contains interfaces and implementations for interacting with the database. These repositories encapsulate CRUD operations for each entity.

- **HubRepository**: Interface and implementation for CRUD operations related to hubs.
- **TeamRepository**: Interface and implementation for CRUD operations related to teams.
- **UserRepository**: Interface and implementation for CRUD operations related to users.

#### `service/`

Business logic layer, where the core functionality is defined. This is where data processing and complex operations occur.

- **HubService**: Service that handles business logic for hub-related operations.
- **TeamService**: Service that handles business logic for team-related operations.
- **UserService**: Service that handles business logic for user-related operations.

#### `handler/`

Contains the HTTP handler functions responsible for processing incoming HTTP requests and sending back appropriate responses. These handlers are used by the router to define endpoints.

- **HubHandler**: HTTP handler for operations related to hubs.
- **TeamHandler**: HTTP handler for operations related to teams.
- **UserHandler**: HTTP handler for operations related to users.

#### `middleware/`

Contains reusable middleware functions that can be applied to the router to handle common concerns, such as authentication, logging, and request validation.

- **AuthMiddleware**: Middleware for verifying JWT tokens and ensuring that the user is authorized to perform certain actions.

#### `router/`

Contains the route definitions and API setup. It connects HTTP requests to the appropriate handlers.

- **router.go**: Defines all the routes and connects them to their respective handlers.

### `migrations/`

Database migration files for setting up or altering the database schema. These files are used to manage database changes over time.

- **0001_initialize_table.sql**: creating tables for hubs, teams, and users.
- **0001_insert_sample_data.sql**: An example migration file for initializing the database records for hubs, teams, and users.


### `.env`

Contains environment variables for local development. It should include sensitive information, such as database credentials, JWT secret, and API keys.


## API Endpoints
### Swagger UI: http://localhost:8081/ 
We can access this swagger UI link to get all APIs information and try to make request.

Here is some APIs as example

### POST /auth/login
Get a JWT token (hardcoded authentication with admin:password).

#### Request
```
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{"username":"admin",
"password":"password"}'
```

#### Response
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM5MzY1NTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.Wh9KCqA39Er7v5TSohY4zxdkQM6GLRevkAkc638Aeng"
}
```

### POST /hubs 
Creates a new hub in the system.

#### Request
```
curl --location 'http://localhost:8080/hubs' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Tech Hub",
  "location": "Viet Nam"
}'
```

#### Response
```
{
    "hub": {
        "id": <id>,
        "name": "Tech Hub",
        "location": "Viet Nam"
    },
    "message": "Hub created successfully"
}
```


### GET /hubs/{id}
Retrieves a hub by its ID.

#### Request
```
curl -X 'GET' \
  'http://localhost:8080/hubs/1' \
  -H 'accept: application/json'
```

#### Response
```
{
  "hub": {
    "id": 1,
    "name": "Hub A",
    "location": "New York"
  }
}
```

### GET /hubs/{id}
Retrieves a hub by its ID.

#### Request
```
curl -X 'GET' \
  'http://localhost:8080/hubs/1' \
  -H 'accept: application/json'
```

#### Response
```
{
  "hub": {
    "id": 1,
    "name": "Hub A",
    "location": "New York"
  }
}
```


### GET /hubs/search?name=<name>
Searches for hubs by name in the system and returns the associated teams.

#### Request
```
curl -X 'GET' \
  'http://localhost:8080/hubs/search?name=Hub%20B' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer <token>'
```

#### Response
```
{
  "hubs": [
    {
      "id": 2,
      "name": "Hub B",
      "location": "San Francisco",
      "teams": [
        {
          "id": 2,
          "name": "Team Beta",
          "hub_id": 2
        }
      ]
    }
  ]
}
```


### POST /teams
Creates a new team in the system.

#### Request
```
curl -X 'POST' \
  'http://localhost:8080/teams' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Team Viet Nam",
  "hub_id": 1
}'
```

#### Response
```

  "message": "Team created successfully",
  "team": {
    "id": 140,
    "name": "Team Viet Nam",
    "hub_id": 1
  }
}
```


### POST /users
Creates a new user in the system.

#### Request
```
curl -X 'POST' \
  'http://localhost:8080/users' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Cuong Tran",
  "email": "cuongtran@gmail.com",
  "team_id": 1
}'
```

#### Response
```
{
  "message": "User created successfully",
  "user": {
    "id": 37,
    "name": "Cuong Tran",
    "team_id": 1,
    "email": "cuongtran@gmail.com"
  }
}
```

