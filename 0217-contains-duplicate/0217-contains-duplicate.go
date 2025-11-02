func containsDuplicate(nums []int) bool {
    mapos := make(map[int]bool)

    for _, n := range nums{
        if mapos[n] != true{
            mapos[n] = true
        }else {
            return true
        }
    }
    return false
}
