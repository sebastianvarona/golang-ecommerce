package database

import (
	"context"
	"database/sql"
	"ecommerce_template/internal/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// Services
	UserInsert(user *models.User) (int, error)

	// Category methods
	CreateCategory(category *models.Category) error
	GetCategoryByID(id int) (*models.Category, error)
	GetAllCategories() ([]*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id int) error

	// CategoryTranslation methods
	CreateCategoryTranslation(ct *models.CategoryTranslation) error
	GetCategoryTranslationByID(id int) (*models.CategoryTranslation, error)
	GetAllCategoryTranslations() ([]*models.CategoryTranslation, error)
	UpdateCategoryTranslation(ct *models.CategoryTranslation) error
	DeleteCategoryTranslation(id int) error

	// Product methods
	CreateProduct(product *models.Product) error
	GetProductByID(id int) (*models.Product, error)
	GetAllProducts() ([]*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error

	// ProductTranslation methods
	CreateProductTranslation(pt *models.ProductTranslation) error
	GetProductTranslationByID(id int) (*models.ProductTranslation, error)
	GetAllProductTranslations() ([]*models.ProductTranslation, error)
	UpdateProductTranslation(pt *models.ProductTranslation) error
	DeleteProductTranslation(id int) error

	// Variant methods
	CreateVariant(variant *models.Variant) error
	GetVariantByID(id int) (*models.Variant, error)
	GetAllVariants() ([]*models.Variant, error)
	UpdateVariant(variant *models.Variant) error
	DeleteVariant(id int) error

	// VariantTranslation methods
	CreateVariantTranslation(vt *models.VariantTranslation) error
	GetVariantTranslationByID(id int) (*models.VariantTranslation, error)
	GetAllVariantTranslations() ([]*models.VariantTranslation, error)
	UpdateVariantTranslation(vt *models.VariantTranslation) error
	DeleteVariantTranslation(id int) error
}

type service struct {
	db *sql.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func NewDB() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	dbInstance = &service{
		db: NewDB(),
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
