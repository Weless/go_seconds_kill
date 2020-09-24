package datamodels

type Product struct {
	ID int64 `json:"id" imooc:"id" sql:"ID"`
	ProductName string `json:"product_name" imooc:"product_name" sql:"product_name"`
	ProductNum int64 `json:"product_num" imooc:"product_num" sql:"product_num"`
	ProductImage string `json:"product_image" imooc:"product_image" sql:"product_image""`
	ProductUrl string `json:"product_url" imooc:"product_image" sql:"product_image"`
}





