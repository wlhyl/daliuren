package daliuren

import (
	"encoding/json"

	"github.com/wlhyl/ganzhiwuxin"
)

// 三传
type SanChuan struct {
	// sike           SiKe
	chu, zhong, mo  ganzhiwuxin.DiZhi
	dunGan, liuQing [3]string
}

func NewSanChuan(tianPan TianPan, sike SiKe) SanChuan {
	// 三传
	var sc SanChuan
	sc.chu, sc.zhong, sc.mo = sanChuan(tianPan, sike)

	// 遁干
	甲, _ := ganzhiwuxin.NewTianGan("甲")
	gan := sike.gan
	zhi := sike.zhi
	delta := gan.Minus(甲)

	xunShou := zhi.Plus(-1 * delta)

	for index, d := range []ganzhiwuxin.DiZhi{sc.chu, sc.zhong, sc.mo} {
		zhiDelta := d.Minus(xunShou)
		if zhiDelta == 10 || zhiDelta == 11 {
			sc.dunGan[index] = ""
		} else {
			sc.dunGan[index] = 甲.Plus(zhiDelta).Name()
		}

		if gan.WuXing().Ke(d.WuXing()) {
			sc.liuQing[index] = "财"
		} else if d.WuXing().Ke(gan.WuXing()) {
			sc.liuQing[index] = "官"
		} else if gan.WuXing().Sheng(d.WuXing()) {
			sc.liuQing[index] = "子"
		} else if d.WuXing().Sheng(gan.WuXing()) {
			sc.liuQing[index] = "父"
		} else {
			sc.liuQing[index] = "兄"
		}
	}

	return sc
}

func (sc SanChuan) MarshalJSON() ([]byte, error) {

	m := map[string]interface{}{
		"chu":     sc.chu.Name(),
		"zhong":   sc.zhong.Name(),
		"mo":      sc.mo.Name(),
		"dunGan":  sc.dunGan,
		"liuQing": sc.liuQing,
	}
	return json.Marshal(m)
}

func sanChuan(tianpan TianPan, sike SiKe) (ganzhiwuxin.DiZhi, ganzhiwuxin.DiZhi, ganzhiwuxin.DiZhi) {
	if sike.zhiYang.Equals(sike.zhi) {
		return get伏呤(sike)
	}

	if sike.zhiYang.LiuChong(sike.zhi) {
		sc := get反呤(tianpan, sike)
		return sc[0], sc[1], sc[2]
	}

	if sc := get贼克(tianpan, sike); len(sc) != 0 {
		return sc[0], sc[1], sc[2]
	}

	if sike.ganYang.Equals(sike.zhiYang) {
		sc := get八专(sike)
		return sc[0], sc[1], sc[2]
	}

	if sc := get遥克(tianpan, sike); len(sc) != 0 {
		return sc[0], sc[1], sc[2]
	}

	//确认是昴星课
	四课上神组 := [4]ganzhiwuxin.DiZhi{sike.ganYang, sike.ganYing, sike.zhiYang, sike.zhiYing}

	无重复的课 := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range 四课上神组 {
		has := false
		for _, v1 := range 无重复的课 {
			if v.Equals(v1) {
				has = true
			}
		}

		if !has {
			无重复的课 = append(无重复的课, v)
		}
	}

	if len(无重复的课) == 4 {
		sc := get昴星(tianpan, sike)
		return sc[0], sc[1], sc[2]
	}
	sc := get别责(tianpan, sike)
	return sc[0], sc[1], sc[2]

}

//有下克上的课
// 返回有贼的课序数,1,2,3,4课
func has贼(sike SiKe) []int {
	var keList []int
	// 有克的课 := make([]ganzhiwuxin.DiZhi, 0)
	// 四课 := sc.sike
	if sike.gan.WuXing().Ke(sike.ganYang.WuXing()) {
		keList = append(keList, 1)
	}
	if sike.ganYang.Ke(sike.ganYing) {
		keList = append(keList, 2)
	}
	if sike.zhi.Ke(sike.zhiYang) {
		keList = append(keList, 3)
	}
	if sike.zhiYang.Ke(sike.zhiYing) {
		keList = append(keList, 4)
	}

	if len(keList) == 0 {

		return []int{}
	}

	//删除重复课
	var shangShenList []ganzhiwuxin.DiZhi
	var keListTmp []int
	for _, v := range keList {
		has := false
		shangShen := getShangShen(sike, v)
		for _, d := range shangShenList {
			if shangShen.Equals(d) {
				has = true
			}
		}
		if !has {
			shangShenList = append(shangShenList, shangShen)
			keListTmp = append(keListTmp, v)
		}
	}

	return keListTmp
}

