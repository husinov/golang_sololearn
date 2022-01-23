package main

import "fmt"

type Timer struct {
	id    string
	value int
}

func (time *Timer) tick() {
	time.value++
	fmt.Println(time.value)

}
func main() {
	var x int
	fmt.Scanln(&x)

	time := Timer{"timer1", 0}

	for i := 0; i < x; i++ {
		time.tick()
	}
}
