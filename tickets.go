package main
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Spaceline         Days  Trip type  Price")
	fmt.Println("========================================")
	for i := 0; i < 10; i++ {
		num := rand.Intn(3)
		spaceline := "SpaceX"
		switch num {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "Space Adventures"
		default:
			break
		}
		speed := rand.Intn(15) + 16
		way := rand.Intn(2)
		price := 20
		time := (62100000 / speed) / (24 * 60 * 60)
		waytype := "Round-trip"
		switch way {
		case 0:
			waytype = "One-way"
			price += speed
		default:
			price = 2 * (price + speed)
			time *= 2
			break
		}
		fmt.Printf("%-16v %4v   %-10v   $ %v\n", spaceline, time, waytype, price)
	}


}
