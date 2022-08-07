package di

// GssktService define GssktService struct
type GssktService struct {
	ClientService ClientService
}

type ClientService struct {
}

// NewGssktService generate GssktService instance
func NewGssktService() *GssktService {
	return &GssktService{}
}
