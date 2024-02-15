package migrations

import (
	"database/sql"
	"github.com/dgryski/trifles/uuid"
	"github.com/liberopassadorneto/clean-arch/internal/entity"
	"log"
)

func CreateOrdersTable(db *sql.DB) {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS orders (
        id VARCHAR(255) NOT NULL,
        price DOUBLE,
        tax DOUBLE,
        final_price DOUBLE,
        PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create orders table: ", err)
	}

	log.Println("orders table created successfully")
}

func InsertSampleData(db *sql.DB) {
	orders := []entity.Order{
		{uuid.UUIDv4(), 100.0, 20.0, 120.0},
		{uuid.UUIDv4(), 200.0, 40.0, 240.0},
		{uuid.UUIDv4(), 300.0, 60.0, 360.0},
		{uuid.UUIDv4(), 400.0, 80.0, 480.0},
		{uuid.UUIDv4(), 500.0, 100.0, 600.0},
		{uuid.UUIDv4(), 600.0, 120.0, 720.0},
		{uuid.UUIDv4(), 700.0, 140.0, 840.0},
		{uuid.UUIDv4(), 800.0, 160.0, 960.0},
		{uuid.UUIDv4(), 900.0, 180.0, 1080.0},
		{uuid.UUIDv4(), 1000.0, 200.0, 1200.0},
		// Add more orders as needed
	}

	for _, order := range orders {
		_, err := db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)
		if err != nil {
			log.Fatal("Failed to insert order: ", err)
		}
	}

	log.Println("Sample data inserted successfully")
}
