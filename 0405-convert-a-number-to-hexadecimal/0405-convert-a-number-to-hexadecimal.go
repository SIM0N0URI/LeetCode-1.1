// toHex converts an integer to its hexadecimal representation (lowercase).
// Handles negative numbers using two's complement (32-bit unsigned).
func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    n, ref := uint32(num), "0123456789abcdef"
    res := ""
    for n != 0 {
        res = string(ref[n&0xF]) + res
        n >>= 4
    }
    return res
}