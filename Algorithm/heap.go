package Algorithm

import "fmt"

type Heap struct {
	Count    int //交换次数
	Total    int //总共的数据量
	Capacity int //总共容量
	Array    []int
}

func NewHeap(total int) (*Heap) {

	return &Heap{0, 0, total, make([]int, total+1)}
}

func (h *Heap) shiftUp(n int) {
	key := h.Total
	for {
		if h.Array[key] > h.Array[key/2] && key > 1 {
			h.swap(key, key/2)
			key /= 2
		} else {
			break
		}
	}

}
func (h *Heap) PushElement(n int) (int) {
	if h.Total > h.Capacity {
		fmt.Println("已经超过最大容量")
		return -1
	}
	h.Total++
	h.Array[h.Total] = n
	h.shiftUp(n)
	return 0
}

func (h *Heap) PopElement() (int) {

	if h.Total == 0 {
		fmt.Println("已经没有元素了")
		return -1
	}

	max := h.Array[1]
	h.swap(1, h.Total)
	h.Total--
	h.shiftDown(1)
	return max
}

func (h *Heap) shiftDown(k int) {

	for {
		if (k * 2) > h.Total {
			h.Count++
			break
		}

		j := k * 2

		if (j+1) < h.Total && (h.Array[j+1] > h.Array[j]) {
			h.Count += 2
			j++

		}

		if h.Array[k] == h.Array[j] {
			break
		}
		h.swap(k, j)
		k = j

	}

}


func (h *Heap) swap(k1 int, k2 int) {
	h.Array[k1], h.Array[k2] = h.Array[k2], h.Array[k1]
}

func (h *Heap) PrintDataPool() {
	fmt.Println(h.Array)
}

type MaxHeap struct {
	Pools      []int  //接收堆数据
	Total      int    //初始堆总数
	Count      int    //堆总数
	CompareNum uint64 //比较次数
}

//获取大根堆的，推顶元素
func (m *MaxHeap) Peek() (int) {
	if m.Count < 0 {
		log.Info("堆数据为空")
		return -1
	}
	Max := m.Pools[0]
	m.Pools[0], m.Pools[m.Count-1] = m.Pools[m.Count-1], m.Pools[0]
	m.Count--
	m.shiftDown(0)
	return Max
}

//初始化堆操作
func NewMaxHeap(data []int, capacity int) (*MaxHeap) {
	ret := &MaxHeap{Pools: data, Total: capacity, Count: capacity, CompareNum: 0}
	ret.initPoolsToHeap()
	return ret

}

//将传入数组通过父亲节点下沉，满足堆需求
func (m *MaxHeap) initPoolsToHeap() {
	for i := (m.Count - 1) / 2; i >= 0; i-- {
		m.shiftDown(i)
	}

}

//添加或删除元素堆的上浮操作
func (m *MaxHeap) shiftUp(key int) {

	for key >= 0 && m.Pools[(key-1)/2] < m.Pools[key] {
		m.Pools[(key-1)/2], m.Pools[key] = m.Pools[key], m.Pools[(key-1)/2]
		key = (key - 1) / 2

	}

}

//添加或删除元素堆的下沉操作
func (m *MaxHeap) shiftDown(key int) {

	left := (2 * key) + 1

	for left <= m.Count-1 {

		if ((left + 1) <= (m.Count - 1)) && (m.Pools[left] < m.Pools[left+1]) {
			m.CompareNum++
			left += 1
		}

		if m.Pools[key] >= m.Pools[left] {
			m.CompareNum++
			break
		}
		m.Pools[key], m.Pools[left] = m.Pools[left], m.Pools[key]
		key = left
		left = (2 * key) + 1
	}

}

//获取堆中排列前num元素
func (m *MaxHeap) GetTopList(num int) ([]int) {
	if num >= m.Count {
		num = m.Count
	}

	top := make([]int, num)

	for i := 0; i < num; i++ {
		top[i] = m.Peek()
	}
	return top
}