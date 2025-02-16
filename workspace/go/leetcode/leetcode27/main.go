package main

func main() {
	testData := []int{3, 2, 2, 3}
	result := removeElement(testData, 3)
	for i := 0; i < result; i++ {
		println(testData[i])
	}
}

// removeElement 移除元素
func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    l, r := 0, len(nums) - 1
    for l <= r {
        if nums[l] == val {
            nums[l] = nums[r]
            nums[r] = -1
            r--
            continue
        }
        l++
    }
    return l
}