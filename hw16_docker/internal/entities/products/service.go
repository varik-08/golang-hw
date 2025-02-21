package products

import (
	"github.com/varik-08/golang-hw/hw16_docker/config"
)

func GetProducts() ([]*Product, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	rows, err := dbPool.Query(ctx, "SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*Product, 0, 100)

	for rows.Next() {
		var i Product
		if err = rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		products = append(products, &i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(product ProductDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	row := dbPool.QueryRow(ctx, "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id",
		product.Name, product.Price)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateProduct(product ProductDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	row := dbPool.QueryRow(ctx, "UPDATE products SET price = $1 WHERE id = $2 RETURNING id",
		product.Price, product.ID)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteProduct(product ProductDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	row := dbPool.QueryRow(ctx, "DELETE FROM products WHERE id = $1 RETURNING id",
		product.ID)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
