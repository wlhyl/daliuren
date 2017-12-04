package shipan

import "ganzhiwuxin"

type TianJiang struct {
	天乙所乘地支 ganzhiwuxin.DiZhi
	顺排     bool
}

func (this *TianJiang) Init(四课 SiKe, 昼占 bool) {
	// 天干数取甲为1，地支数，取寅为1
	昼贵 := [...]int{0, 6, 7, 8, 10, 12, 11, 12, 1, 2, 4}
	夜贵 := [...]int{0, 12, 11, 10, 8, 6, 7, 6, 5, 4, 2}

	if 昼占 {
		var d ganzhiwuxin.DiZhi
		d.Init(昼贵[四课.gan.GetNum()])
		this.天乙所乘地支 = d
	} else {
		var d ganzhiwuxin.DiZhi
		d.Init(夜贵[四课.gan.GetNum()])
		this.天乙所乘地支 = d
	}

	// 计算顺排或是逆排 天乙所临地支>=4 and <=9，即 在巳与戌间，为逆排
	天乙所临地支 := 四课.tp.Get临(this.天乙所乘地支)
	if 天乙所临地支.GetNum() >= 4 && 天乙所临地支.GetNum() <= 9 {
		this.顺排 = false
	} else {
		this.顺排 = true
	}
}
func (this *TianJiang) Get天将(d ganzhiwuxin.DiZhi) string {
	数字映射名字 := [...]string{"贵", "蛇", "雀", "合", "勾", "龙", "空", "虎", "常", "玄", "阴", "后"}

	if this.顺排 {
		return 数字映射名字[(d.Sub(this.天乙所乘地支)+12)%12]
	}
	return 数字映射名字[(12-(d.Sub(this.天乙所乘地支)+12)%12)%12]
}
