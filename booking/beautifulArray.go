package booking

import "fmt"

func BeautifulSubarrays1(arr []int32, numOdds int32) int64 {
	// Write your code here
	dp := make([]int, 0)
	count, result := 0, 0
	for i := 0; i < len(arr); i++ {
		count++
		if arr[i]%2 == 1 {
			dp = append(dp, count)
			count = 0
		}
		if len(dp) >= int(numOdds) {
			result += dp[len(dp)-int(numOdds)]
		}
	}
	return int64(result)
}

func beautifulSubarrays(arr []int32, numOdds int32) int64 {
	// Write your code here
	left, right, oddCount, res := 0, 0, 0, 0
	for right < len(arr) {
		if arr[right]&1 == 1 {
			oddCount++ // 奇数
		}
		right++
		if oddCount == int(numOdds) {
			// right向右直到找到奇数或越界停下
			tmpRight := right // 记下当前right位置
			for right < len(arr) && arr[right]&1 == 0 {
				right++
			}
			rightEvenCount := right - tmpRight
			// left向左找到奇数或越界之后停下
			leftEvenCount := 0
			for arr[left]&1 == 0 {
				left++ // left往右走，到第一个奇数时停下
				leftEvenCount++
			}
			res += (leftEvenCount + 1) * (rightEvenCount + 1)
			// 此时left为第一个奇数的位置，left++就在左侧让出一个奇数
			left++
			// 左侧少了一个奇数，所以奇数的计数器要减一
			oddCount--
		}
	}
	return int64(res)
}

func BeautifulSubarrays3(arr []int32, numOdds int32) int64 {
	list := make([]int32, 0)
	beautifulSubarraysTraverse(0, numOdds, arr, &list)
	return 0
}

func beautifulSubarraysTraverse(index, numOdds int32, arr []int32, list *[]int32) {
	if numOdds == 0 {
		fmt.Println(list)
		return
	}

	for i := index; i < int32(len(arr)); i++ {
		if arr[i]%2 == 1 {
			numOdds--
		}
		*list = append(*list, arr[i])
		beautifulSubarraysTraverse(i+1, numOdds, arr, list)
		*list = (*list)[:len(*list)-1]
		if arr[i]%2 == 1 {
			numOdds++
		}
	}
}
