package minimumwindow

func minWindow(s string, t string) string {
	need := [128]int{}   // 缁熻T
	window := [128]int{} // 缁熻绐楀彛鐨?
	needKinds := 0       // 闇€瑕佺殑瀛楃绉嶇被鏁?
	validKinds := 0      // 绗﹀悎鏉′欢鐨勫瓧绗︾绫绘暟
	res := ""            // 缁撴灉
	resLen := len(s) + 1 // 鐢ㄤ簬鍒ゆ柇鏈€鐭潯浠?

	// T 涓渶瑕佹弧瓒崇殑瀛楃绉嶇被鏁?
	for _, c := range t {
		// 鏌愪釜瀛楃绗竴娆″嚭鐜?needKinds++
		if need[c] == 0 {
			needKinds++
		}
		need[c]++
	}

	start := 0
	for end, ec := range s {
		//1.鏂板鍏冪礌
		window[ec]++
		// 闇€瑕佽瀛楃涓旀暟閲忎篃杈炬爣
		if need[ec] > 0 && need[ec] == window[ec] {
			//鏈変竴涓瓧绗︾绫绘暟閲忚揪鏍?+1
			//姣斿 A闇€瑕?涓紝绐楀彛鍐呭垰濂芥湁2涓紝杈炬爣
			validKinds++
		}

		//2.鏀剁缉绐楀彛 锛堟墍鏈夊瓧绗︾绫婚渶瑕佺殑鏁伴噺閮芥弧瓒宠姹傦級
		for validKinds == needKinds {
			//鏇存柊缁撴灉
			if end-start+1 < resLen {
				res = s[start : end+1]
				resLen = end - start + 1
			}
			sc := s[start]
			//鏀剁缉杈圭晫
			//棰勫垽鏂紝濡傛灉绐楀彛宸﹁竟鐣屽垰濂芥槸闇€瑕佺殑瀛楃绉嶇被涓旀暟閲忚揪鏍囷紝鍑虹獥鍙ｅ悗锛岃揪鏍囧瓧绗︾绫绘暟-1
			if need[sc] > 0 && need[sc] == window[sc] {
				//绐楀彛宸﹁竟鐣屽瓧姣嶅噺灏戝悗锛屽彲鑳藉氨涓嶈揪鏍囦簡
				validKinds--
			}
			window[sc]--
			start++
		}
	}

	//鎵句笉鍒?
	if resLen == len(s)+1 {
		return ""
	}
	return res
}
