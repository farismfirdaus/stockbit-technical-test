package main

import "sort"

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func anagramSolver(str []string) (result [][]string) {
	anagramContainer := map[string][]string{}
	for _, v := range str {
		runes := RuneSlice([]rune(v))
		sort.Sort(runes)
		if data, ok := anagramContainer[string(runes)]; data != nil && ok {
			anagramContainer[string(runes)] = append(anagramContainer[string(runes)], v)
		} else {
			anagramContainer[string(runes)] = []string{v}
		}
	}
	for _, v := range anagramContainer {
		result = append(result, v)
	}
	return
}
