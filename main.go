package main

import (
	"fmt"

	"github.com/zmirk/graf"
	"github.com/zmirk/zmlib"
)

func main() {
	CharFull := ">"
	CharEmpty := " "
	LenBar := 50
	Len := 2840

	fmt.Println("Test graf 1")
	graf.Test()

	fmt.Println("Test graf 2")
	graf.Test2()

	fmt.Println("Test graf 3")
	graf.Test4()

	println("Обработка данных. Это может занять некоторое время...")

	for i := 0; i <= Len; i++ {
		zmlib.RandSleep(11000, 12000)

		fmt.Printf("\033[1A\033[K")
		println(zmlib.Percentage("Обработано:", CharEmpty, CharFull, LenBar, Len, i))
	}

	fmt.Println("")

	stud := graf.Student{}

	stud.SetName("Alex")
	stud.SetMark(92)
	stud.SetAge(18)

	stud2 := graf.Student{}

	stud2.SetName("Max")
	stud2.SetMark(8)
	stud2.SetAge(34)

	students := [2]graf.Student{stud, stud2}

	for _, v := range students {
		fmt.Println(v.String())
	}

	fmt.Println("Test graf 5")
	for _, v := range graf.Test5() {
		fmt.Println(v.String())
	}

	fmt.Println("Test graf 6")
	graf.Test6()

	fmt.Println("Test graf 7")
	graf.Test7()

	fmt.Println("Test graf 8")
	for i := 0; i < 10; i++ {
		fmt.Printf("Test %d\n", i+1)
		graf.Test8()
	}
}
