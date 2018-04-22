package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	fmt.Println("let's qrcode")

	err := qrcode.WriteFile("https://blog.naver.com/lolgo7", qrcode.Medium, 256, "qr.png")

	if err != nil {
		fmt.Println(err)
	}

}
