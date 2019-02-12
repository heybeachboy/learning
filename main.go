package main

import (
	"fmt"
	"learning/Algorithm"
	)

func main() {
	 data := Algorithm.CreateRandomList(10000,10000)
	  s := Algorithm.NewSort(data)
	  resp := s.BubbleSort()
	  fmt.Println(resp)
	  fmt.Println("execute time:",s.Interval)
	
}
