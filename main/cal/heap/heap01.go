package heap

type HeapObject struct {
	Heap     []int
	Capacity int
	Size     int
}

func (h *HeapObject) Push(value int) {
	if h.Size >= h.Capacity {
		panic("heap is full")
	}

	h.Heap = append(h.Heap, value)
	h.HeapInsert(h.Heap, h.Size)
	h.Capacity++
	h.Size++
}

//Pop 弹出最大值
func (h *HeapObject) Pop() int {
	if h.Size == 0 {
		panic("heap is empty")
	}

	result := h.Heap[0]
	h.Size--
	swap(h.Heap, 0, h.Size)
	h.Heapify(h.Heap, 0, h.Size)
	return result
}

//HeapInsert 堆插入 对比自己的父节点（index-1）/ 2
func (h *HeapObject) HeapInsert(arr []int, idx int) {
	for arr[idx] > arr[(idx-1)/2] {
		swap(arr, idx, (idx-1)/2)

		idx = (idx - 1) / 2
	}
}

//Heapify 节点下沉 直到不比子节点小
func (h *HeapObject) Heapify(arr []int, idx, heapSize int) {
	left := 2*idx + 1
	for left < heapSize {
		largest := left
		// 判断是否有右孩子节点
		if left+1 < heapSize && arr[left] < arr[left+1] {
			largest = left + 1
		}

		if arr[largest] > arr[idx] {
			swap(arr, largest, idx)
			idx = largest
			left = 2*idx + 1
		} else {
			break
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
