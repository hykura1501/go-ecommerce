package repositories

import (
	"BE_Ecommerce/src/entity"

	"gorm.io/gorm"
)

func GetAllUsers(dbInstance *gorm.DB, query *entity.UserQuery) ([]entity.User, error) {
	var users []entity.User

	querySQL := `
		SELECT 
			user_id, 
			username, 
			permission, 
			provider_id, 
			created_at,
			login_provider,
			fullname,
			avatar,
			phone,
			address
		FROM users
		WHERE (@search = '' OR username ILIKE '%' || @search || '%' OR fullname ILIKE '%' || @search || '%')
		LIMIT @limit OFFSET @offset
	`

	params := map[string]interface{}{
		"search": query.Search,
		"limit":  query.Limit,
		"offset": query.Offset,
	}

	err := dbInstance.Raw(querySQL, params).Scan(&users).Error

	if err != nil {
		return nil, err
	}
	var count int64
	countSQL := ` 
		SELECT COUNT(*) FROM users
		WHERE (@search = '' OR username ILIKE '%' || @search || '%' OR fullname ILIKE '%' || @search || '%')
	`
	err = dbInstance.Raw(countSQL, params).Count(&count).Error

	if err != nil {
		return nil, err
	}
	query.SetTotalPages(count)
	return users, nil
}

func GetUserById(dbInstance *gorm.DB, userId int) (*entity.User, error) {
	var user entity.User
	err := dbInstance.Table("users").Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetNewUsersStatistics(dbInstance *gorm.DB) ([]entity.NewUserStatistic, error) {
	var newUserStatistic []entity.NewUserStatistic
	query := `
		SELECT 
			EXTRACT(MONTH FROM created_at)::INTEGER AS month, 
			EXTRACT(YEAR FROM created_at)::INTEGER AS year, 
			COUNT(*)::INTEGER AS quantity
		FROM users
		GROUP BY 
			EXTRACT(MONTH FROM created_at), 
			EXTRACT(YEAR FROM created_at)
		ORDER BY 
			year, month
		LIMIT 12
	`
	err := dbInstance.Raw(query).Scan(&newUserStatistic).Error
	if err != nil {
		return nil, err
	}
	return newUserStatistic, nil
}

func ResetPassword(dbInstance *gorm.DB, userId int, newPassword string) error {
	err := dbInstance.Model(&entity.User{}).Where("user_id = ?", userId).Update("password", newPassword).Error
	if err != nil {
		return err
	}
	return nil
}
