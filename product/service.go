package product

func GetByIds(ids []int) []Product {
	var products []Product
	db.Raw("select p.name, p.price, b.name, b.status,  from products as p inner join brands as b on p.id in (?)", ids).Scan(&products)
	return products
}