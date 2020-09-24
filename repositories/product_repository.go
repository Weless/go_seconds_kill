package repositories

import (
	"database/sql"
	"imooc-product/datamodels"
)

type IProduct interface {
	// 连接数据库
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(productId int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(productId int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManger struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManger(table string, db *sql.DB) IProduct {
	return &ProductManger{
		table:     table,
		mysqlConn: db,
	}
}

func (p *ProductManger) Conn() error {

}

func (p *ProductManger) Insert(product *datamodels.Product) (int64, error) {

}

func (p *ProductManger) Delete(productId int64) bool {

}

func (p *ProductManger) Update(product *datamodels.Product) error {

}

func (p *ProductManger) SelectByKey(productId int64) (*datamodels.Product, error) {

}

func (p *ProductManger) SelectAll() ([]*datamodels.Product, error) {

}
