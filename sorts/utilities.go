package sorts

//GenerateRandomNumbers es una función generadora de números pseudoaleatorios
func GenerateRandomNumbers(size int, seed int) []int {

	m := 2048 //Periodo
	a := 109  //Multiplicador
	c := 0    //Incremento = 0, es Método de Congruencia Lineal Multiplicativa

	slice := make([]int, size, size)
	slice[0] = formula(a, seed, c, m)
	for i := 1; i < size; i++ {
		slice[i] = formula(a, slice[i-1], c, m)
	}
	return slice
}

func formula(a int, xn int, c int, m int) int {
	return (a*xn + c) % m % 31
}
