package main

func toLowerCase(str string) string {
	result := ""
	for i, _ := range str {
		result += toLower(str[i])
	}
	return result
}

func toLower(v byte) string {
	if 'a' <= v && v <= 'z' {
		return string(v)
	} else if 'A' <= v && v <= 'Z' {
		return string((v - 'A') + 'a')
	}
	return string(v)
}

func main() {

}
