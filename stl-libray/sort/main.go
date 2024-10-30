package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // This manupluates the original array
	fmt.Println(ages)

	index := sort.SearchInts(ages, 45)
	fmt.Println(index)

	name := []string{"youshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(name)
	fmt.Println(name)
	fmt.Println(sort.SearchStrings(name, "mario"))
}
