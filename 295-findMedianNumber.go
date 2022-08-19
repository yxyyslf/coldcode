//The median is the middle value in an ordered integer list. If the size of the
//list is even, there is no middle value and the median is the mean of the two
//middle values.
//
//
// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
//
//
// Implement the MedianFinder class:
//
//
// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data
//structure.
// double findMedian() returns the median of all elements so far. Answers
//within 10⁻⁵ of the actual answer will be accepted.
//
//
//
// Example 1:
//
//
//Input
//["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//[[], [1], [2], [], [3], []]
//Output
//[null, null, null, 1.5, null, 2.0]
//
//Explanation
//MedianFinder medianFinder = new MedianFinder();
//medianFinder.addNum(1);    // arr = [1]
//medianFinder.addNum(2);    // arr = [1, 2]
//medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
//medianFinder.addNum(3);    // arr[1, 2, 3]
//medianFinder.findMedian(); // return 2.0
//
//
//
// Constraints:
//
//
// -10⁵ <= num <= 10⁵
// There will be at least one element in the data structure before calling
//findMedian.
// At most 5 * 10⁴ calls will be made to addNum and findMedian.
//
//
//
// Follow up:
//
//
// If all integer numbers from the stream are in the range [0, 100], how would
//you optimize your solution?
// If 99% of all integer numbers from the stream are in the range [0, 100], how
//would you optimize your solution?
//
//
// Related Topics Two Pointers Design Sorting Heap (Priority Queue) Data Stream
//👍 7867 👎 140

package main

import "container/heap"

//leetcode submit region begin(Prohibit modification and deletion)
type SmallIntHeap []int
type BigIntHeap []int

func (s SmallIntHeap) Len() int {
	return len(s)
}

func (b BigIntHeap) Len() int {
	return len(b)
}
func (s SmallIntHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (b BigIntHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (s SmallIntHeap) Less(i, j int) bool {
	return s[i] < s[j]
}

func (b BigIntHeap) Less(i, j int) bool {
	return b[i] > b[j]
}

func (s *SmallIntHeap) Push(val interface{}) {
	*s = append(*s, val.(int))
}

func (b *BigIntHeap) Push(val interface{}) {
	*b = append(*b, val.(int))
}

func (s *SmallIntHeap) Pop() interface{} {
	n := len(*s)
	val := (*s)[n-1]
	*s = (*s)[:n-1]
	return val
}

func (b *BigIntHeap) Pop() interface{} {
	n := len(*b)
	val := (*b)[n-1]
	*b = (*b)[:n-1]
	return val
}

type MedianFinder struct {
	h1 *BigIntHeap
	h2 *SmallIntHeap
}

func MedianFinderConstructor() MedianFinder {
	m := MedianFinder{
		h1: &BigIntHeap{},
		h2: &SmallIntHeap{},
	}
	return m
}

func (this *MedianFinder) AddNum(num int) {
	x := this.h1
	y := this.h2

	heap.Push(x, num)

	if x.Len()-y.Len() > 1 {
		heap.Push(y, heap.Pop(x))
	}

	if x.Len()-y.Len() == 1 && y.Len() > 0 {
		smallVal := (*x)[0]
		bigVal := (*y)[0]
		if smallVal > bigVal {
			a := heap.Pop(x)
			b := heap.Pop(y)
			heap.Push(y, a)
			heap.Push(x, b)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	m := this.h1.Len()
	n := this.h2.Len()
	x := this.h1
	y := this.h2
	if m == 0 {
		return 0
	}
	if (m+n)%2 == 0 {
		return (float64((*x)[0]) + float64((*y)[0])) / 2
	} else {
		return float64((*x)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)
