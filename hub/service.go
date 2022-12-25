package hub

func HubById(id int) (Hub, error) {
	var hub Hub
	result := db.First(&hub, id)
	return hub, result.Error
}
