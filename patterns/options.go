package patterns

// Options pattern is a way to pass optional parameters to a function.
// It is useful when you have a function with many parameters, and you want to make it more readable.
// It is also useful when you want to add new parameters to a function without breaking the existing code.

// --

type Service struct {
	optionA string
	optionB string
}

func NewService(opts ...Option) *Service {
	svc := new(Service)
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

// ---

type Option func(*Service)

func WithOptionA(optionA string) Option {
	return func(s *Service) {
		s.optionA = optionA
	}
}

func WithOptionB(optionB string) Option {
	return func(s *Service) {
		s.optionB = optionB
	}
}