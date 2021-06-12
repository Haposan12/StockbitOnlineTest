package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	fmt.Println(anagram(input))
}

func anagram(s []string) (result [][]string) {
	anagramMap := make(map[string][]string)
	for _, v := range s {
		//split string
		split := strings.Split(v, "")
		sort.Strings(split)

		anagramMap[strings.Join(split, "")] = append(anagramMap[strings.Join(split, "")], v)
	}

	for _, v := range anagramMap {
		result = append(result, v)
	}

	return result
}

/*Solution explanation
	 - Create a map with key string and value array of string
	 - Looping from string input
	 - Split string with separation ""
	 - Then sort the split string ascending order
	 - Then insert into map
	 - Result like this = [aakmn:[makan] aik:[kia] aikt:[kita atik tika] aku:[aku kua]]
	 - Then append value in map into returning string
*/
