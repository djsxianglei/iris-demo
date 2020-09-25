package repositories

import (
	"database/sql"
	"github.com/djsxianglei/iris-demo/common"
	"github.com/djsxianglei/iris-demo/models"
	"strconv"
)

//第一步，先开发对应的接口
//第二步，实现定义的接口
type IProduct interface {
	Conn() error
	Insert(product *models.Product) (int64, error)
	Delete(int64) bool
	Update(product *models.Product) error
	SelectByKey(int64) (*models.Product, error)
	SelectAll() ([]*models.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

//数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return nil
}

//插入
func (p *ProductManager) Insert(product *models.Product) (productId int64, err error) {
	//判断连接是否存在
	if err = p.Conn(); err != nil {
		return
	}
	//准备sql
	sql := "INSERT " + p.table + " SET product_name=?,product_num=?,product_image=?,product_url=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	//3.传入参数
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}
func (p *ProductManager) Delete(productId int64) bool {
	//判断连接是否存在
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "DELETE FROM " + p.table + " WHERE id=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(productId)
	if err != nil {
		return false
	}
	return true
}

func (p *ProductManager) Update(product *models.Product) error {
	//判断连接是否存在
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "UPDATE " + p.table + " set product_name=?,product_num=?,product_image=?,product_url=? WHERE id=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductImage)
	if err != nil {
		return err
	}
	return nil
}

//根据商品ID查询商品
func (p *ProductManager) SelectByKey(productId int64) (product *models.Product, err error) {
	//1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return &models.Product{}, err
	}
	sql := "SELECT * FROM " + p.table + " where id = " + strconv.FormatInt(productId, 10)
	row, err := p.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return &models.Product{}, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &models.Product{}, err
	}
	product = &models.Product{}
	common.DataToStructByTagSql(result, product)
	return
}

//获取所有商品
func (p *ProductManager) SelectAll() (productArray []*models.Product, err error) {
	//1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT * FROM" + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	for _, v := range result {
		product := &models.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return
}
