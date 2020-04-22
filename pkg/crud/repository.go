package crud

import "github.com/jinzhu/gorm"

type repository struct {
	db      *gorm.DB
	creator DataCreator
}

// NewRepository initializes new instance of repository
func NewRepository(db *gorm.DB, creator DataCreator) *repository {
	return &repository{db: db, creator: creator}
}

/**-- Implement Repository --**/

func (r repository) Create(data interface{}) error {
	return nil
}
func (r repository) Get(ID int64) (interface{}, error) {
	return nil, nil
}
func (r repository) Delete(ID int64) error {
	return nil
}
func (r repository) Update(data interface{}) error {
	return nil
}
func (r repository) All() (interface{}, error) {
	return nil, nil
}
