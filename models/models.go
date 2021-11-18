package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       *string
	Mobile      string
	VerifyAt    time.Time
	Status      uint
	Password    string
	Roles       *[]Role       `gorm:"many2many:role_user;"`
	Permissions *[]Permission `gorm:"many2many:permission_user;"`
	Zipcodes    *[]Zipcode    `gorm:"many2many:user_zipcode;"`
}

type Role struct {
	gorm.Model
	Name        string
	Permissions *[]Permission `gorm:"many2many:permission_role;"`
}

type Permission struct {
	gorm.Model
	Name  string
	Roles *[]Role `gorm:"many2many:permission_role;"`
}

type Product struct {
	gorm.Model
	Name        string
	Description string `gorm:"type:text"`
	Images      datatypes.JSON
	Attributes  *[]Attribute `gorm:"many2many:attribute_product"`
	Categories  *[]Category  `gorm:"many2many:category_product;"`
	Vendors     *[]Vendor    `gorm:"many2many:product_vendor;"`
}

type Attribute struct {
	gorm.Model
	Name     string
	Products *[]Product `gorm:"many2many:attribute_product"`
}

type AttributeProduct struct {
	AttributeID uint   `gorm:"primaryKey"`
	ProductID   uint   `gorm:"primaryKey"`
	Value       string `gorm:"type:text"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type ProductVendor struct {
	ProductID uint `gorm:"primaryKey"`
	VendorID  uint `gorm:"primaryKey"`
	Quantity  uint
	Price     uint
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Category struct {
	gorm.Model
	Name     string
	Products *[]Product `gorm:"many2many:category_product;"`
}

type Order struct {
	gorm.Model
	Name          string
	Tax           int
	Discount      int
	TotalPrice    uint
	ShipDate      time.Time
	ShipperID     uint
	Shipper       Shipper
	UserID        uint
	User          User
	OrderStatusID uint
	OrderStatus   OrderStatus `gorm:"foreignKey:OrderStatusID;"`
}

type Transaction struct {
	gorm.Model
	Code    string
	Mode    uint
	Status  uint
	Type    uint
	OrderID int
	Order   Order
}

type Shipper struct {
	gorm.Model
	Name string
	Fee  int
}

type OrderStatus struct {
	gorm.Model
	Name        string
	Description string `gorm:"type:text"`
}

type Cart struct {
	gorm.Model
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Quantity  uint
	Price     uint
}

type Vendor struct {
	gorm.Model
	Name     string
	Bio      string     `gorm:"type:text"`
	Products *[]Product `gorm:"many2many:product_vendor;"`
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:text"`
	Status    uint
	ParentID  int `gorm:"TYPE:integer REFERENCES comments"`
	Parent    *Comment
	UserID    int
	User      User
	ProductID int
	Product   Product
}

type Zipcode struct {
	gorm.Model
	Code      string
	Country   string
	City      string
	Address   string
	Longitude float64
	Latitude  float64
	Users     *[]User `gorm:"many2many:user_zipcode;"`
}

func (AttributeProduct) TableName() string {
	return "attribute_product"
}

func (ProductVendor) TableName() string {
	return "product_vendor"
}
