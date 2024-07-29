package main

import ("fmt"
		"sync"
		)

func printSomething(s string, wg *sync.WaitGroup){
	fmt.Println(s)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, x := range words {
		wg.Add(1)
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
	
	wg.Add(1)
	go printSomething("This is the second thing to be printed!", &wg)
	wg.Wait()
}