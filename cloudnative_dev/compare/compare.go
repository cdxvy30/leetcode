package main

import "fmt"

type Coin int

const (
	Real  Coin = 0
	Light      = -1
	Heavy      = 1
)

func compare(coins []Coin, start1, end1, start2, end2 int) int {
	sum1, sum2 := 0, 0
	for i := start1; i < end1; i++ {
		sum1 += int(coins[i])
	}
	for i := start2; i < end2; i++ {
		sum2 += int(coins[i])
	}
	if sum1 < sum2 {
		return -1
	} else if sum1 == sum2 {
		return 0
	} else {
		return 1
	}
}

func findFake(coins []Coin) (int, Coin) {
	comp1 := compare(coins, 0, 4, 4, 8)

	switch comp1 {
	case 0: // Fake is in C (or doesn't exist)
		comp2 := compare(coins, 8, 10, 0, 2)
		switch comp2 {
		case 0:
			comp3 := compare(coins, 10, 11, 11, 12)
			return 10 + comp3, Coin(comp3)
		default:
			comp3 := compare(coins, 8, 9, 9, 10)
			return 8 + comp3, Coin(comp3)
		}
	case -1: // Fake is in B
		comp2 := compare(coins, 4, 7, 0, 3)
		switch comp2 {
		case -1:
			comp3 := compare(coins, 4, 5, 5, 6)
			return 4 + comp3, Coin(comp3)
		default:
			comp3 := compare(coins, 7, 8, 0, 1)
			return 7, Coin(comp3)
		}
	case 1: // Fake is in A
		comp2 := compare(coins, 0, 3, 4, 7)
		switch comp2 {
		case 1:
			comp3 := compare(coins, 0, 1, 1, 2)
			return 0 + comp3, Coin(comp3)
		default:
			comp3 := compare(coins, 3, 4, 4, 5)
			return 3, Coin(comp3)
		}
	}
	return -1, Real // Should not reach here
}

func main() {
	coins := []Coin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1} // Last coin is heavy (fake)
	index, result := findFake(coins)
	if result == Real {
		fmt.Println("No fake coin found.")
	} else {
		fmt.Printf("Fake coin is at index %d and it is ", index)
		if result == Light {
			fmt.Println("lighter.")
		} else {
			fmt.Println("heavier.")
		}
	}
}
