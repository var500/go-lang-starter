package main

import "fmt"

func main() {
	// an array has fixed length and type
	var ages [3]int = [3]int{20, 25, 30}

	var ageArray = [3]int{20, 25, 30}

	//ShortHand
	names := [4]string{"Joe", "Mario", "Luigi", "Peach"}
	names[1] = "Bowser"

	fmt.Println(ages, len(ages))
	fmt.Println(ageArray, len(ageArray))
	fmt.Println(names, len(names))

	//Slices (use array under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 76)
	fmt.Println(scores)

	fmt.Printf("Scores is of Type %T \n", scores)
	fmt.Printf("Names is of Type %T \n", names)

	//Slice Ranges
	rangeOne := names[1:3]  //Including 1 excluding position 3
	rangeTwo := names[2:]   //Including 1 excluding position 3
	rangeThree := names[:3] //Including 1 excluding position 3
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Dorel")
	fmt.Println(rangeOne)
}
