package main

func getRow(rowIndex int) []int {
	arr := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		arr[i] = make([]int, 0, rowIndex+1)
		for j := 0; j < i+1; j++ {
			value := 1
			if i > 0 && j > 0 && j < i {
				value = arr[i-1][j-1] + arr[i-1][j]
			}
			arr[i] = append(arr[i], value)
		}
	}
	return arr[rowIndex]
}

func main() {

}
