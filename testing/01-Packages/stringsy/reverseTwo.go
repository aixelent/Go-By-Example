package stringsy

func reverseString(st string) string {
	var rev string
	for i := len(st) - 1; i >= 0; i-- {
		rev += string(st[i])
	}
	return rev
}
