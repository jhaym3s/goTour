package main

import "fmt"
func main()  {
	ch := make(chan string)
	ss := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	for _, v := range ss {
		go func(str string) {
			ch <- str
			}(v)	
	}
	for i := 0; i < len(ss); i++ {
		fmt.Println(<-ch)
	}
	c := make(chan int, 10)
	for v := range c {  // why do you not need the index when working with channels 
		fmt.Println(v)	
	}
}
