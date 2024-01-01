// Упражнение 2.3. Перепишите функцию PopCount так, чтобы она использовала 
// цикл вместо единого выражения.

package popcount
// pc[i] - количество единичных битов в i.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count 
// (number of set bits) of the value x. 
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}