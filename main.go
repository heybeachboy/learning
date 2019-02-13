package main

import (
	"fmt"
	"learning/Algorithm"
	)

func main() {
	 data := Algorithm.CreateRandomList(10000,10000)
	 fmt.Println("sort before:",data)
	  s := Algorithm.NewSort(data)
	  resp :=s.MergeSort()
	  fmt.Println("sort after:",resp)
	  fmt.Println("execute time:",s.Interval)
	
}
