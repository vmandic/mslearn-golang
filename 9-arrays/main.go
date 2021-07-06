package main

import "fmt"

func main() {
	var a [3]int
	a[1] = 10

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	// NOTE: this is an array with implicit length [...]
	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}

	fmt.Println(cities)

	numbers := [...]int{99: -1, 95: 15}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	// NOTE: slices are a portion or a window to an underlying array
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	fmt.Println("months array:", months)
	fmt.Println("len:", len(months), "cap:", cap(months))

	// NOTE: this slice has an underlying array of len 12

	// NOTE: this is a new slice based on existing slice, takes element on [x:y] x index and goes for y-x len ie in case below we take 3 elements ie. 6-3 = 3 resulting in ["April", "May", "June"]
	quarter2 := months[3:6] // NOTE: cap is 9 because we skip first three elements from original slice

	// because the quarter2 is slice on existing slice of 12 months, you can extend the slice len from start and make it more "lengthy" 4-0 = 4 elements, if x is not specified it means 0, ie. from first element, ie. index 0
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	for i, v := range months {
		monthMsg := fmt.Sprintf("Month %d: %s", i+1, v)
		fmt.Println(monthMsg)
	}

	// NOTE: append() adds elements at end of an array
	var numbers2 []int

	for i := 0; i < 10; i++ {
		// NOTE how capacity grows by factor of 2, a geometric progression
		numbers2 = append(numbers2, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers2), numbers2)
	}

	newArr := removeStringFromArray(months, "August")
	fmt.Println("Removed August:", newArr)
}

func removeStringFromArray(arr []string, element string) []string {
	for i, v := range arr {
		if v == element {
			newArr := append(arr[:i], arr[i+1:]...) // NOTE: You can pass a slice s directly to a variadic function if you unpack it with the s... notation.
			return newArr
		}
	}

	return arr
}
