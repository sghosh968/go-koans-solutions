package go_koans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == "apple") // slices seem like arrays
	assert(len(fruits) == 3)  // in nearly all respects

	tasty_fruits := fruits[1:3]           // we can even slice slices
	assert(tasty_fruits[0] == "orange") // slices of slices also share the underlying data

	pregnancy_slots := []string{"baby", "baby", "lemon"}
	assert(cap(pregnancy_slots) == 3) // the capacity is initially the length

	pregnancy_slots = append(pregnancy_slots, "baby!")
	assert(len(pregnancy_slots) == 4) // slices can be extended with append(), much like realloc in C
	assert(cap(pregnancy_slots) == 6) // but with better optimizations

	pregnancy_slots = append(pregnancy_slots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	assert(len(pregnancy_slots) == 7) // append() can take N arguments to append to the slice
	assert(cap(pregnancy_slots) == 12) // the capacity optimizations have a guessable algorithm
}

// NOTE: caveat on how cap fot a slice is calculated. ref: https://www.dotnetperls.com/slice-go
// - When on append the slice has no room for the new element(s) the cap doubles it existing capacity.
// Ex: on line #15 on an append when slice runs out of capacity it doubles its cap size to 6 (i.e 3*2)
// Again on another append (in line #19) slice runs out of capacity on last append ("they must be Catholic") it doubles its cap size i.e. 12 (6*2)
