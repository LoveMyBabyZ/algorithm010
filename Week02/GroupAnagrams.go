package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAngarams(str))
}

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i,j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func groupAngarams(str []string) [][]string {
	ret := [][]string{}
	if len(str) == 0{
		return ret
	}
	keyMap := make(map[string]int)
	for _,s := range str{
		strByte := bytes(s)
		sort.Sort(strByte)
		key := string(strByte)
		if i ,ok := keyMap[key];!ok{
			keyMap[key] = len(ret)
			ret = append(ret, []string{s})
		} else {
			ret[i] = append(ret[i], s)
		}

	}
	return  ret
}
