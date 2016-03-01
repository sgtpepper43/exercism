package hello
import "fmt"
const TestVersion = 1

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
