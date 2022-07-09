package daliuren

import "github.com/wlhyl/ganzhiwuxin"

// 天干的寄宫
func jiGong(tianGan ganzhiwuxin.TianGan) ganzhiwuxin.DiZhi {
	子, _ := ganzhiwuxin.NewDiZhi("子")
	甲, _ := ganzhiwuxin.NewTianGan("甲")

	寅 := 子.Plus(2)
	辰 := 寅.Plus(2)
	巳 := 辰.Plus(1)
	未 := 巳.Plus(2)
	申 := 未.Plus(1)
	戌 := 申.Plus(2)
	亥 := 戌.Plus(1)
	丑 := 亥.Plus(2)
	d := [...]ganzhiwuxin.DiZhi{寅, 辰, 巳, 未, 巳, 未, 申, 戌, 亥, 丑}
	n := tianGan.Minus(甲)
	return d[n]
}

// 获取某地支所寄的天干
func jiGongGan(d ganzhiwuxin.DiZhi) []ganzhiwuxin.TianGan {
	甲, _ := ganzhiwuxin.NewTianGan("甲")

	switch d.Name() {
	case "丑":
		return []ganzhiwuxin.TianGan{甲.Plus(-1)}
	case "寅":
		return []ganzhiwuxin.TianGan{甲}
	case "辰":
		return []ganzhiwuxin.TianGan{甲.Plus(1)}
	case "巳":
		return []ganzhiwuxin.TianGan{甲.Plus(2), 甲.Plus(4)}
	case "未":
		return []ganzhiwuxin.TianGan{甲.Plus(3), 甲.Plus(5)}
	case "申":
		return []ganzhiwuxin.TianGan{甲.Plus(6)}
	case "戌":
		return []ganzhiwuxin.TianGan{甲.Plus(7)}
	case "亥":
		return []ganzhiwuxin.TianGan{甲.Plus(-2)}
	default:
		return []ganzhiwuxin.TianGan{}
	}
}
