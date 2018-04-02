package main

import "fmt"

type Animal interface {
	moves() string
	eates() string
	sound() string
}

type Lion struct {
	movement string
	food     string
	says     string
}

type Bird struct {
	movement string
	food     string
	beeps    string
}

type Dog struct {
	movement string
	food     string
	barks    string
}

// Lion
func (l Lion) moves() string {
	return "I stalk"
}
func (l Lion) eates() string {
	return "I'm eating meat"
}
func (l Lion) sound() string {
	return "I roar"
}

// Bird
func (b Bird) moves() string {
	return "I fly"
}
func (b Bird) eates() string {
	return "I'm eating bird seed"
}
func (b Bird) sound() string {
	return "I beep"
}

// Dog
func (d Dog) moves() string {
	return "I run"
}
func (d Dog) eates() string {
	return "I'm eating Eukanuba"
}
func (d Dog) sound() string {
	return "I bark"
}

func main() {
	animal1 := Lion{movement: "stalking", food: "meat", says: "roar"}
	animal2 := Bird{movement: "flying", food: "seeds", beeps: "chilp"}
	animal3 := Dog{movement: "running", food: "Eukanuba", barks: "barks"}
	animals := []Animal{animal1, animal2, animal3}

	for n, _ := range animals {
		fmt.Println("I'm an animal and: ", animals[n].eates())
		fmt.Println("I'm an animal and: ", animals[n].moves())
		fmt.Println("I'm an animal and: ", animals[n].sound())
	}

	fmt.Println(animals)
	l := new(Lion)
	l.movement = "crawl"
	b := new(Bird)
	b.beeps = "chirps"

	fmt.Println("This animal does: ", l.movement, "and: ", b.beeps)

}

/*
In the above code within the for loop, we are able to ask for the eating habbits of an animal,
but get different responses each time based on what is the intrinsic type,
whether Lion or Bird. At the point of calling, we only know that it is an Animal -
so it appears that the animal is morphing each time. Polymorphism!
*/
