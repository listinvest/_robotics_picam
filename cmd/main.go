package main

import (
	"fmt"
	"picam"
)

func main() {
	cam := picam.NewCam("/tmp/pictures")

	p1, e := cam.TakePicture("shot-today.jpg")
	if e != nil {
		panic(e)
	}
	fmt.Println(p1)
	p2, e := cam.TakePicture("")
	if e != nil {
		panic(e)
	}
	fmt.Println(p2)
}
