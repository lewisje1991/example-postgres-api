package tasks

type Service struct {
	store *Store
}

func NewService(s *Store) *Service {
	return &Service{
		store: s,
	}
}