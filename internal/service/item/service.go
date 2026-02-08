package itemservice


type Repository interface {}

type Config struct {}

type Service struct {
	config Config
	repo Repository
}

func New(cfg Config ,repository Repository) *Service {
	return &Service{
		config: cfg,
		repo: repository,
	}
}