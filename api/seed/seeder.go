package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/localhearts/framework/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Panji Dwi Setiawan",
		Email:    "panjidwisetiawan@localhearts.id",
		Password: "password",
	},
	models.User{
		Nickname: "Nadia Slavina",
		Email:    "nadia.slavina@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Madog Hernades",
		Email:    "mad.dog99@gmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
