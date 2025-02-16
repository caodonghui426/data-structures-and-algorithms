package main

func main() {
	testData := []int{-1, 0, 3, 5, 9, 12}
	result := search(testData, 9)
	println(result)
}

// search 二分查找
func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        m := (r - l)/2 + l
        if target == nums[m] {
            return m
        } else if target < nums[m] {
            r = m - 1
        }else{
            l = m + 1
        }
    }
    return -1
}