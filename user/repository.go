package user

type Repository interface {
	FindAll() ([]User, error)
	Create(user User) (User, error)
}

type repository struct {
	nextID int
	users  []User
}

func NewRepository() *repository {
	return &repository{
		nextID: 1,
		users:  []User{},
	}
}

func (r *repository) FindAll() ([]User, error) {
	return r.users, nil
}

func (r *repository) Create(user User) (User, error) {
	user.ID = r.nextID
	r.users = append(r.users, user)
	r.nextID++

	return user, nil
}
