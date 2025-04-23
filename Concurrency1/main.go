package main

import (
	"fmt"
	"time"
)

// func someFunc(num string) {
// 	fmt.Println(num)
// }

// func doWork(done <-chan bool) {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println("Doing Work")
// 		}
// 	}
// }

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			time.Sleep(time.Second * 3)
		}
		close(out)
	}()
	return out
}

func main() {

	// myChannel := make(chan string)
	// anotherChannel := make(chan string)

	// for i := 0; i < 100; i++ {
	// 	go func(i int) {
	// 		myChannel <- fmt.Sprintf("data %d", i)
	// 	}(i)

	// 	go func(i int) {
	// 		anotherChannel <- fmt.Sprintf("data2 %d", i)
	// 	}(i)
	// }

	// l := list.New()

	// for i := 0; i < 100; i++ {
	// 	select {
	// 	case msgFromMyChannel := <-myChannel:
	// 		l.PushBack(msgFromMyChannel)
	// 	case msgFromOtherChannnel := <-anotherChannel:
	// 		l.PushBack(msgFromOtherChannnel)
	// 	}
	// }

	// for e := l.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	// ---------------------------

	// charChannel := make(chan string, 3)
	// chars := []string{"a", "b", "c"}

	// for _, s := range chars {
	// 	select {
	// 	case charChannel <- s:
	// 	}
	// }

	// close(charChannel)

	// for result := range charChannel {
	// 	fmt.Println(result)
	// }

	// --------------------------------

	// done := make(chan bool)

	// go doWork(done)

	// time.Sleep(time.Second * 3)

	// close(done) // - signal to the doWork coroutine to stop

	// --------------------------------

	nums := []int{2, 3, 4, 7, 1}
	dataChannel := sliceToChannel(nums)
	finalChannel := sq(dataChannel)
	for n := range finalChannel {
		fmt.Println(n)
	}

}
