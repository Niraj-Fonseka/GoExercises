package animals

type Cat struct {
	name   string
	age    int
	gender string
	color  string
}

func NewCat(name string, gender string, age int, color string) *Cat {
	return &Cat{
		name:   name,
		age:    age,
		gender: gender,
	}
}
func (c Cat) Name() string {
	return c.name
}

func (c Cat) Age() int {
	return c.age
}

func (c Cat) Gender() string {
	return c.gender
}

func (c Cat) Color() string {
	return c.color
}
