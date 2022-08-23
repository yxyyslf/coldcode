package booking

/*
breakfast beach city center location metro view staff price
not
5
1
2
1
1
2
5
This hotel has a nice view of the city center. The location is perfect.
The breakfast is ok. Regarding location, it is quite far from the city center but price is cheap so it is worth it.
Location is excellent, 5 minutes from the city center. There is also a metro station very close to the hotel.
They said I couldn't take my dog and there were other guests with dogs! That is not fair.
Very friendly staff and a good cost-benefit ratio. Its location is a bit far from the city center.
2
*/
import (
	"container/heap"
	"fmt"
	"strings"
)

type Hotel struct {
	ID    int
	Score int
}
type HotelHeap []Hotel

func (h HotelHeap) Len() int {
	return len(h)
}
func (h HotelHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h HotelHeap) Less(i, j int) bool {
	return h[i].Score < h[j].Score
}
func (h *HotelHeap) Push(val interface{}) {
	*h = append(*h, val.(Hotel))
}
func (h *HotelHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
func AwardTopK(positiveWord, negativeWord string, hotels []int, reviews []string, k int) []int {
	pWords := strings.Split(positiveWord, " ")
	nWords := strings.Split(negativeWord, " ")
	pWordMap := make(map[string]bool, 0)
	nWordsMap := make(map[string]bool, 0)
	for _, word := range pWords {
		pWordMap[word] = true
	}
	for _, word := range nWords {
		nWordsMap[word] = true
	}

	scoreMap := make(map[int]int, 0)
	for i := range hotels {
		scoreMap[hotels[i]] = 0
	}

	for i := range reviews {
		review := reviews[i]
		wordList := strings.Split(review, " ")
		posCount := 0
		negCount := 0
		for _, word := range wordList {
			if word[len(word)-1] == '.' || word[len(word)-1] == ',' {
				word = word[:len(word)-1]
			}
			if _, ok := pWordMap[word]; ok {
				posCount++
			}
			if _, ok := nWordsMap[word]; ok {
				negCount++
			}
		}

		scoreMap[hotels[i]] += posCount*3 - negCount
	}
	fmt.Println(scoreMap)
	h := &HotelHeap{}
	for i, v := range scoreMap {
		hotel := Hotel{
			ID:    i,
			Score: v,
		}
		heap.Push(h, hotel)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, 0)
	for h.Len() > 0 {
		data := heap.Pop(h)
		result = append(result, data.(Hotel).ID)
	}
	return result
}
