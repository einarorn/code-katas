package merge_lists

func SortTwoLists(list1, list2 []int) []int {
	mergedList := make([]int, 0, len(list1)+len(list2))

	for {
		if len(list1)+len(list2) == 0 {
			return mergedList
		}

		if len(list1) == 0 {
			mergedList, list2 = appendAndPop(mergedList, list2, true)
			continue
		}

		if len(list2) == 0 {
			mergedList, list1 = appendAndPop(mergedList, list1, true)
			continue
		}

		if list1[0] < list2[0] {
			mergedList, list1 = appendAndPop(mergedList, list1, false)
		} else {
			mergedList, list2 = appendAndPop(mergedList, list2, false)
		}

	}
}

func appendAndPop(merged, list []int, popAll bool) ([]int, []int) {
	if popAll {
		merged = append(merged, list...)
		list = []int{}
	} else {
		merged = append(merged, list[0])
		list = list[1:]
	}
	return merged, list
}
