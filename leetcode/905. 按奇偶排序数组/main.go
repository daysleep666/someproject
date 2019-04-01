package main

func sortArrayByParity(A []int) []int {
	j := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			j++
		}
	}
	return A
}

func main() {

}
