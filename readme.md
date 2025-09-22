# Todo App

A full-stack Todo application built with Go (Gin + GORM) backend and Vue.js + TypeScript frontend.

## Features

- Create, read, update, and delete todo items
- RESTful API with Go/Gin
- MySQL database with GORM ORM
- Vue.js frontend with TypeScript

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16+)
- [MySQL](https://dev.mysql.com/downloads/) (version 5.7+)
- [Node.js](https://nodejs.org/) (version 14+)
- [npm](https://www.npmjs.com/) or [yarn](https://yarnpkg.com/)

### Clone the Repository

```bash
# Clone the repository
git clone https://github.com/Zhaobo-Wang/go-projects.git

# Navigate to the project directory
cd go-projects/os-go
```

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file cope `.env.sample`

4. Create the MySQL database:
   ```bash
   mysql -u root -p
   ```
   
   In the MySQL prompt:
   ```sql
   CREATE DATABASE todo_app;
   EXIT;
   ```

5. Run the backend server:
   ```bash
   go run main.go
   ```

   The server will start on http://localhost:8080

## API Endpoints

- `GET /api/v1/todos` - Get all todos
- `POST /api/v1/todos` - Create a new todo
- `GET /api/v1/todos/:id` - Get a single todo
- `PUT /api/v1/todos/:id` - Update a todo
- `DELETE /api/v1/todos/:id` - Delete a todo

## Built With

- [Go](https://golang.org/) - Backend language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Vue.js](https://vuejs.org/) - Frontend framework
- [TypeScript](https://www.typescriptlang.org/) - Type-safe JavaScript
- [MySQL](https://www.mysql.com/) - Database

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Vue.js Documentation](https://vuejs.org/guide/introduction.html)