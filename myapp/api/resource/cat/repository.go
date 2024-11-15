package cat

import (
	"gorm.io/gorm"
)

type CatRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) *CatRepository {
	return &CatRepository{
		db: db,
	}
}

func (r *CatRepository) List() (Cats, error) {
	cats := make([]*Cat, 0)
	if err := r.db.Find(&cats).Error; err != nil {
		return nil, err
	}

	return cats, nil
}
