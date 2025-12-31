package trappingrainwater

func trap(height []int) int {
	// 鏈寸礌瑙ｆ硶鎬濊矾锛?姣忓垪鍌ㄦ按楂樺害鍙鑰冭檻姣忓垪鐨勫乏杈瑰拰鍙宠竟鐨勫偍姘撮珮搴﹀氨琛?
	// LeftMax,cur,RightMax
	// curWaterHeight= if(min(LeftMax,RightMax)>cur) min-cur
	// 鍔ㄦ€佽鍒? 涓嶇敤姣忔閬嶅巻閲岄潰鍐峟or鎵綥eftMax锛孯ightMax
	// 鍙互鍙亶鍘嗕竴娆℃妸缁撴灉鐢ㄦ暟缁勫瓨鍌ㄤ笅鏉?
	// 鍙屾寚閽堟槸鍔ㄦ€佽鍒掍笂浼樺寲锛?鎶婂乏鍙充袱杈规渶楂橀珮搴︿綅缃敤涓や釜鎸囬拡瀛樺偍灏卞彲浠ヤ簡

	res := 0
	LeftMax := 0
	RightMax := 0
	l := 1               //宸︽寚閽?鏈€宸﹀拰鏈€鍙冲瓨鍌ㄤ笉浜嗘按锛岀洿鎺ヤ粠 1閬嶅巻鍒?len(height)-2
	r := len(height) - 2 //鍙虫寚閽?

	// 涓や釜鎸囬拡閮藉悜涓棿闈犳嫝
	for l <= r {
		// 閲嶈锛佷袱杈瑰悓鏃跺紑宸?锛坙鎸囬拡锛夋渶宸LeftMax鏄繀瀹氬皬浜庣瓑浜庯紙r鎸囬拡锛夋渶鍙砵LeftMax锛堟渶鍙冲彸鐨勫乏鎷ユ湁鏇村ぇ鐨勮寖鍥达級
		// 鍙嶈繃鏉?锛坙鎸囬拡锛夋渶宸︾殑鍙砵RightMax 蹇呴』澶т簬绛変簬锛坮鎸囬拡锛夋渶鍙砵RightMax
		// 鐜板湪鎴戜滑鍙鑰冭檻涓よ竟鐨勭煭鏉跨殑灏卞彲浠ヤ簡锛堜篃灏辨槸鏈€宸?l鎸囬拡)鐨刬LeftMax鍜屾渶鍙?r鎸囬拡)鐨刯RightMax
		if height[l-1] < height[r+1] {
			LeftMax = max(LeftMax, height[l-1])
			cur := height[l]
			if LeftMax > cur {
				res += LeftMax - cur
			}
			l++
		} else {
			RightMax = max(RightMax, height[r+1])
			cur := height[r]
			if RightMax > cur {
				res += RightMax - cur
			}
			r--
		}
	}
	return res
}
