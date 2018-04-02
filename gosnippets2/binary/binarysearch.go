package main

import "fmt"

func main() {
	searchArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91} // note that array is sorted from low to high
	searchItem := 56
	_, numberOfSearches := BinarySearch(searchItem, searchArray)
	fmt.Println("Number of searches for ", searchItem, "was: ", numberOfSearches)

}

var High, Low, Mid int

func BinarySearch(s int, sa []int) (bool, int) {
	count := 1
	midPositionNumber := int((len(sa) - 1) / 2) // 5
	fmt.Println("midPositionNumber is: ", midPositionNumber)

	Low = sa[0]
	Mid = sa[midPositionNumber]
	High = sa[len(sa)-1]

	for s != Mid {
		if s == Mid || s == Low || s == High {
			return true, count
		}

		if s > Mid {
			Low = Mid
			midPositionNumber = int((midPositionNumber + len(sa)) / 2)
			Mid = sa[midPositionNumber]
			High = sa[len(sa)-1]
			count++
		}
		if s < Mid {
			High = Mid
			midPositionNumber = int(midPositionNumber / 2)
			Mid = sa[midPositionNumber]
			Low = sa[0]
			count++
		}
	}

	return true, count
}

/*

r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(1000))
	search := r.Intn(1000)
	fmt.Println("The random generated search item is: ", binarysearch(search))
*/
