package database

import "tipo-server/app/models"

func (d *DB) SaveUserToken(model *models.UserToken) (*models.UserToken, error) {
	err := d.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, err
}
