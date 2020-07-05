package Week04

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList,beginWord)
		ids[beginWord] = len(wordList) -1
	}
	if _, ok := ids[endWord]; !ok {
		return nil
	}
	n := len(wordList)
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if helper(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}
	var res [][]string
	cost := make([]int, n)
	queue := [][]int{{ids[beginWord]}}
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0
	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			temp := []string{}
			for _, index := range now {
				temp = append(temp, wordList[index])
			}
			res = append(res, temp)
		} else {
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					temp := make([]int,len(now))
					copy(temp,now)
					temp  = append(temp,to)
					queue = append(queue,temp)
				}
			}
		}
	}
	return res

}

func helper(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return str1[i+1:] == str2[i+1:]
		}
	}
	return false
}