//有上克下的课
func has克(sike SiKe) []int {
	var keList []int

	if sike.ganYang.WuXing().Ke(sike.gan.WuXing()) {
		keList = append(keList, 1)
	}
	if sike.ganYing.Ke(sike.ganYang) {
		keList = append(keList, 2)
	}
	if sike.zhiYang.Ke(sike.zhi) {
		keList = append(keList, 3)
	}
	if sike.zhiYing.Ke(sike.zhiYang) {
		keList = append(keList, 4)
	}

	if len(keList) == 0 {

		return keList
	}

	//删除重复课
	var shangShenList []ganzhiwuxin.DiZhi
	var keListTmp []int
	for _, v := range keList {
		has := false
		shangShen := getShangShen(sike, v)
		for _, d := range shangShenList {
			if shangShen.Equals(d) {
				has = true
			}
		}
		if !has {
			shangShenList = append(shangShenList, shangShen)
			keListTmp = append(keListTmp, v)
		}

	}
	return keListTmp
}

func get贼克(tianPan TianPan, sike SiKe) []ganzhiwuxin.DiZhi {
	贼 := has贼(sike)
	克 := has克(sike)

	if len(贼) == 1 {
		chu := getShangShen(sike, 贼[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)
		return []ganzhiwuxin.DiZhi{chu, zhong, mo}
	}

	if len(贼) > 1 {
		sc := get比用(tianPan, sike, 贼)
		return sc[:]

	}

	if len(克) == 1 {
		chu := getShangShen(sike, 克[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)
		return []ganzhiwuxin.DiZhi{chu, zhong, mo}
	}
	if len(贼) > 1 {
		sc := get比用(tianPan, sike, 克)
		return sc[:]
	}

	return []ganzhiwuxin.DiZhi{}
}

func get比用(tianPan TianPan, sk SiKe, keList []int) [3]ganzhiwuxin.DiZhi {
	/* len(results) ==1, >1, ==0
	   results中保存 取比用後的结果
	*/
	results := make([]int, 0)
	for _, it := range keList {
		if getShangShen(sk, it).Masculine() == sk.gan.Masculine() {
			results = append(results, it)
		}
	}

	if len(results) == 1 {
		chu := getShangShen(sk, results[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)
		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	}
	// 俱不比
	if len(results) == 0 {
		return get涉害(tianPan, sk, keList)
	}
	//多个俱比 len(result) >1
	return get涉害(tianPan, sk, results)

}

func get涉害(tianPan TianPan, sike SiKe, keList []int) [3]ganzhiwuxin.DiZhi {
	var is贼 = func(it int) bool {
		if it == 1 {
			return sike.gan.WuXing().Ke(sike.ganYang.WuXing())
		}
		if it == 2 {
			return sike.ganYang.Ke(sike.ganYing)
		}
		if it == 3 {
			return sike.zhi.Ke(sike.zhiYang)
		}
		return sike.zhiYang.Ke(sike.zhiYing)

	}

	//使用涉归，不用孟仲季取法
	// __课的涉害深度中的key存储地支的num，value存储涉害深度
	var __课的涉害深度 = make(map[int]int)
	for _, v := range keList {
		课 := getShangShen(sike, v)
		临地盘 := tianPan.down(课)

		count := 0
		for i := 0; i < 课.Minus(临地盘); i++ {
			__d := 临地盘.Plus(i)
			jiGan := jiGongGan(__d)
			if is贼(v) {
				if __d.Ke(课) {
					count++
				}
				for _, g := range jiGan {
					if g.WuXing().Ke(课.WuXing()) {
						count++
					}
				}

			} else {
				if 课.Ke(__d) {
					count++
				}
				for _, g := range jiGan {
					if 课.WuXing().Ke(g.WuXing()) {
						count++
					}
				}
			}
		}
		__课的涉害深度[v] = count
	}

	最大涉害深度 := 0
	var 有最大涉害深度的支组 []int
	for _, v := range __课的涉害深度 {
		if v > 最大涉害深度 {
			最大涉害深度 = v
		}
	}

	for index, v := range __课的涉害深度 {
		if v == 最大涉害深度 {
			// var 有最大涉害深度的支 ganzhiwuxin.DiZhi
			// 有最大涉害深度的支.Init(k)
			有最大涉害深度的支组 = append(有最大涉害深度的支组, index)
		}
	}

	if len(有最大涉害深度的支组) == 1 {

		chu := getShangShen(sike, 有最大涉害深度的支组[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)

		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}

	}

	//涉害深度相同
	子, _ := ganzhiwuxin.NewDiZhi("子")
	寅 := 子.Plus(2)
	// 找出四孟
	var 四孟支组 []int
	for _, v := range 有最大涉害深度的支组 {
		临地盘 := tianPan.down(getShangShen(sike, v))

		if 临地盘.Minus(寅)%3 == 0 {
			四孟支组 = append(四孟支组, v)
		}
	}

	if len(四孟支组) == 1 {
		chu := getShangShen(sike, 四孟支组[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)

		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	}

	if len(四孟支组) > 0 {
		if sike.gan.Masculine() {
			chu := sike.ganYang
			zhong := tianPan.up(chu)
			mo := tianPan.up(zhong)

			return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
		} else {
			chu := sike.zhiYang
			zhong := tianPan.up(chu)
			mo := tianPan.up(zhong)
			return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
		}
	}

	// 从仲发用
	var 四仲支组 []int
	for _, v := range 有最大涉害深度的支组 {
		临地盘 := tianPan.down(getShangShen(sike, v))

		if 临地盘.Minus(子)%3 == 0 {
			四仲支组 = append(四仲支组, v)
		}
	}

	if len(四仲支组) == 1 {
		chu := getShangShen(sike, 四仲支组[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)

		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	}
	// 剩余情况有两种:
	// len(四仲支组) > 1
	// len(四仲支组) == 0，此种情况，支组临于四季，与前一种情况一并按“复等课”取三传
	if sike.gan.Masculine() {
		chu := sike.ganYang
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)

		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	} else {
		chu := sike.zhiYang
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)
		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	}
}

func get遥克(tianPan TianPan, sike SiKe) []ganzhiwuxin.DiZhi {

	var ke []int
	if sike.ganYing.WuXing().Ke(sike.gan.WuXing()) {
		ke = append(ke, 2)

	}
	if sike.zhiYang.WuXing().Ke(sike.gan.WuXing()) {
		ke = append(ke, 3)
	}
	if sike.zhiYing.WuXing().Ke(sike.gan.WuXing()) {
		ke = append(ke, 4)
	}

	if len(ke) == 0 {
		if sike.gan.WuXing().Ke(sike.ganYing.WuXing()) {
			ke = append(ke, 2)
		}
		if sike.gan.WuXing().Ke(sike.zhiYang.WuXing()) {
			ke = append(ke, 3)
		}
		if sike.gan.WuXing().Ke(sike.zhiYing.WuXing()) {
			ke = append(ke, 4)
		}
	}
	if len(ke) == 0 {
		return []ganzhiwuxin.DiZhi{}
	}

	keList := make([]int, 0)
	for _, v := range ke {
		has := false
		k0 := getShangShen(sike, v)
		for _, v1 := range keList {
			k1 := getShangShen(sike, v1)
			if k0.Equals(k1) {
				has = true
			}
		}

		if !has {
			keList = append(keList, v)
		}
	}

	if len(keList) == 1 {
		chu := getShangShen(sike, keList[0])
		zhong := tianPan.up(chu)
		mo := tianPan.up(zhong)
		return []ganzhiwuxin.DiZhi{chu, zhong, mo}
	}
	// 如果有两课遥克，取比用，遥克课取比用之後，不会出现涉害（需要验证）
	sc := get比用(tianPan, sike, keList)
	return sc[:]
}

// 不判断是否为昴星课，只按昴星课取三传
func get昴星(tianPan TianPan, sike SiKe) [3]ganzhiwuxin.DiZhi {

	酉, _ := ganzhiwuxin.NewDiZhi("酉")
	if sike.gan.Masculine() {
		chu := tianPan.up(酉)
		zhong := sike.zhiYang
		mo := sike.ganYang
		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	} else {
		chu := tianPan.down(酉)
		zhong := sike.ganYang
		mo := sike.zhiYang
		return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
	}
}

// 不判断是否为别责课，只按别责取传
func get别责(tianpan TianPan, sike SiKe) [3]ganzhiwuxin.DiZhi {
	var d ganzhiwuxin.DiZhi
	if sike.gan.Masculine() {
		d = jiGong(sike.gan.Plus(5))
	} else {
		//干为阴的情况
		d = sike.zhi.Plus(4)
	}
	//中末皆取干上神
	// 阴日，取三合上神为初传，不取三合为初传，《六壬直指》原文的括号中的注解似乎有错
	chu := tianpan.up(d)
	zhong := sike.ganYang
	mo := sike.ganYang
	return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
}

// 调用此函数前，需要先确定是八专课，此函数只按八专取三传，不判断是否为八专课
func get八专(sike SiKe) [3]ganzhiwuxin.DiZhi {

	if sike.gan.Masculine() {
		chu := sike.ganYang.Plus(2)
		return [3]ganzhiwuxin.DiZhi{chu, sike.ganYang, sike.ganYang}
	} else {
		chu := sike.zhiYing.Plus(-2)
		return [3]ganzhiwuxin.DiZhi{chu, sike.ganYang, sike.ganYang}
	}
}

func get伏呤(sike SiKe) (ganzhiwuxin.DiZhi, ganzhiwuxin.DiZhi, ganzhiwuxin.DiZhi) {
	甲, _ := ganzhiwuxin.NewTianGan("甲")
	乙, 癸 := 甲.Plus(1), 甲.Plus(-1)

	var chu, zhong, mo ganzhiwuxin.DiZhi
	// 六乙、六癸日 无克，阳日取干上神发用
	if sike.gan.Equals(乙) || sike.gan.Equals(癸) || sike.gan.Masculine() {
		chu = sike.ganYang
	} else {
		// 阴日，非六乙日、六癸
		chu = sike.zhiYang
	}

	//求得中传
	for i := 0; i < 12; i++ {
		d := chu.Plus(i)
		if chu.Xing(d) {
			zhong = d
			break
		}
	}
	// 初为自刑，阳日、六乙日、六癸日取支上神为中传
	if chu.Equals(zhong) {
		if sike.gan.Equals(乙) || sike.gan.Equals(癸) || sike.gan.Masculine() {
			//六乙、六癸、阳日，初传传自刑，取支上神为中传
			zhong = sike.zhiYang
		} else {
			zhong = sike.ganYang
		}
	}

	//求末传
	for i := 0; i < 12; i++ {
		d := chu.Plus(i)
		if zhong.Xing(d) {
			mo = d
			break
		}
	}

	//中传自刑，取中所冲之神
	if zhong.Equals(mo) {
		mo = zhong.Plus(6)
	}

	// 初、中互刑，如：子、卯，末取中所冲之神
	if zhong.Xing(chu) {
		mo = zhong.Plus(6)
	}
	return chu, zhong, mo
}

func get反呤(tianPan TianPan, sike SiKe) [3]ganzhiwuxin.DiZhi {
	sc := get贼克(tianPan, sike)
	if len(sc) != 0 {
		var s [3]ganzhiwuxin.DiZhi
		copy(s[:], sc)
		return s
	}

	// 驿马计算
	zhi := sike.zhi
	// 根据日支找到三合日支的四孟地支
	// 驿马 = 四孟地支 + 6
	var yiMa ganzhiwuxin.DiZhi
	for i := 0; i < 3; i++ {
		寅, _ := ganzhiwuxin.NewDiZhi("寅")
		if zhi.Minus(寅)%3 == 0 {
			yiMa = zhi.Plus(6)
			break
		}
		// 下一个三合支 = 地支 + 4
		zhi = zhi.Plus(4)
	}
	chu := yiMa
	zhong := sike.zhiYang
	mo := sike.ganYang
	return [3]ganzhiwuxin.DiZhi{chu, zhong, mo}
}

// 获得某课的上神
// 第一课:1,第二课：2，第三课：3，第四课：4
func getShangShen(sk SiKe, n int) ganzhiwuxin.DiZhi {
	if n == 1 {
		return sk.ganYang
	}
	if n == 2 {
		return sk.ganYing
	}
	if n == 3 {
		return sk.zhiYang
	}

	return sk.zhiYing
}
