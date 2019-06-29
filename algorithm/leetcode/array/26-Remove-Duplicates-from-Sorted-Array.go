func removeDuplicates(nums []int) int {
    var cur int = 0
    var curValue int = 0
    for i,v := range nums {
        if i == 0 {
            curValue = v
            continue
        }
        if v == curValue {
            continue
        }
        cur = cur + 1
        nums[cur] = v
        curValue = v
    }
    return cur+1
}
