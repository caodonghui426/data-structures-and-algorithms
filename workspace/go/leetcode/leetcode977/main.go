package main

func main() {
	testData := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(testData)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}

// sortedSquares 有序数组的平方
func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}

// sortedSquares2 排序法
func sortedSquares2(nums []int) []int {
    for i, val := range nums {
        nums[i] *= val
    }
    sort.Ints(nums)
    return nums
}

func sortedSquares3(nums []int) []int {
    if len(nums) == 1 {
        return []int{nums[0]*nums[0]}
    }
    // 1. 先找到分界点
    // 完全是负的，完全是正的，有正负分界点的
    l, r := 0, 1
    for r < len(nums) - 1 {
        if nums[l] < 0 && nums[r] < 0 {
            l++
            r++
        }
        if nums[l] < 0 && 0 <= nums[r] {
            break
        }
        if 0 <= nums[l] {
            break
        }
    }

	// 2. 从分界点开始，左右两边分别向两边扩散
    var result []int
    if l == 0 && r == 1 && 0 <= nums[0]{
        for i := 0; i < len(nums); i++ {
            result = append(result, nums[i]*nums[i])
        }
        return result
    }
    for 0 <= l && r < len(nums) {
        if nums[l]*nums[l] <= nums[r]*nums[r] {
            result = append(result, nums[l]*nums[l])
            l--
        } else {
            result = append(result, nums[r]*nums[r])
            r++
        }
    }
    if 0 <= l {
        for 0 <= l {
            result = append(result, nums[l]*nums[l])
            l--
        }
    }
    if r < len(nums) {
        for r < len(nums) {
            result = append(result, nums[r]*nums[r])
            r++
        }
    }
    return result
}