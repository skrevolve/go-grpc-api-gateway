package services

import (
	"grpc/config/database"
	"grpc/models"

	"github.com/pkg/errors"
)

func GetUserInfo(email string, password string) ([]models.GetUserInfo, error) {
	db := database.Conn.Mysql
	user := []models.GetUserInfo{}
	if err := db.Raw(`
		SELECT
			  user_info_id
			, email
			, social
			, lang
			, block
		FROM user_info
		WHERE
			email = ? AND
			password = ?
	`, email, password).Scan(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}