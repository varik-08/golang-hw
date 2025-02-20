package users

type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password"`
}

type UserOrderStatistics struct {
	UserID          int     `json:"userId"`
	OrdersSum       float64 `json:"ordersSum"`
	AvgProductPrice float64 `json:"avgProductPrice"`
}

type UserDTO struct {
	ID       int
	Name     string
	Email    string
	Password string
}
