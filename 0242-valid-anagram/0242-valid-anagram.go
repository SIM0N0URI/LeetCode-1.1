import "fmt"
func isAnagram(s string, t string) bool {
    msp:=make(map[rune]int)
    if len(s)!=len(t){
        return false
    }
    for _,char:=range t{
        msp[char]++
    }
    for _,char2:=range s{
        msp[char2]--
        if msp[char2]<0{
            return false
        }
    }
    return true
}