package main

import (
    "employees/controller"
    "employees/repository"
    "employees/routes"
    "employees/service"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

func main() {
    // Create a new Fiber app instance
    app := fiber.New()

    // Configure CORS middleware with allowed origins from environment variables
    app.Use(cors.New(cors.Config{
        AllowOrigins: os.Getenv("ALLOWED_ORIGINS"),
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Initialize database connection
    db := initializeDatabaseConnection()

    // Run migrations to ensure database schema is up to date
    repository.RunMigrations(db)

    // Initialize repositories, services, and controllers
    employeeRepository := repository.NewEmployeeRepository(db)
    employeeService := service.NewEmployeeService(employeeRepository)
    employeeController := controller.NewEmployeeController(employeeService)

    // Register routes for the application
    routes.RegisterRoute(app, employeeController)

    // Start the server on port 8080
    err := app.Listen(":8080")
    if err != nil {
        log.Fatalln(fmt.Sprintf("error starting the server: %s", err.Error()))
    }
}

// initializeDatabaseConnection sets up the database connection using GORM
func initializeDatabaseConnection() *gorm.DB {
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  createDsn(), // Create the DSN for the database connection
        PreferSimpleProtocol: true,         // Use simple protocol (avoid using psql's features)
    }), &gorm.Config{})
    if err != nil {
        log.Fatalln(fmt.Sprintf("error connecting to database: %s", err.Error()))
    }
    return db
}

// createDsn constructs the Data Source Name (DSN) for the PostgreSQL connection
func createDsn() string {
    dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
    dbHost := os.Getenv("DB_HOST")        // Get database host from environment variable
    dbUser := os.Getenv("DB_USER")        // Get database user from environment variable
    dbPassword := os.Getenv("DB_PASSWORD") // Get database password from environment variable
    dbName := os.Getenv("DB_NAME")        // Get database name from environment variable
    dbPort := os.Getenv("DB_PORT")        // Get database port from environment variable
    return fmt.Sprintf(dsnFormat, dbHost, dbUser, dbPassword, dbName, dbPort) // Return formatted DSN string
}

