package daos

import (
	db "Cloudbee/internal/database"
	"fmt"
	"log"

	"Cloudbee/internal/database/models"

	"github.com/google/uuid"
)

func SaveCustomer(req *models.Post) error {
	if req != nil {
		if req.Author == "" || req.Content == "" || req.Id == uuid.Nil || req.Title == "" {
			fmt.Println("first case")
		}
		if len(req.Tags) > 0 {
			fmt.Println("tags")
		}

	}
	log.Println("check 6")
	err := db.DB.Debug().Unscoped().Table("posts").Create(req).Error
	if err != nil {
		fmt.Println("err", err)
		log.Println(err)
		return err
	}
	log.Println("check 3")
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
	if id == "" {
		log.Println("error")
	}
	log.Println("check 2")
	err := db.DB.Debug().Unscoped().Table("posts").Where("id =?", &id).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Update(req *models.Post) (*models.Post, error) {
	err := db.DB.Debug().Unscoped().Table("posts").Save(req).Error
	if err != nil {
		return nil, err
	}

	res, err := GetPost(req.Id.String())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Delete(PostId string) error {

	err := db.DB.Debug().Exec(`DELETE FROM posts WHERE id = ? `, PostId).Error
	if err != nil {
		return err
	}
	return nil
}
