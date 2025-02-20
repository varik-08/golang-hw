package orders

import "time"

type OrderProduct struct {
	OrderID   int `db:"order_id" json:"orderId"`
	ProductID int `db:"product_id" json:"productId"`
	Count     int `db:"count" json:"count"`
}

type Order struct {
	ID            int            `db:"id" json:"id"`
	UserID        int            `db:"user_id" json:"userId"`
	OrderDate     time.Time      `db:"order_date" json:"orderDate"`
	TotalAmount   float64        `db:"total_amount" json:"totalAmount"`
	OrderProducts []OrderProduct `json:"orderProducts"`
}

type OrderProductDTO struct {
	OrderID   int
	ProductID int
	Count     int
}

type OrderDTO struct {
	ID            int
	UserID        int
	OrderDate     time.Time
	TotalAmount   float64
	OrderProducts []OrderProductDTO
}
