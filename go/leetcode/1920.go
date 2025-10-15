func buildArray(nums []int) []int {
    ans := []int{}

    for i := range nums {
        ans = append(ans, nums[nums[i]])
    }

    return ans
}
