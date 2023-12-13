package storage

import (
	"database/sql"
	"fmt"
)

const (
	CreateTableDog = `CREATE TABLE IF NOT EXISTS dogs(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		description TEXT NOT NULL,
		weight FLOAT(3,2) NOT NULL,
		max_age INT NOT NULL,
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		update_at TIMESTAMP
		)ENGINE = InnoDB;
		`
)

type MySQLDog struct {
	db *sql.DB
}

func NewMySQLDog(db *sql.DB) *MySQLDog {
	return &MySQLDog{db}
}

func (m *MySQLDog) Migrate() error {
	stm, err := m.db.Prepare(CreateTableDog)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migrate of Dog was a successfully")
	return nil
}
