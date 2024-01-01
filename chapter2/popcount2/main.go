// Выражение х&(х-1) сбрасывает крайний справа ненулевой бит х. 
// Напишите версию PopCount, которая подсчитывает биты с использованием 
// этого факта 

package popcount
// pc[i] - количество единичных битов в i.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount counts the number of set bits in a uint64 using the expression x&(x-1)
func PopCount(x uint64) int {
	count := 0
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
