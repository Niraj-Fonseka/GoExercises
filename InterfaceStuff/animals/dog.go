package animals

type Dog struct {
	name   string
	age    int
	gender string
}

func NewDog(name string, gender string, age int) *Dog {
	return &Dog{
		name:   name,
		age:    age,
		gender: gender,
	}
}
func (d Dog) Name() string {
	return d.name
}

func (d Dog) Age() int {
	return d.age
}

func (d Dog) Gender() string {
	return d.gender
}
