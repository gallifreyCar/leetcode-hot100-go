package firstmissingpositive

func firstMissingPositive(nums []int) int {
	//鍘熷湴绠楁硶锛岀敤涓嬫爣鏉ユ爣璁拌繖涓暟鏈夋病鏈夊嚭鐜拌繃
	//鐢ㄦ璐熸暟鏉ユ爣璁帮紙涓嬫爣i-1瀵瑰簲鏁板瓧i)

	// 鍘熸潵鐨?鎴栬礋鏁扮洿鎺ョ疆涓虹敤杈圭晫澶栫殑浠绘剰涓€涓鏁帮紙涓嶇敤澶勭悊
	for i, num := range nums {
		if num <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	//濡傛灉num鐨勪笅鏍噉um-1瀛樺湪锛屽垯缃负璐熸暟
	for _, num := range nums {
		value := abs(num) //杩欓噷鍑虹幇鐨勮礋鏁版槸鍥犱负鍓嶇疆鎿嶄綔瀵艰嚧鐨?
		if value-1 < len(nums) {
			nums[value-1] = -abs(nums[value-1])
		}
	}

	//瀛樺湪姝ｆ暟灏辨槸缂哄皯鐨勯偅涓暟瀛?
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
