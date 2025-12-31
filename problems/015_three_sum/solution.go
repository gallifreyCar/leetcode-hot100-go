package threesum

import "sort"

func threeSum(nums []int) [][]int {

	res := [][]int{}
	// 鍏堟帓搴?
	sort.Ints(nums)

	// i 鏄涓€涓暟瀛?
	for i := 0; i < len(nums); i++ {
		l := i + 1         // 绗簩涓暟瀛?
		r := len(nums) - 1 // 绗笁涓暟瀛?
		// 绗竴涓暟瀛楀幓閲?
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			// 鎵句釜瑙?
			if nums[i]+nums[l]+nums[r] == 0 {
				// 鍏堝幓閲?鎵炬渶涓棿鐨勬暟锛?
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r-- // 澶ぇ浜?
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++ // 澶皬浜?
			}

		}
	}
	return res
}
