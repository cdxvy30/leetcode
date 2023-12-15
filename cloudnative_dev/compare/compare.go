package main

import "fmt"

type Coin int

const (
	Real  Coin = 0
	Light      = -1
	Heavy      = 1
)

func compare(coins []Coin, a, b, c, d int) int {
	return int(coins[a] + coins[b] + coins[c] + coins[d])
}

func findFake(coins []Coin) (int, Coin) {
	comp1 := compare(coins, 0, 1, 2, 3) - compare(coins, 4, 5, 6, 7)

	switch {
	case comp1 == 0: // C contains fake coin or no fake coin exists
		comp2 := compare(coins, 8, 9, 10, 11) - compare(coins, 0, 1, 2, 3)
		switch {
		case comp2 == 0:
			return -1, Real
		case comp2 < 0: // 8, 9, 10, 11, and one of them is Light
			if coins[8] != coins[9] {
				return 8, coins[8]
			}
			return 10, coins[10]
		case comp2 > 0: // 8, 9, 10, 11, and one of them is Heavy
			if coins[8] != coins[9] {
				return 8, coins[8]
			}
			return 10, coins[10]
		}
	case comp1 < 0: // B contains the fake coin, and it is Light
		comp2 := compare(coins, 4, 5, 6, 7) - compare(coins, 0, 1, 2, 3)
		if comp2 != 0 {
			if coins[4] != coins[5] {
				return 4, coins[4]
			}
			return 6, coins[6]
		}
	case comp1 > 0: // A contains the fake coin, and it is Heavy
		comp2 := compare(coins, 0, 1, 2, 3) - compare(coins, 4, 5, 6, 7)
		if comp2 != 0 {
			if coins[0] != coins[1] {
				return 0, coins[0]
			}
			return 2, coins[2]
		}
	}

	return -1, Real // should not reach here
}

func main() {
	coins := []Coin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1} // the last coin is fake and Heavy
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
