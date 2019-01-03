package models

// Users struct stdin
type Users struct {
	Model
	Username string `gorm:"unique" json:"username" form:"username" binding:"required,min=4,max=255"`
	Password string `gorm:"" json:"-" form:"password" binding:"required,min=4,max=255"`
}

// GetAllUsers Model
func GetAllUsers() ([]*Users, error) {
	var users []*Users
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID Model
func GetUserByID(id int) (*Users, error) {
	var user *Users
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ExistUserByID Model
func ExistUserByID(id int) (bool, error) {
	var user Users
	if err := db.Select("id").Where("id = ?", id).First(&user).Error; err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// ExistUserByName Model
func ExistUserByName(name string) (bool, error) {
	var user Users
	db.Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// ExistUserByNamePut Model
func ExistUserByNamePut(id int, name string) (bool, error) {
	var user Users
	db.Where("username = ? AND id <> ?", name, id).First(&user)
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// PostUser Model
func PostUser(user Users) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// PutUserByID Model
func PutUserByID(id int, user Users) error {
	if err := db.Model(&user).Where("id = ?", id).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUserByID Model
func DeleteUserByID(id int) error {
	var user Users
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
