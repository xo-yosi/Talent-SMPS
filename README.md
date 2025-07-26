
# Digital Meal Pass System

A simple backend system to track student meal attendance (breakfast, lunch, dinner) via student IDs. Built with Go, Gin, Gorm, and PostgreSQL, following the repository pattern for clean architecture.

---

## Features

- Register students with unique IDs  
- Query student info and meal status for the current day  
- Mark meals (breakfast, lunch, dinner) as taken, preventing duplicates  
- Simple RESTful API design  
- Clear separation of concerns using repository and service layers

---

## Tech Stack

- [Go](https://golang.org/)  
- [Gin](https://github.com/gin-gonic/gin) (HTTP web framework)  
- [Gorm](https://gorm.io/) (ORM for database interaction)  
- [PostgreSQL](https://www.postgresql.org/) (Relational database)  

---

## Project Structure

```
/cmd
/config
/handlers
/models
/repositories
/services
/database
/routes
/utils
```

---

## Getting Started

### Prerequisites

- Go 1.20+ installed  
- PostgreSQL database running  
- (Optional) Postman or curl for testing APIs

### Setup

1. Clone the repo:  
   ```bash
   git clone https://github.com/yourusername/meal-pass-system.git
   cd meal-pass-system
   ```

2. Configure database credentials in `config/config.go` or environment variables.

3. Run database migrations (if any).

4. Run the application:  
   ```bash
   go run ./cmd/main.go
   ```

5. The server will start at `http://localhost:8080`

---

## API Endpoints

| Method | Endpoint                  | Description                      |
|--------|---------------------------|--------------------------------|
| POST   | `/students`               | Register a new student          |
| GET    | `/students/:id`           | Get student info and meal status|
| POST   | `/students/:id/meals`     | Mark a meal as taken            |

---

## How to Use

1. Register a student with a unique ID.  
2. Query the student’s meal status anytime during the day.  
3. Mark meals as taken by submitting the student ID and meal type.

---

## Contributing

Feel free to open issues or submit pull requests. Keep it simple and focused on core features.

---

## License

MIT License © 2025 Your Name
