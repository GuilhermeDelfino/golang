package repository

import (
	"database/sql"

	"github.com/GuilhermeDelfino/golang/class3/internal/entity"
)

type ProductRepository interface {
	FindAll() (*[]entity.Product, error)
	Save(p *entity.Product) error
	FindById(id *int) (*entity.Product, error)
}
type productRepository struct {
	conn *sql.DB
}

func CreateProductRepository(conn *sql.DB) productRepository {
	return productRepository{conn}
}

func (r productRepository) FindAll() (*[]entity.Product, error) {
	query := "SELECT id, title, description, price FROM products"
	rows, err := r.conn.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product
		rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price)
		products = append(products, p)
	}

	return &products, nil
}

func (r productRepository) Save(p *entity.Product) error {
	query := "INSERT INTO products (title, description, price) VALUES (?,?,?)"
	result, err := r.conn.Exec(query, p.Title, p.Description, p.Price)

	if err != nil {
		return err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id64)

	return nil
}

func (r productRepository) FindById(id *int) (*entity.Product, error) {
	query := "SELECT id, title, description, price FROM products WHERE id = ?"
	row, err := r.conn.Query(query, &id)

	if err != nil {
		return nil, err
	}

	var product entity.Product
	if row.Next() {
		row.Scan(&product.ID, &product.Title, &product.Description, &product.Price)
	}

	return &product, nil
}
