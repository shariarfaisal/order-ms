package brand

func GetPartnerById(id int) (partner Partner, err error) {
	err = db.First(&partner, id).Error
	return partner, err
}

func GetPartnerMany() (partners []Partner, err error) {
	err = db.Find(&partners).Error
	return partners, err
}

func GetBrandMany() (brands []Brand, err error) {
	err = db.Find(&brands).Error
	return brands, err
}

func GetBrandById(id int) (brand Brand, err error) {
	err = db.First(&brand, id).Error
	return brand, err
}

func GetCategoryById(id int) (category BrandCategory, err error) {
	err = db.First(&category, id).Error
	return category, err
}

func GetBrandByNameAndPartnerId(name string, partnerId int) (brand Brand, err error) {
	err = db.Where("name = ? AND partner_id = ?", name, partnerId).First(&brand).Error
	return brand, err
}

func GetCategoryByCategoryId(id int) (category BrandCategory, err error) {
	err = db.Where("category_id = ?", id).First(&category).Error
	return category, err
}

func GetCategoryMany() (categories []BrandCategory, err error) {
	err = db.Find(&categories).Error
	return categories, err
}