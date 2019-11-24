package animals

type Alien struct {
	name string
	age  int
}

func NewAlien(name string, age int) *Alien {
	return &Alien{
		name: name,
		age:  age,
	}
}
func (a Alien) Name() string {
	return a.name
}

func (a Alien) Age() int {
	return a.age
}
