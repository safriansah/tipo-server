package database

import "tipo-server/app/models"

func (d *DB) SaveUser(model *models.User) (*models.User, error) {
	err := d.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, err
}

func (d *DB) FindUserByEmail(email *string) (*models.User, error) {
	var data = &models.User{
		Email: *email,
	}
	err := d.db.Where(data).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}
