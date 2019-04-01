package main

func peakIndexInMountainArray(A []int) int {
	i := 0
	for ; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			break
		}
	}
	return i
}

func main() {

}
