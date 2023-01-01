package product

func GetByIds(ids []int) ([]Product, error) {
	var products []Product
	db.Raw("select p.name, p.price, b.name, b.status,  from products as p inner join brands as b on p.id in (?)", ids).Scan(&products)
	return products, nil
}

func GetById(id int) (product Product, err error) {
	err = db.First(&product, id).Error
	return product, err
}

func GetByName(name string) (product Product, err error) {
	err = db.Where("name = ?", name).First(&product).Error
	return product, err
}

func GetByNameAndBrandId(name string, brandId int) (product Product, err error) {
	err = db.Where("name = ? AND brand_id = ?", name, brandId).First(&product).Error
	return product, err
}