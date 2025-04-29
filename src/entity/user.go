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

type UserQuery struct {
	Paging
	Search string `json:"search" query:"search"`
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

type LoginFacebookRequest struct {
	AccessToken string `json:"access_token"`
}

type FacebookUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture struct {
		Data struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"data"`
	} `json:"picture"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password"`
}

func (User) TableName() string {
	return "users"
}
