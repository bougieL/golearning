package models

// Users struct
type Users struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAllUsers services
func GetAllUsers() ([]*Users, error) {
	var users []*Users
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById
