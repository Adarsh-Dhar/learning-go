package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FirstName      *string
	LastName       *string
	PassWord       *string
	Email          *string
	Phone          *string
	Token          *string
	RefreshToken   *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         string
	UserCart       []ProductUser
	AddressDetails []Address
	OrderStatus    []Order
}

type Product struct {
	ProductId   primitive.ObjectID
	ProductName *string
	Price       *uint64
	Rating      *uint8
	Image       *string
}

type ProductUser struct {
	ProductId   primitive.ObjectID
	ProductName *string
	Price       *int
	Rating      *uint
	Image       *string
}

type Address struct {
	AddressId primitive.ObjectID
	House     *string
	Street    *string
	City      *string
	PinCode   *string
}

type Order struct {
	OrderId       primitive.ObjectID
	OrderCart     []ProductUser
	OrderAt       time.Time
	Price         int
	Discount      *int
	PaymentMethod Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
