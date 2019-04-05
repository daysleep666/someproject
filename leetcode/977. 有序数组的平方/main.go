package main

func sortedSquares(A []int) []int {
	i := 0
	for i, _ = range A {
		if A[i] >= 0 {
			break
		}
	}
	reverse(A, 0, i-1)
	newA := make([]int, len(A))
	k, a, b := 0, 0, i
	for a < i && b < len(A) {
		if square(A[a]) > square(A[b]) {
			newA[k] = square(A[b])
			k++
			b++
		} else {
			newA[k] = square(A[a])
			k++
			a++
		}
	}
	for a < i {
		newA[k] = square(A[a])
		k++
		a++
	}
	for b < len(A) {
		newA[k] = square(A[b])
		k++
		b++
	}
	return newA
}

func reverse(nums []int, from, to int) {
	for ; from < to; from, to = from+1, to-1 {
		nums[from], nums[to] = nums[to], nums[from]
	}
}

func square(a int) int {
	return a * a
}

func main() {

}
