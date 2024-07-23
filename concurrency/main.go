package main

import (
	// "fmt"
	"fmt"
	"math/rand"
	// "strings"
	"sync"
)

var lock sync.Mutex

func justTest(wg *sync.WaitGroup, num int, ch chan int) {
	defer wg.Done()
	r := rand.Intn(20)
	fmt.Println(r, num)
	ch <- r + num
}

func modifyStr(st *string, i int) {
	// without lock unlock race condition will happen
	// and if you see the count of the string seperated by -
	// it will be incorrect
	lock.Lock()
	defer lock.Unlock()
	r := rand.Intn(20)
	*st = fmt.Sprintf("%s%d%d-", *st, r, i)
}

func genrateEvenOdd(even chan int, odd chan int) {
	// adding even or odd numbers to different channels
	r := rand.Intn(100)
	if r%2 == 0 {
		even <- r
	} else {
		odd <- r
	}
}

func main() {
	// trying waitgroups and data sharing with channels
	// var wg sync.WaitGroup
	// ch:= make(chan int, 4)
	// wg.Add(4)
	// go justTest(&wg, 1, ch)
	// go justTest(&wg, 1, ch)
	// go justTest(&wg, 2, ch)
	// go justTest(&wg, 2, ch)
	// wg.Wait()
	// close(ch)
	// for {
	// 	val, ok := <- ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(val)
	// }

	// trying resource locking
	// var wg2 sync.WaitGroup
	// wg2.Add(100)
	// var str string = ""
	// for i:=0;i<100;i++ {
	// 	i:=i
	// 	fmt.Println("number", i)
	// 	go func () {
	// 		defer wg2.Done()
	// 		modifyStr(&str, i)
	// 	}()
	//
	// }
	// wg2.Wait()
	// fmt.Println(str)
	// // code to chek
	// ll:=len(strings.Split(str, "-"))
	// fmt.Println(ll)

	// trying multi channel with select

	oddChan := make(chan int, 100)
	evenChan := make(chan int, 100)
	var wg3 sync.WaitGroup
	const val int = 100
	wg3.Add(val)
	for range val {
		go func() {
			defer wg3.Done()
			genrateEvenOdd(oddChan, evenChan)
		}()
	}
	for range 100 {
		select {
		case oddNum := <-oddChan:
			{
				fmt.Println("holla got a even number", oddNum)
			}
		case evenNum := <-evenChan:
			{
				fmt.Println("holla got a odd number", evenNum)

			}
		}
	}
	wg3.Wait()
	close(oddChan)
	close(evenChan)
}
