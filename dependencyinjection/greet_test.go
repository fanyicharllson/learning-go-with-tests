package dependencyinjection

import "testing"

func TestGreet(t *testing.T) {
	t.Run("greet user", func(t *testing.T) {
		// Properly create a Greeter with a MessageProvider
		greeter := Greeter{provider: MessageProvider{}}
		got := greeter.Greet()
		want := "Hello1" // matches MessageProvider.Message()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("greet user with fake provider (dependency injection)", func(t *testing.T) {

		greeter := GreeterWithInterface{provider: FakeProvider{}}
		got := greeter.provider.Message()
		want := "Test!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}