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



## API Routes
Here are the API routes provided by this application:

### User Routes
- `POST /users` - Create a new user
- `GET /users` - Get a list of all users
- `GET /users/:id` - Get details of a specific user
- `PUT /users/:id` - Update a specific user
- `DELETE /users/:id` - Delete a specific user

### Ticket Routes
- `POST /tickets` - Create a new ticket
- `GET /tickets` - Get a list of all tickets
- `GET /tickets/:id` - Get details of a specific ticket
- `PUT /tickets/:id` - Update a specific ticket
- `DELETE /tickets/:id` - Delete a specific ticket

### Order Routes
- `POST /orders` - Create a new order
- `GET /orders` - Get a list of all orders
- `GET /orders/:id` - Get details of a specific order
- `PUT /orders/:id` - Update a specific order
- `DELETE /orders/:id` - Delete a specific order

### Payment Routes
- `POST /payments` - Create a new payment
- `GET /payments` - Get a list of all payments
- `GET /payments/:id` - Get details of a specific payment
- `PUT /payments/:id` - Update a specific payment
- `DELETE /payments/:id` - Delete a specific payment

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
