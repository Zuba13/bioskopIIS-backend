package main

import (
	"log"
	"net/http"

	"bioskop.com/projekat/bioskopIIS-backend/handler"
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
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
	database.AutoMigrate(&model.Hall{})
	database.AutoMigrate(&model.Ticket{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Contract{})
	database.AutoMigrate(&model.ContractItem{})
	database.AutoMigrate(&model.StockItem{})
	database.AutoMigrate(&model.Supplier{})
	database.AutoMigrate(&model.MenuItem{})
	database.AutoMigrate(&model.Menu{})
	database.AutoMigrate(&model.SupplierProduct{})
	database.AutoMigrate(&model.OrderSuggestion{})
	return database
}

func startServer(router *mux.Router) {
	// Create a new CORS middleware with desired options
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Use the CORS middleware to wrap the router
	handler := corsMiddleware.Handler(router)

	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", handler))
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

	hallRepo := &repo.HallRepository{DatabaseConnection: database}
	hallService := &service.HallService{HallRepository: *hallRepo}
	hallHandler := &handler.HallHandler{HallService: *hallService}

	ticketRepo := &repo.TicketRepository{DB: database}
	ticketService := &service.TicketService{TicketRepository: *ticketRepo}
	ticketHandler := &handler.TicketHandler{TicketService: *ticketService, UserService: *userService, ProjectionService: *projectionService}

	productRepo := &repo.ProductRepository{DatabaseConnection: database}
	productService := &service.ProductService{ProductRepository: *productRepo}
	productHandler := &handler.ProductHandler{ProductService: *productService}

	contractRepo := &repo.ContractRepository{DatabaseConnection: database}
	contractService := &service.ContractService{ContractRepository: *contractRepo}
	contractHandler := &handler.ContractHandler{ContractService: *contractService}

	stockItemRepo := &repo.StockItemRepository{DatabaseConnection: database}
	stockItemService := &service.StockItemService{StockItemRepository: *stockItemRepo, ContractService: *contractService}
	stockItemHandler := &handler.StockItemHandler{StockItemService: *stockItemService}

	menuItemRepo := &repo.MenuItemRepository{DatabaseConnection: database}
	menuItemService := &service.MenuItemService{MenuItemRepository: *menuItemRepo}
	menuItemHandler := &handler.MenuItemHandler{MenuItemService: *menuItemService}

	menuRepo := &repo.MenuRepository{DatabaseConnection: database}
	menuService := &service.MenuService{MenuRepository: *menuRepo}
	menuHandler := &handler.MenuHandler{MenuService: *menuService}



	contractItemRepo := &repo.ContractItemRepository{DatabaseConnection: database}
	contractItemService := &service.ContractItemService{ContractItemRepository: *contractItemRepo}
	contractItemHandler := &handler.ContractItemHandler{ContractItemService: *contractItemService}
	
	supplierRepo := &repo.SupplierRepository{DatabaseConnection: database}
	supplierService := &service.SupplierService{SupplierRepository: *supplierRepo}
	supplierHandler := &handler.SupplierHandler{SupplierService: *supplierService}


	// Create a new router
	router := mux.NewRouter().StrictSlash(true)

	// Register your handlers with the router
	userHandler.RegisterUserHandler(router)
	movieHandler.RegisterMovieHandler(router)
	projectionHandler.RegisterProjectionHandler(router)
	hallHandler.RegisterHallHandler(router)
	ticketHandler.RegisterTicketHandler(router)
	menuItemHandler.RegisterMenuItemHandler(router)
	stockItemHandler.RegisterStockItemHandler(router)
	menuHandler.RegisterMenuHandler(router)
	productHandler.RegisterProductHandler(router)
	contractHandler.RegisterContractHandler(router)
	contractItemHandler.RegisterContractItemHandler(router)
	supplierHandler.RegisterSupplierHandler(router)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	c := cron.New()
	// Scheduled to run every 3 minutes for testing
	//_, err := c.AddFunc("/3 * * * *", func() { stockItemService.DailyTaskForDelivery() })
	// Schedule the task to run every day at 12:00
	//_, err := c.AddFunc("0 12 * * *", func() { stockItemService.DailyTaskForDelivery() })
	// if err != nil {
	// 	fmt.Println("Error scheduling task:", err)
	// 	return
	// }

	c.Start()
	
	//go stockItemService.DailyTaskForDelivery()

	go startServer(router)

	// Start the server with the CORS-enabled handler
	//log.Println("Server starting")
	//log.Fatal(http.ListenAndServe(":8085", handler))
	

	// Wait indefinitely
	select {}

}




func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
