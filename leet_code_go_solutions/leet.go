package leet_code_go_solutions

import (
	"fmt"
	"sort"
)

func hasPairWithSum(a []int, sumVal int) []int {

	var result []int
	cur := 0
	nxt := 1
	for i := 0; i < len(a); i++ {

		if (a[cur] + a[nxt]) == sumVal {
			result = append(result, cur, nxt)
			return result
		} else {
			if nxt != (len(a) - 1) {
				cur = cur + 1
				nxt = nxt + 1
			}
		}
	}
	return result
}

//func checkMe(one []string, two map[string]bool) bool {
//	for _, i := range one {
//		fmt.Println("Val:%s", i)
//		if two[i] {
//			return true
//		}
//	}
//	return false
//}

func robotPositions(roboPositions [][]int) map[int]int {

	var result map[int]int
	for i := range roboPositions {
		if i+1 < len(roboPositions) {
			initialArray := roboPositions[i]
			nextArray := roboPositions[i+1]
			countOfNxtArray := len(nextArray)
			if countOfNxtArray > 0 {
				for j := i; j < len(initialArray); j++ {
					currVal := initialArray[j]
					if currVal == 1 {
						//now that we have found the robot we need to check verify the movement
						isValid := checkItsPosition(currVal, j, nextArray)
						if !isValid {
							fmt.Println("Val:", currVal, j)
						}
					}

				}
			}

		}
	}
	return result

}

func checkItsPosition(val, position int, arrayToCheckAgainst []int) bool {

	if len(arrayToCheckAgainst) > position {

		currPosVal := arrayToCheckAgainst[position]
		//if current position doesn't have same value it means the robot has moved
		if currPosVal != val && (position+1) < len(arrayToCheckAgainst) {
			currPosVal = arrayToCheckAgainst[position+1]
		}
		if currPosVal != val && (position-1) < len(arrayToCheckAgainst) {
			currPosVal = arrayToCheckAgainst[position-1]
		}
		if currPosVal == val {
			return true
		}

	}
	return false
}

//Array Related
func add(val int, arr []int) []int {
	arr = append(arr, val)
	return arr
}

//with order change
func remove(index int, arr []int) []int {
	var temp = arr[len(arr)-1]
	arr[index] = temp

	return arr[:len(arr)-1]
}

// 9  1 5 3 _ 7 2
func removeWithOrder(index int, arr []int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func reverseString(s string) string {
	var newString []string
	for i := len(s) - 1; i >= 0; i-- {
		newString = append(newString, string(s[i]))
	}
	return fmt.Sprintf("%s", newString)
}

//[0,3,4,31] [4,6,30] // n log(n)
func mergeSortedArray(a, b []int) {
	result := append(a, b...)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	fmt.Println(result)

}

func maxSubArray(nums []int) {
	// -2 1 -3 4 -1 2 1 -5 4
	length := len(nums)
	max_sum := nums[0]
	sum := nums[0]
	fmt.Println(max_sum)
	fmt.Println(sum)
	for i := 1; i < length; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		fmt.Println(max_sum, sum)
		fmt.Println()
		if sum > max_sum {
			max_sum = sum
		}
		fmt.Println("2tttt->", max_sum, sum)
	}
	fmt.Println(max_sum)

}

func maxSubArrayDivideConquer(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums, 0, right)
}

func findMaxSubArray(nums []int, left, right int) int {
	// -2 1 -3 4 -1 2 1 -5 4
	if left == right {
		return nums[left]
	} else {
		mid := (left + right) / 2
		leftMax := findMaxSubArray(nums, left, mid)
		rightMax := findMaxSubArray(nums, mid+1, right)
		crossMax := maxCrossing(nums, left, right, mid)
		return max(leftMax, rightMax, crossMax)
	}

}

func max(values ...int) int {
	maxVal := values[0]
	for i, val := range values {
		if maxVal < values[i] {
			maxVal = val
		}
	}
	return maxVal
}

func maxCrossing(nums []int, left, right, mid int) int {

	var sum = 0
	leftSum := 0
	rightSum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}

//func moveZeroes(nums []int) []int {
//
//	var newNums []int
//	var count []int
//	for i := 0; i < len(nums); i++ {
//
//		if nums[i] != 0 {
//			newNums = append(newNums, nums[i])
//
//		} else {
//			count = append(count, 0)
//		}
//	}
//
//	newNums = append(newNums, count...)
//	fmt.Println(newNums)
//	return newNums
//}
func moveZeroes(nums []int) []int {

	var count = 0
	var j = 0
	for i := 0; i < len(nums); i++ {
		fmt.Println("check1:", nums[i])
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++

		} else {
			count++
		}
	}

	newj := len(nums) - count
	for i := 0; i < count; i++ {
		nums[newj] = 0
		newj++
	}

	return nums
}

func facto(num int) int {
	if num == 1 || num == 0 {
		return num
	}

	return num * facto(num-1)

}

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)

}

func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func wordBreak(s string, wordDict []string) bool {
	// initialize []bool containing where solved words live
	// zero value of bool in Go is false
	dp := make([]bool, len(s)+1)

	for i := 0; i < len(s); i++ {
		// to speed up the processing
		// no need to iterate on mid-word letters
		// except the 1st index
		if i > 0 {
			if !dp[i] {
				continue
			}
		}
		// strategy to find the last index of the dictionary words
		// get length of word, using the two idx, see if that word
		// matches in the string
		for _, w := range wordDict {
			j := i + len(w)
			if j <= len(s) && w == s[i:j] {
				dp[j] = true
			}
		}
	}
	return dp[len(s)]
}

type ListNode struct {
	Val  string
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var out = newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			l1 = l1.Next
			newList = newList.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
			newList = newList.Next
		}
	}
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return out.Next
}
