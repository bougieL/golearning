package models

// Users struct
type Users struct {
	Model
	Username string `json:"username" form:"username" binding:"required,min=4,max=255"`
	Password string `json:"password" form:"password" binding:"required,min=4,max=255"`
}

// GetAllUsers services
func GetAllUsers() ([]*Users, error) {
	var users []*Users
	if err := db.Select("id, username, password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById
func GetUserById(id int) (*Users, error) {
	var user Users
	if err := db.Select("id, username, password").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ExistUserById(id int) (bool, error) {
	var user Users
	if err := db.Select("id").Where("id = ?", id).First(&user).Error; err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func PostUser(user Users) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func PutUserById(id int, user Users) error {
	if err := db.Model(&user).Where("id = ?", id).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int) error {
	var user Users
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
