package main

import (
	"fmt"
	"learning/Algorithm"
	)

func main() {
	 data := Algorithm.CreateRandomList(10,100)
	  s := Algorithm.NewSort(data)
	  resp := s.SelectSort()
	  fmt.Println(resp)
	  fmt.Println("execute time:",s.Interval)
	
}
