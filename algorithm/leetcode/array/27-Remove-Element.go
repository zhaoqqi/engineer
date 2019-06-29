func removeElement(nums []int, val int) int {
    var curPos int = 0
    for _, v := range nums {
        if v == val {
            continue
        }
        nums[curPos] = v
        curPos += 1
    }
    return curPos
}
