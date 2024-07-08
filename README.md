# go-gin-postgres

A RESTful API built with Go, Gin framework, and PostgreSQL.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Database Migrations](#database-migrations)
- [Running the Application](#running-the-application)
- [Project Structure](#project-structure)
- [API Routes](#api-routes)
- [Seeding Data](#seeding-data)
- [Contributing](#contributing)

## Introduction
This project is a sample RESTful API built using the Go programming language and the Gin framework. It uses PostgreSQL as the database and includes various features such as authentication, logging, and seeding.

## Features
- User authentication
- CRUD operations for users, tickets, orders, and payments
- Middleware for logging
- Database migrations and seeding

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/go-gin-postgres.git
    cd go-gin-postgres
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your PostgreSQL database and update the database configuration in `database/database.go`.

## Database Migrations
To apply the database migrations, use the following commands:

- Apply all migrations:
    ```sh
    migrate -path ./database/migrations -database 'postgres://user:password@localhost:5432/dbname?sslmode=disable' up
    ```

- Rollback the last migration:
    ```sh
    migrate -path ./database/migrations -database 'postgres://user:password@localhost:5432/dbname?sslmode=disable' down 1
    ```

## Running the Application
Start the application using the following command:
```sh
go run main.go
```

## Project Structure

```sh

go-gin-postgres/
├── .git/
├── auth/
│   └── auth.go
├── database/
│   ├── database.go
│   └── migrations/
│       ├── 000001_create_users_table.down.sql
│       ├── 000001_create_users_table.up.sql
│       ├── 000002_create_tickets_table.down.sql
│       ├── 000002_create_tickets_table.up.sql
│       ├── 000003_create_orders_table.down.sql
│       ├── 000003_create_orders_table.up.sql
│       ├── 000004_create_payments_table.down.sql
│       ├── 000004_create_payments_table.up.sql
├── handlers/
│   ├── authentication-handlers.go
│   ├── order-handlers.go
│   ├── payment-handlers.go
│   ├── ticket-handlers.go
│   ├── user-handlers.go
│   └── handlers.go
├── middleware/
│   └── logging.go
├── models/
│   └── models.go
├── seeder/
│   └── seed.go
├── app.log
├── go.mod
├── go.sum
└── main.go


```

## API Documentation

## Authentication

### Login
- `POST /login` - Authenticates a user and provides a JWT token

## User Routes

- `POST /users` - Create a new user
- `GET /users` - Get a list of all users
- `GET /users/:id` - Retrieve a user by their ID
- `PUT /users/:id` - Update a user by their ID
- `DELETE /users/:id` - Delete a user by their ID
- `GET /users/range/:start_id/:end_id` - Retrieve users within a range of IDs
- `GET /users/byname/:name` - Retrieve a user by their name

## Ticket Routes

- `GET /tickets/date/:start_date/:end_date` - Retrieve tickets within a date range
- `GET /tickets/date/time/:start_date/:end_date` - Retrieve tickets within a date and time range
- `GET /tickets/:user_id` - Retrieve tickets by user ID
- `GET /tickets/payment/:status` - Retrieve tickets by payment status
- `GET /records/date/:date_created` - Retrieve records by the ticket's date of creation
- `GET /records/:date/:start_time/:end_time` - Retrieve records within a specific date and time range

## Order Routes

- `GET /orders/date/:start_date/:end_date` - Retrieve orders within a date range

## Payment Routes

- `GET /payments/date/:start_date/:end_date` - Retrieve payments within a date range


## Seeding Data
To seed the database with initial data, use the provided seeder script.

### Prerequisites
Ensure your database is set up and configured.

### Running the Seeder
1. Run the seeder script:
    ```sh
    go run seeder/seed.go
    ```

### Seeder Script (`seeder/seed.go`)
Here is a brief overview of what the seeder script does:

- Connects to the PostgreSQL database.
- Inserts initial data for users, tickets, orders, and payments tables.

### Example Seed Data
The seed data might look something like this:

#### Users
```sql
INSERT INTO users (id, name, email, password) VALUES
(1, 'John Doe', 'john.doe@example.com', 'hashedpassword'),
(2, 'Jane Smith', 'jane.smith@example.com', 'hashedpassword');
```


## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.
