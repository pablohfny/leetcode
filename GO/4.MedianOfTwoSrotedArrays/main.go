package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArray := []int{}
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				mergedArray = append(mergedArray, nums1[i])
				i++
			} else {
				mergedArray = append(mergedArray, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			mergedArray = append(mergedArray, nums1[i])
			i++
		} else {
			mergedArray = append(mergedArray, nums2[j])
			j++
		}
	}

	lengthMA := len(mergedArray)
	half := lengthMA / 2

	if lengthMA%2 == 0 {
		return float64((mergedArray[half-1] + mergedArray[half])) / 2

	} else {
		return float64(mergedArray[int(half)])
	}
}
