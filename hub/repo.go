package hub

func GetById(id int) (Hub, error) {
	var hub Hub
	result := db.First(&hub, id)
	return hub, result.Error
}

func GetMany() ([]Hub, error) {
	var hubs []Hub
	result := db.Find(&hubs)
	return hubs, result.Error
}