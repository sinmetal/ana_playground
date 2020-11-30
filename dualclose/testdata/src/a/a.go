package a

func f() {
	ch := make(chan int)
	close(ch)
	close(ch) // want "close dual chan"

	ch2 := make(chan int)
	close(ch2)

	myclose := func(ch chan int) {
		close(ch) // want "close dual chan" 追いたいけど追えてない
	}
	ch3 := make(chan int)
	myclose(ch3)
	myclose(ch3)
}
