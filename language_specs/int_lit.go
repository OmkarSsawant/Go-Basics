package main

import "fmt"

func main() {
	binary_num := 0b1
	hex_num := 0x244567
	oct_num := 0o34456
	img_num := 34456i

	//with underscores
	uhex_num := 0x2_44_567
	uoct_num := 0o34_456
	uimg_num := 34_456i

	fmt.Println(binary_num, hex_num, oct_num, uhex_num, uoct_num, img_num, uimg_num)
}
