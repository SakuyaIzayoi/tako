package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var basepoints int
	var handstr string
	var dealer_payout int
	var nondealer_payout int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Fu: ")
	fustr, _ := reader.ReadString('\n')
	fustr = strings.Trim(fustr, "\n")
	fu, _ := strconv.Atoi(fustr)

	fmt.Print("Han: ")
	hanstr, _ := reader.ReadString('\n')
	hanstr = strings.Trim(hanstr, "\n")
	han, _ := strconv.Atoi(hanstr)

	fmt.Print("Dealer: ")
	dealer, _ := reader.ReadString('\n')
	dealer = strings.Trim(dealer, "\n")

	fmt.Print("Win Type (tsumo/ron): ")
	wintype, _ := reader.ReadString('\n')
	wintype = strings.Trim(wintype, "\n")

	if han > 4 {

		switch {
		case han == 5:
			basepoints = 2000
			handstr = "Mangan"
			break
		case han > 5 && han < 8:
			basepoints = 3000
			handstr = "Haneman"
			break
		case han > 7 && han < 11:
			basepoints = 4000
			handstr = "Baiman"
			break
		case han > 10 && han < 13:
			basepoints = 6000
			handstr = "Sanbaiman"
			break
		case han > 12:
			fmt.Print("Yakuman Multiplier: ")
			yakuman_multiplier, _ := reader.ReadString('\n')
			yakuman_multiplier = strings.Trim(yakuman_multiplier, "\n")
			yakuman_mult, _ := strconv.Atoi(yakuman_multiplier)
			basepoints = 8000 * yakuman_mult
			handstr = "Yakuman"
			break
		default:
			break
		}

	} else {
		basepoints = fu * iPow(2, 2+han)
		if basepoints > 2000 {
			basepoints = 2000
			handstr = "Mangan"
		} else {
			handstr = "Regular"
		}
	}

	fmt.Println("--------------------")

	fmt.Printf("Fu: %d\n", fu)
	fmt.Printf("Han: %d\n", han)
	fmt.Printf("Dealer: %s\n", dealer)
	fmt.Printf("Base Points: %d\n", basepoints)
	if dealer == "y" {
		if wintype == "tsumo" {
			dealer_payout = roundUp(2 * basepoints)
		} else {
			dealer_payout = roundUp(6 * basepoints)
		}
		fmt.Printf("Payout: %d\n", dealer_payout)
	} else {
		if wintype == "tsumo" {
			nondealer_payout = basepoints
			dealer_payout = roundUp(2 * basepoints)
			fmt.Printf("Payout: %d/%d\n", nondealer_payout, dealer_payout)
		} else {
			dealer_payout = roundUp(4 * basepoints)
			fmt.Printf("Payout: %d\n", dealer_payout)
		}
	}
	fmt.Println(handstr)
}

func iPow(a, b int) int {
	var result int = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}
	return result
}

func roundUp(a int) int {
	return ((a + 99) / 100) * 100
}
