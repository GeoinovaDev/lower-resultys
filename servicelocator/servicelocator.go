package servicelocator

// ServiceLocator ...
type ServiceLocator struct {
	services map[string]interface{}
}

var current *ServiceLocator

func create() *ServiceLocator {
	return &ServiceLocator{services: make(map[string]interface{})}
}

// GetInstance ...
func GetInstance() *ServiceLocator {
	if current == nil {
		current = create()
	}

	return current
}

// Add ...
func (s *ServiceLocator) Add(name string, service interface{}) *ServiceLocator {
	s.services[name] = service

	return s
}

// Get ...
func (s *ServiceLocator) Get(name string) interface{} {
	return s.services[name]
}

// Clear ...
func (s *ServiceLocator) Clear() {
	for service := range s.services {
		delete(s.services, service)
	}

	s.services = make(map[string]interface{})
}
