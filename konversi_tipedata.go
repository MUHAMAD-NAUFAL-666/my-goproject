package main

import "fmt"

func main(){
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("Nilai 32:", nilai32)
	fmt.Println("Nilai 64:", nilai64)
	fmt.Println("Nilai 16:", nilai16)
}