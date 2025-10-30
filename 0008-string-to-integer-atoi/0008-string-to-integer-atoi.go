func myAtoi(s string) int {
    s = strings.TrimSpace(s)

    if len(s) == 0 {
        return 0
    }

    var negative bool
    
    if s[0] == '-' {
        negative = true
        s = s[1:]
    } else if s[0] == '+' {
        s = s[1:]
    }

    var res int

    for i, c := range s {
        if c < '0' || c > '9' {
            break
        }

        res = res * 10 + int(s[i]-'0')

        if res > 2147483647 && !negative {
            return 2147483647
        }

        if -res < -2147483648 && negative{
            return -2147483648
        }
    }

    if negative {
        res = -res
    }

    return res
}