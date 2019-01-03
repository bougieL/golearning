package models

type Files struct {
	Model
	UserId   int    `gorm:"" form:"userId" json:"userId"`
	Filename string `gorm:"" form:"filename" json:"filename"`
	Realname string `gorm:"uniqe" form:"realname" json:"realname"`
}

func PostFile(file Files) error {
	return db.Create(&file).Error
}

func GetFileByID(id int) (*Files, error) {
	var file *Files
	if err := db.Where("id = ?", id).Find(&file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileByUserID(id int) ([]*Files, error) {
	var files []*Files
	if err := db.Where("user_id = ?", id).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
