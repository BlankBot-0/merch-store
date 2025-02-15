package models

type CredentialsDTO struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type InfoForTokenDTO struct {
	Password string `db:"password"`
	Id       int64  `db:"id"`
}
