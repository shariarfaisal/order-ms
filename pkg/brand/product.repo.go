package brand

import "gorm.io/gorm"

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) GetByIds(ids []int) ([]Product, error) {
	var products []Product
	r.DB.Raw("select p.name, p.price, b.name, b.status,  from products as p inner join brands as b on p.id in (?)", ids).Scan(&products)
	return products, nil
}

func (r *ProductRepo) GetById(id int) (product Product, err error) {
	err = r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepo) GetByName(name string) (product Product, err error) {
	err = r.DB.Where("name = ?", name).First(&product).Error
	return product, err
}

func (r *ProductRepo) GetByNameAndBrandId(name string, brandId uint) (product Product, err error) {
	err = r.DB.Where("name = ? AND brand_id = ? AND type = ? ", name, brandId, "single").First(&product).Error
	return product, err
}
