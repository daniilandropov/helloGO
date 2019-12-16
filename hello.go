package main

import (
	"fmt"
	"os"
)

func main() {
	var(
		hello string = "здорова"
		name string = GiveName()
	)

	consoles := []string{"ps4", "x-box one", "nswitch"}
	consoles = append(consoles, "psp")

	for _, value := range consoles{
		defer Enderman(value)
	}

	Hello(hello,name, ReadKey, func(x, y string) string{ return x+" "+y})
}

func Hello(hello string, name string, zaderzka func(), textSetter func(string, string) string){
	fmt.Println(textSetter(hello, name))
	zaderzka()
}

func ReadKey(){
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
}

func GiveName() string{
	return "Иван"
}

func Enderman(consolename string){
	fmt.Println("console "+ consolename)
}