package main

import "fmt"

func main() {
	hex_num := 0x2445p67
	e_hex_num := 0x2445e67
	img_num := 344.56i

	//with underscores
	uhex_num := 0x2_44_5p67
	uimg_num := 34_4.56i

	fmt.Println(hex_num, uhex_num, e_hex_num, uimg_num, img_num)
}
