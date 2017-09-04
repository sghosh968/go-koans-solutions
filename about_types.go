package go_koans

type coolNumber int

func (cn coolNumber) multiplyByTwo() int {
	return int(cn) * 2
}

func aboutTypes() {
	i := coolNumber(4)
	assert(i == coolNumber(4))     // values can be converted between compatible types
	assert(i.multiplyByTwo() == 8) // you can add methods on any type you define
}
// NOTE: Read more about naked return and how the execution works here https://tour.golang.org/basics/7
// TODO check how statement in line #12 works, not sure about it.
