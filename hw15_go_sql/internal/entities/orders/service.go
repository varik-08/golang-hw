package orders

import (
	"github.com/jackc/pgx/v5"
	"github.com/varik-08/golang-hw/hw15_go_sql/config"
)

func GetOrdersByUser(dto OrderDTO) ([]*Order, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	rows, err := dbPool.Query(ctx, `
		SELECT id, order_date, total_amount
		FROM orders
		WHERE user_id = $1
	`, dto.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*Order, 0, 100)

	for rows.Next() {
		order := new(Order)
		if err = rows.Scan(&order.ID, &order.OrderDate, &order.TotalAmount); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, rows.Err()
}

func CreateOrder(dto OrderDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	tx, err := dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `INSERT INTO orders (user_id, order_date, total_amount)
VALUES ($1, $2, $3) RETURNING id`,
		dto.UserID, dto.OrderDate, dto.TotalAmount)

	var orderID int

	err = row.Scan(&orderID)
	if err != nil {
		return 0, err
	}

	for _, product := range dto.OrderProducts {
		_, err = tx.Exec(ctx, `INSERT INTO order_product (order_id, product_id, count)
VALUES ($1, $2, $3)`,
			orderID, product.ProductID, product.Count)
		if err != nil {
			return orderID, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return orderID, err
	}

	return orderID, nil
}

func DeleteOrder(dto OrderDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	tx, err := dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	_, err = dbPool.Exec(ctx, "DELETE FROM order_product WHERE order_id = $1",
		dto.ID)

	if err != nil {
		return 0, err
	}

	_, err = dbPool.Exec(ctx, "DELETE FROM orders WHERE id = $1",
		dto.ID)

	if err != nil {
		return dto.ID, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return dto.ID, err
	}

	return dto.ID, nil
}
