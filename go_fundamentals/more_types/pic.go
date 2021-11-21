package moretypes

func Pic(dx, dy int) [][]uint8 {
	modpic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pixels := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pixels[j] = uint8((i + j) / 2)
		}
		modpic[i] = pixels
	}
	return modpic
}
