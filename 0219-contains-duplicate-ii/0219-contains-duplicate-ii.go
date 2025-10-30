import "fmt"
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func containsNearbyDuplicate(nums []int, k int) bool {
    newmap1:=make(map[int]int)
    for i,num:=range nums{
        if j,find:=newmap1[num];find{
            if abs(i-j)<=k{
                return true
            }
        }
        newmap1[num]=i
    }
    return false
}