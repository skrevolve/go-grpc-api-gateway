package stores

import (
	"github.com/pkg/errors"
	"github.com/skrevolve/grpc-gateway/models"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore{
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetByEmailAndPassword(email string, password string) (*models.User, error) {
 	var user models.User
	if err := s.db.Raw(`
		SELECT
				user_info_id
			, email
			, social
			, lang
			, block
		FROM user_info
		WHERE
			email=? AND
			password=?
	`, email, password).Scan(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (s *UserStore) GetProfileById(userInfoId uint) (*models.User, error) {
	var user models.User
	if err := s.db.Raw(`
		SELECT
				img_path
			, name
			, gender
			, email
			, country
			, lang
		FROM user_info
		WHERE user_info_id = ?
	`, userInfoId).Scan(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}