package Algorithm

import "time"

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

func (s *Sort) SelectSort() ([]int) {

	return nil
}

func (s *Sort) InsertSort() ([]int) {

	return nil
}

func (s *Sort) MergeSort() ([]int) {

	return nil
}

func (s *Sort) ShellSort() ([]int) {

	return nil
}

func (s *Sort) QuickSort() ([]int) {

	return nil
}

func (s *Sort) HeapSort() ([]int) {

	return nil
}

func (s *Sort) RadixSort() ([]int) {

	return nil
}
