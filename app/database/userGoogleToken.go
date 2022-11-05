package database

import "tipo-server/app/models"

func (d *DB) SaveUserGoogleToken(model *models.UserGoogleToken) (*models.UserGoogleToken, error) {
	err := d.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, err
}

func (d *DB) FindGoogleTokenByUserId(userId *uint) (*models.UserGoogleToken, error) {
	var data = &models.UserGoogleToken{
		UserId: *userId,
	}
	err := d.db.Where(data).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (d *DB) UpdateGoogleToken(data *models.UserGoogleToken) error {
	err := d.db.Save(&data).Error
	return err
}
