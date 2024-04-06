package main

import (
	"log"
	"net/http"

	"bioskop.com/projekat/bioskopIIS-backend/handler"
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=cinema port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	// Auto-migrate models for the User and Movie modules
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Movie{})
	database.AutoMigrate(&model.Projection{})

	return database
}

func startServer(userHandler *handler.UserHandler, movieHandler *handler.MovieHandler) {
	router := mux.NewRouter().StrictSlash(true)

	userHandler.RegisterUserHandler(router)
	movieHandler.RegisterMovieHandler(router)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", router))
}

func main() {
	database := initDB()
	if database == nil {
		log.Println("Failed to connect to database!")
		return
	}

	// Initialize User repository, service, and handler
	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepo: *userRepo}
	userHandler := &handler.UserHandler{UserService: *userService}

	// Initialize Movie repository, service, and handler
	movieRepo := &repo.MovieRepository{DatabaseConnection: database}
	movieService := &service.MovieService{MovieRepository: *movieRepo}
	movieHandler := &handler.MovieHandler{MovieService: *movieService}

	projectionRepo := &repo.ProjectionRepository{DatabaseConnection: database}
	projectionService := &service.ProjectionService{ProjectionRepo: *projectionRepo}
	projectionHandler := &handler.ProjectionHandler{ProjectionService: *projectionService}

	// Create a new router
	router := mux.NewRouter().StrictSlash(true)

	// Register your handlers with the router
	userHandler.RegisterUserHandler(router)
	movieHandler.RegisterMovieHandler(router)
	projectionHandler.RegisterProjectionHandler(router)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Create a new CORS middleware with desired options
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Use the CORS middleware to wrap the router
	handler := corsMiddleware.Handler(router)

	// Start the server with the CORS-enabled handler
	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", handler))
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
