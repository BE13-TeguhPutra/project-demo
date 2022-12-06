package auth

type ServiceInterface interface {
	Login(email string, pass string) (string, string, error)
}

type RepositoryInterface interface {
	Login(email string, pass string) (string, string, error)
}
