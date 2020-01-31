package animal

import "fmt"

// An Animal can introduce itself.
type Animal interface {
	Greet() string
}

type cat struct {
	Name  string
	Color string
	Age   uint
}

type dog struct {
	Name  string
	Breed string
	Age   uint
}

type bird struct {
	Name    string
	Species string
	Age     uint
}

// I've made Greet accept pointers because if it used direct values,
// the structs would be copied in memory. Though these structs are small,
// an Animal implemented by a package consumer could be much larger.

// Make a cat say Meow.
func (c *cat) Greet() string {
	return fmt.Sprintf("%s says Meow!", c.Name)
}

// Make a cat say Bark.
func (d *dog) Greet() string {
	return fmt.Sprintf("%s the %s says Bark!", d.Name, d.Breed)
}

// Make a cat say Tweet.
func (b *bird) Greet() string {
	return fmt.Sprintf("%s the %s says Tweet!", b.Name, b.Species)
}

// Speak makes an Animal introduce itself.
func Speak(a Animal) string {
	return a.Greet()
}
