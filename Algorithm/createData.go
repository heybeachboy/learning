package Algorithm

import (
	"kubernetes/staging/src/k8s.io/apimachinery/pkg/util/rand"
	"time"
)

func CreateRandomList(size, max int) ([]int) {
	var list []int
	for i := 0; i < size; i++ {
		time.Sleep(time.Microsecond)
		list = append(list, rand.Intn(max))
	}
	return list
}
