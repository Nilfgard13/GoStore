package seeder

import (
	"github.com/Nilfgard13/GOSTORE/app/database/faker"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.UserFaker(db)},
		{Seeder: faker.ProductFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeder(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return nil
		}
	}

	return nil
}
