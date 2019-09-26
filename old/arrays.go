package main
import("fmt")
func main() {
	tanks := []string{"Paladin", "Dark Knight", "GunBreaker", "Warrior"}
	fmt.Println(tanks)
	fmt.Println(len(tanks))
	fmt.Println(tanks[2])
}

// $ go run arrays.go
// [Paladin Dark Knight GunBreaker Warrior]
// 4
// GunBreaker