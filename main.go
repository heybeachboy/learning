package main

import (
	"fmt"
	"learning/Algorithm"
	)

func main() {
	 data := Algorithm.CreateRandomList(20,100)
	 fmt.Println("sort before:",data)
	  s := Algorithm.NewSort(data)
	  resp := s.InsertSort()
	  fmt.Println("sort after:",resp)
	  fmt.Println("execute time:",s.Interval)
	
}
