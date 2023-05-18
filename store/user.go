package store

import (
	"github.com/pkg/errors"
	"github.com/skrevolve/grpc/database"
	"github.com/skrevolve/grpc/models"
)

func GetUserInfoByEmailAndPassword(email string, password string) ([]models.GetUserInfoByEmailAndPassword, error) {
	db := database.Conn.Mysql
	user := []models.GetUserInfoByEmailAndPassword{}
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