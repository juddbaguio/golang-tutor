package somepkg

import "fmt"

type Person struct {
	Name string
	Age  int
}

// suitable for Read
func (mystruct Person) FormatAgePerson() string {
	return fmt.Sprintf("%s - %v", mystruct.Name, mystruct.Age)
}

// suitable for Read and Write
func (mystruct *Person) ChangeName(newName string) {
	mystruct.Name = newName
}
