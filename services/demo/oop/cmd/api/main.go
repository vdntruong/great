package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"oop/internal/api"
	"oop/internal/repositories"
	"oop/internal/services"
)

// EmailNotifier implements OrderNotifier interface for email notifications
type EmailNotifier struct {
	// Email configuration would go here
}

func (n *EmailNotifier) NotifyOrderCreated(order *domain.Order) error {
	// Implementation to send email notification
	return nil
}

func (n *EmailNotifier) NotifyOrderStatusChanged(order *domain.Order, oldStatus domain.OrderStatus) error {
	// Implementation to send email notification
	return nil
}

func main() {
	// Connect to database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/ecommerce")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	productRepo := repositories.NewPostgresProductRepository(db)
	orderRepo := repositories.NewPostgresOrderRepository(db)

	// Initialize notifier
	notifier := &EmailNotifier{}

	// Initialize services
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, productRepo, notifier)

	// Initialize handlers
	productHandler := api.NewProductHandler(productService)
	orderHandler := api.NewOrderHandler(orderService)

	// Initialize router
	router := api.NewRouter(productHandler, orderHandler)

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router.SetupRoutes()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
