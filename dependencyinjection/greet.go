package dependencyinjection

// Use an interface for more flexible dependency injection
type MessageProviderInterface interface {
	Message() string
}

type GreeterWithInterface struct {
	provider MessageProviderInterface
}

type MessageProvider struct{}
type FakeProvider struct{}

type Greeter struct {
	provider MessageProvider
}

// real provider
func (m MessageProvider) Message() string {
	return "Hello1"
}
func (m FakeProvider) Message() string {
	return "Test!"
}

func (g Greeter) Greet() string {
	return g.provider.Message()
}
