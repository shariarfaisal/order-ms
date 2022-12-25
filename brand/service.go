package brand


func getByIds(ids []int) []Brand {
	var brands []Brand
	db.Where("id IN ?", ids).Find(&brands)
	return brands
}