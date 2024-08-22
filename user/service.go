package user

type Service interface {
	GetAllUsers() ([]User, error)
	RegisterUser(input UserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repository.FindAll()
}

func (s *service) RegisterUser(input UserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Name
	user.Age = input.Age

	return s.repository.Create(user)
}
