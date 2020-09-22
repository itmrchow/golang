package main

func main() {
	animals := make([]IAnimal, 0)

	duck := New(Duck)
	dog := New(Dog)
	tiger := New(Tiger)

	animals = append(animals, duck, dog, tiger)

	for _, a := range animals {
		a.Speak()
	}

}
