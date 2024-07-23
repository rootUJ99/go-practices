package main

import (
	"fmt"
	"sync"
	"time"
)


func waitG(d time.Duration, task string, wg *sync.WaitGroup){
/* func channel(d time.Duration, task string){ */
	time.Sleep(d)
	fmt.Println("checking channel", task)
	wg.Done()
}

func channel(d time.Duration, task string, ch chan <- string){
/* func channel(d time.Duration, task string){ */
	time.Sleep(d)
	fmt.Println("checking channel", task)
	ch <- task
}


func main(){
	startTime:= time.Now()
	// wg:= sync.WaitGroup{}
	// wg.Add(2)
	// go waitG(time.Duration(2 * 1000), "one", &wg)
	// go waitG(time.Duration(2 * 2 *1000), "one", &wg)


	ch := make(chan string)

	go channel(time.Duration(2 * 1000), "one", ch)
	go channel(time.Duration(2 * 2 * 1000), "two", ch)
	
	val1:= <- ch
	val2:= <- ch
	
	fmt.Println(val1, val2)
	endTime:= time.Now()
	fmt.Println(endTime, startTime, "yoo")
	fmt.Println(endTime.Sub(startTime), "yoo")
}
