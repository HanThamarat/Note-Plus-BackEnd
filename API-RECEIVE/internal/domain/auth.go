package domain

type AuthEntity struct {
	Email		string	`json:"email"`
	Name		string	`json:"name"`
	AuthToken	string	`json:"authToken"`
}

type AuthDTO struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type AuthRepository interface {
	CreateCredentialAuth(authDTO AuthDTO) (*AuthEntity, error)
}

type AuthUsecase interface {
	CredentialAuth(authDTO AuthDTO) (*AuthEntity, error)
}