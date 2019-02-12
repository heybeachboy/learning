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
 *merge sort
 *
 *
 */

func (s *Sort) MergeSort() ([]int) {
	s.startRun()
	list := s.Data
	s.endRun()
	return list
}

func (s *Sort)merge() {

}

func (s *Sort) QuickSort() ([]int) {
	s.startRun()
	list := s.Data
	s.endRun()
	return list
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
