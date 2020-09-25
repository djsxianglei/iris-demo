package models

type Product struct {
	ID           int64  `json:"id" sql:"ID" cashub:"id"`
	ProductName  string `json:"ProductName" sql:"product_name" cashub:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"product_num" cashub:"productNmb"`
	ProductImage string `json:"ProductImage" sql:"product_image" cashub:"productImage"`
	ProductUrl   string `json:"ProductUrl" sql:"product_url" cashub:"ProductUrl"`
}
