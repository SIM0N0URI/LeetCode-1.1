func twoSum(nums []int, target int) []int {
    newmap:=make(map[int]int)
    for i,num:=range nums{
        take:=target-num
        if j,get:=newmap[take];get{
            return []int{j,i}
        }
        newmap[num]=i
    }

	return nil 
}

