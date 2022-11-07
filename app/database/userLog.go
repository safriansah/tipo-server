package database

import "tipo-server/app/models"

func (d *DB) SaveUserLog(model *models.UserLog) (*models.UserLog, error) {
	err := d.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, err
}
