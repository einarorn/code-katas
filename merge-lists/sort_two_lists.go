package merge_lists

func SortTwoLists(list1, list2 []int) []int {
	mergedList := make([]int, 0, len(list1)+len(list2))

	for {
		if len(list1)+len(list2) == 0 {
			return mergedList
		}

		switch true {
		case len(list1) == 0:
			mergedList, list2 = appendAndPop(mergedList, list2)
		case len(list2) == 0:
			mergedList, list1 = appendAndPop(mergedList, list1)
		default:
			if list1[0] < list2[0] {
				mergedList, list1 = appendAndPop(mergedList, list1)
			} else {
				mergedList, list2 = appendAndPop(mergedList, list2)
			}
		}
	}
}

func appendAndPop(merged, list []int) ([]int, []int) {
	merged = append(merged, list[0])
	return merged, list[1:]
}
