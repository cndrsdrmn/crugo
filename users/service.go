package users

func NewService(repo IRepository) *service {
	return &service{
		repo: repo,
	}
}

func (srvs *service) All() ([]User, error) {
	return srvs.repo.All()
}

func (srvs *service) Store(user *User) error {
	return srvs.repo.Store(user)
}

func (srvs *service) Show(id uint) (*User, error) {
	return srvs.repo.Show(id)
}

func (srvs *service) Update(id uint, user *User) error {
	return srvs.repo.Update(id, user)
}

func (srvs *service) Destroy(id uint) error {
	return srvs.repo.Destroy(id)
}
