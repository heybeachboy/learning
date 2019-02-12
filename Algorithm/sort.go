package Algorithm

import (
	"time"
)

type Sort struct {
	Data      []int
	Size      int
	Interval  time.Duration
	startTime time.Time
	endTime   time.Time
}

func NewSort(data []int) (*Sort) {
	return &Sort{Data: data, Size: len(data), Interval: 0}

}

/**
 *sort start time
 */
func (s *Sort) startRun() {
	s.startTime = time.Now()
}

/**
 *sort end time
 */
func (s *Sort) endRun() {
	s.endTime = time.Now()
	s.Interval = s.endTime.Sub(s.startTime)

}

/**
 * Bubble Sort
 * O(n^2)
 */

func (s *Sort) BubbleSort() ([]int) {
	s.startRun()
	list := s.Data
	for i := len(list) - 1; i > 0; i-- {
		index := i
		for j := i - 1; j >= 0; j-- {
			if list[j] <= list[index] {
				continue
			}
			index = j
		}
		list[i], list[index] = list[index], list[i]

	}
	s.endRun()
	return list
}

/**
 *select sort
 *O(n^2)
 */

func (s *Sort) SelectSort() ([]int) {
	s.startRun()
	list := s.Data
	size := len(list)
	for i := 0; i < size-1; i++ {
		index := i
		for j := i + 1; j < size; j++ {
			if list[j] >= list[index] {
				continue
			}
			index = j
		}

		if i == index {
			continue
		}

		list[i], list[index] = list[index], list[i]
	}

	s.endRun()
	return list
}

func (s *Sort) SelectSortForOptimize() {

}

/**
 *insert sort
 *O(n^2)
 */

func (s *Sort) InsertSort() ([]int) {
	s.startRun()
	list := s.Data
	size := len(list)

	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {

			if list[j] <= list[j+1] {
				continue
			}
			list[j], list[j+1] = list[j+1], list[j]
		}

	}
	s.endRun()
	return list
}

func (s *Sort) ShellSort() ([]int) {
	s.startRun()
	list := s.Data
	
	s.endRun()
	return list
}

/**
 时间复杂度O(nlogn)
 归并排序
 */

func (s *Sort)MergeSort(data []int,start, end int) {
	if len(data) <=0 || start >= end {
		return
	}

	mid := start + ((end - start) / 2)
	s.MergeSort(data,start,mid)
	s.MergeSort(data,mid+1, end)
	s.merge(data, start, mid, end)


}

func (s *Sort)merge(data []int, start, mid, end int) {

	total := (end - start) + 1
	temp := make([]int,total)
	i := start
	j := mid+1

	for t:=0;t < total; t++ {
		if i <=mid && j > end {
			temp[t] = data[i]
			i++
			continue
		}

		if i > mid && j <= end {
			temp[t] = data[j]
			j++
			continue

		}

		if data[i] <  data[j]  {
			temp[t] = data[i]
			i++
		} else {
			temp[t] = data[j]
			j++
		}
	}

	for k,v :=range temp {
		data[start + k] = v
	}


}

/**
  时间复杂度(nlogn)
  快速排序法
 */

func (s *Sort)QuickSort(array []int, left int, right int)  {


	if left >= right {
		return
	}
	n := s.partition(array,left,right)
	s.QuickSort(array, left, n-1)
	s.QuickSort(array, n+1, right)

}

func (s *Sort)partition(data []int,l,r int)(int) {

	focus := data[l]
	j := l
	for i := l+1; i<= r; i++ {
		if data[i] < focus {
			data[j+1], data[i] = data[i],data[j+1]
			j++
		}
	}

	data[l],data[j] = data[j], data[l]
	return j

}

func (s *Sort) HeapSort() ([]int) {
	s.startRun()
	list := s.Data
	s.endRun()
	return list
}

func (s *Sort) RadixSort() ([]int) {
	s.startRun()
	list := s.Data
	s.endRun()
	return list
}
