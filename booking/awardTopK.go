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
	if h[i].Score == h[j].Score {
		return h[i].ID > h[i].ID
	}
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

func AwardTopKHotels(positiveKeywords, negativeKeywords string, hotelIds []int32, reviews []string, k int32) []int32 {
	pWords := strings.Split(positiveKeywords, " ")
	nWords := strings.Split(negativeKeywords, " ")
	pWordMap := make(map[string]bool, 0)
	nWordsMap := make(map[string]bool, 0)
	for _, word := range pWords {
		pWordMap[word] = true
	}
	for _, word := range nWords {
		nWordsMap[word] = true
	}

	scoreMap := make(map[int]int, 0)
	for i := range hotelIds {
		scoreMap[int(hotelIds[i])] = 0
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
			word = strings.ToLower(word)
			if _, ok := pWordMap[word]; ok {
				posCount++
			}
			if _, ok := nWordsMap[word]; ok {
				negCount++
			}
		}
		scoreMap[int(hotelIds[i])] += posCount*3 - negCount
	}
	h := &HotelHeap{}
	for i, v := range scoreMap {
		hotel := Hotel{
			ID:    i,
			Score: v,
		}
		heap.Push(h, hotel)
		if h.Len() > int(k) {
			heap.Pop(h)
		}
	}
	result := make([]int32, 0)
	for h.Len() > 0 {
		data := heap.Pop(h)
		result = append(result, int32(data.(Hotel).ID))
	}
	l, r := 0, len(result)-1
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}
	return result
}
