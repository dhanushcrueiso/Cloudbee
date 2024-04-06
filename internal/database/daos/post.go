package daos

import (
	db "Cloudbee/internal/database"
	"Cloudbee/internal/database/models"
)

func SaveCustomer(req *models.Post) error {

	err := db.DB.Debug().Unscoped().Table("posts").Create(&req).Error
	if err != nil {
		return err
	}
	return err
}

func GetAllPosts() ([]*models.Post, error) {
	res := []*models.Post{}
	err := db.DB.Debug().Unscoped().Table("posts").Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetPost(id string) (*models.Post, error) {
	res := models.Post{}
	err := db.DB.Debug().Unscoped().Table("posts").Where("id =?", id).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
