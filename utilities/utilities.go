package utilities

import "strings"
import "strconv"
import "io/ioutil"
import "fmt"

func ReadFileIntoLines(filename string) []string {
	b,err := ioutil.ReadFile(filename)
	if(err != nil){
		panic(err)
	}
	file := string(b)
	lines := strings.Split(file,"\n")
	return lines
}

func StringsToInts(strs []string) []int {
	var output = make([]int, len(strs))
	for i,s := range strs {
		temp,err := strconv.Atoi(s)
		if(err != nil) {
			panic(err)
		}
		output[i] = temp
	}

	return output
}

func IntsToString (input []int) string {
	var s string = fmt.Sprintf("%d", input[0] )
	for _,item := range input[1:] {
		s = fmt.Sprintf("%s,%d", s, item)
	}

	return s
}