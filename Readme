# Simple CRUD Application - Go Backend

This repository contains a simple backend application built with Go. It demonstrates how to perform Create, Read, Update, and Delete (CRUD) operations using a SQL database. This is an ideal starting point for learning backend development with Go and SQL.

## Features

- CRUD operations for managing resources.
- SQL database integration for data persistence.
- Modular and clean Go codebase.
- Easy to set up and extend.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (v1.18 or later)
- A SQL database (e.g., MySQL, PostgreSQL, SQLite)
- [Git](https://git-scm.com/)

## Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd <repository-name>
```

### 2. Install Dependencies

This project uses Go modules. Run the following command to install required dependencies:

```bash
go mod tidy
```

### 3. Configure the Database

Update the `config.yaml` file (or environment variables) with your database connection details:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "your_username"
  password: "your_password"
  dbname: "your_database"
```

### 4. Run Database Migrations

Ensure your database schema is set up before running the application. Use the provided SQL migration script:

```bash
go run db/migrate.go
```

### 5. Run the Application

Start the server:

```bash
go run main.go
```

The server will start at `http://localhost:8080`.

## API Endpoints

### Resource: `users`

| Method | Endpoint      | Description       | Body (JSON)           |
|--------|---------------|-------------------|-----------------------|
| GET    | `/users`      | List all users    | N/A                   |
| POST   | `/users`      | Create a user     | `{ "name": "", "email": "" }` |
| GET    | `/users/{id}` | Get a specific user | N/A                 |
| PUT    | `/users/{id}` | Update a user     | `{ "name": "", "email": "" }` |
| DELETE | `/users/{id}` | Delete a user     | N/A                   |

## Project Structure

```
├── main.go           # Entry point of the application
├── db/
│   ├── migrate.go    # Database migration script
│   └── schema.sql    # SQL schema definition
├── handlers/         # HTTP handlers for each endpoint
├── models/           # Database models
├── config.yaml       # Configuration file
└── README.md         # Project documentation
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) for HTTP routing.
- [GORM](https://gorm.io/) or `database/sql` for SQL interaction.
