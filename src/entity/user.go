package entity

type User struct {
	UserId        int     `json:"user_id" gorm:"primaryKey"`
	Username      *string `json:"username" gorm:"unique;not null"`
	Password      *string `json:"-"`
	Fullname      *string `json:"fullname"`
	Permission    int     `json:"permission"`
	LoginProvider *int    `json:"login_provider"`
	ProviderId    *string `json:"provider_id"`
	CreatedAt     *string `json:"created_at"`
	Phone         *string `json:"phone"`
	Address       *string `json:"address"`
	Avatar        *string `json:"avatar"`
	AccountId     *string `json:"account_id"`
}

type LoginLocalRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginGoogleRequest struct {
	Code string `json:"code"`
}

func (User) TableName() string {
	return "users"
}
