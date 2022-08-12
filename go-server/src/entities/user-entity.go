package entities

type UserEntity struct {
	Id        int    `json:"id"`
	Username  string `json:"name"`
	Password  string `json:"description"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
