package users

import (
	"github.com/varik-08/golang-hw/hw15_go_sql/config"
)

func GetUsers() ([]*User, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	rows, err := dbPool.Query(ctx, "SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User

	for rows.Next() {
		var i User
		if err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		users = append(users, &i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUsersOrdersStatistics() ([]*UserOrderStatistics, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	rows, err := dbPool.Query(ctx, `
		select u.id as user_id, sum(coalesce(o.total_amount, 0)) as orders_sum,
		       avg(coalesce(p.price, 0)) as avg_product_price
		from users u
		left join orders o on o.user_id = u.id
		left join order_product op on o.id = op.order_id
		left join products p on p.id = op.product_id
		group by u.id
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var usersOrderStatistics []*UserOrderStatistics

	for rows.Next() {
		var i UserOrderStatistics
		if err = rows.Scan(
			&i.UserID,
			&i.OrdersSum,
			&i.AvgProductPrice,
		); err != nil {
			return nil, err
		}
		usersOrderStatistics = append(usersOrderStatistics, &i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usersOrderStatistics, nil
}

func CreateUser(user UserDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	row := dbPool.QueryRow(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateUser(user UserDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	row := dbPool.QueryRow(ctx, "UPDATE users SET name = $1 WHERE id = $2 RETURNING id",
		user.Name, user.ID)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteUser(user UserDTO) (int, error) {
	ctx := config.GetCTX()
	dbPool := config.GetDB()

	row := dbPool.QueryRow(ctx, "DELETE FROM users WHERE id = $1 RETURNING id",
		user.ID)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
