package database

import "tipo-server/app/models"

func (d *DB) SaveUserLog(model *models.UserLog) (*models.UserLog, error) {
	err := d.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, err
}

func (d *DB) FindUserLogByUserId(userId uint) (*[]models.UserLog, error) {
	var data = &[]models.UserLog{}
	err := d.db.Where(models.UserLog{
		UserId: userId,
	}).Limit(10).Preload("User").Preload("Word").Find(&data).Error
	return data, err
}
