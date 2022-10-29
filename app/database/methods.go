package database

import "tipo-server/app/models"

func (d *DB) CreateWord(w *models.Word) (*models.Word, error) {
	err := d.db.Create(&w).Error
	if err != nil {
		return nil, err
	}

	return w, err
}

func (d *DB) FindWordByInput(input *string) (*models.Word, error) {
	var w = &models.Word{
		Input: *input,
	}
	err := d.db.Where(w).First(&w).Error
	if err != nil {
		return w, err
	}

	return w, err
}
