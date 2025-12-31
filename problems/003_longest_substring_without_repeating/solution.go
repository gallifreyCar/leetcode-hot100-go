package longestsubstring

func lengthOfLongestSubstring(s string) int {
	// 婊戝姩绐楀彛
	// start涓嶅姩锛宔nd鍚戝悗绉诲姩
	// 褰揺nd閬囧埌閲嶅瀛楃锛宻tart搴旇鏀惧湪涓婁竴涓噸澶嶅瓧绗︾殑浣嶇疆鐨勫悗涓€浣嶏紝鍚屾椂璁板綍鏈€闀跨殑闀垮害
	// key鍒ゆ柇鏄惁閲嶅锛寁alue鐢ㄦ壘鍒伴噸澶嶅瓧绗︾殑涓嬩竴涓綅缃?

	res := 0
	m := make(map[byte]int)
	for start, end := 0, 0; end < len(s); end++ {
		c := s[end]
		//濡傛灉褰撳墠瀛楃宸茬粡鍑虹幇杩囷紝鑰屼笖鍦ㄦ垜浠獥鍙ｅ唴锛屽垯涓鸿涓洪噸澶嶅瓧绗?
		if loc, ok := m[c]; ok && loc >= start {
			start = loc + 1
		}
		m[c] = end
		res = max(res, end-start+1)
	}
	return res
}
