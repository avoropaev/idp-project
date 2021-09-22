package newservice

// +kit:endpoint
// Service contract
type Service interface {
	// add service methods here...
}

// Service constructor
func NewService(store Store) Service {
	return service{store: store}
}

// Service realization
type service struct {
	store Store
}
