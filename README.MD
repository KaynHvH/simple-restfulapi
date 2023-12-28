# Very Simple RestfulAPI in Go

A very simple RESTful API implementation in Go (Golang) with SQLite database support.

## Features
- **GET /tasks:** Retrieve a list of tasks.
- **GET /tasks/{id}:** Retrieve a specific task by ID.
- **POST /tasks:** Create a new task.
- **PUT /tasks/{id}:** Update an existing task.
- **DELETE /tasks/{id}:** Delete a task.

## Usage
1. Ensure you have Go installed on your machine.
2. Clone the repository:
   ```bash
   git clone https://github.com/KaynHvH/simple-restfulapi.git
   ```
3. Navigate to the project directory:
   ```bash
   cd simple-restfulapi
   ```
4. Initialize the SQLite database:
   ```bash
   go run main.go
   ```
5. The API will be accessible at [http://localhost:8080](http://localhost:8080).

## Endpoints
- **GET /tasks:** Retrieve all tasks.
- **GET /tasks/{id}:** Retrieve a specific task.
- **POST /tasks:** Create a new task.
  - Request Body:
    ```json
    {
      "title": "Task Title"
    }
    ```
- **PUT /tasks/{id}:** Update an existing task.
  - Request Body:
    ```json
    {
      "title": "Updated Task Title"
    }
    ```
- **DELETE /tasks/{id}:** Delete a task.

## Dependencies
- [gorilla/mux](https://github.com/gorilla/mux): HTTP request router and dispatcher.
- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3): SQLite3 driver for Go's database/sql package.

## Database
- SQLite3 is used as the backend database.
- Database file: `tasks.db`

## Project Structure
- **main.go:** Application entry point.
- **handlers:** Package containing HTTP request handlers.
- **db:** Package handling database initialization and queries.
- **models:** Package defining the data model.

## Acknowledgments
This project utilizes the power of Go and several open-source packages to create a simple and efficient RESTful API.

Feel free to contribute or modify as needed!
