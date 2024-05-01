package model

type User struct {
	ID           int64  `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"-" db:"password"`
	DisplayName  string `json:"display_name" db:"display_name"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
	// Other user fields
}
