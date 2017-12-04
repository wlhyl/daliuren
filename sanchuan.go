package shipan

import "ganzhiwuxin"

type SanChuan struct {
	sike           SiKe
	chu, zhong, mo ganzhiwuxin.DiZhi
}

func (sc *SanChuan) Init(sk SiKe) {
	sc.sike = sk
	sc.getSanChuan()
}

func (sc *SanChuan) Get初传() ganzhiwuxin.DiZhi {
	return sc.chu
}

func (sc *SanChuan) Get中传() ganzhiwuxin.DiZhi {
	return sc.zhong
}

func (sc *SanChuan) Get末传() ganzhiwuxin.DiZhi {
	return sc.mo
}

func (sc *SanChuan) getSanChuan() {
	if sc.get伏呤() {
		return
	}
	if sc.get反呤() {
		return
	}
	if sc.get贼克() {
		return
	}
	if sc.get遥克() {
		return
	}

	if sc.get昂星() {
		return
	}

	if sc.get别责() {
		return
	}

	sc.get八专()

}

//有下克上的课
func (sc *SanChuan) has贼() ([]ganzhiwuxin.DiZhi, bool) {
	有克的课 := make([]ganzhiwuxin.DiZhi, 0)
	四课 := sc.sike
	if 四课.gan.Ke(四课.gyang) {
		有克的课 = append(有克的课, 四课.gyang)
	}
	if 四课.gyang.Ke(四课.gying) {
		有克的课 = append(有克的课, 四课.gying)
	}
	if 四课.zhi.Ke(四课.zyang) {
		有克的课 = append(有克的课, 四课.zyang)
	}
	if 四课.zyang.Ke(四课.zying) {
		有克的课 = append(有克的课, 四课.zying)
	}

	if len(有克的课) == 0 {

		return 有克的课, false
	}

	//删除重复课
	listTmp := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range 有克的课 {
		has := false
		for _, v1 := range listTmp {
			if v.Eq(v1) {
				has = true
			}
		}

		if !has {
			listTmp = append(listTmp, v)
		}
	}

	return listTmp, true
}

//有上克下的课
func (sc *SanChuan) has克() ([]ganzhiwuxin.DiZhi, bool) {
	有克的课 := make([]ganzhiwuxin.DiZhi, 0)
	四课 := sc.sike
	if 四课.gyang.Ke(四课.gan) {
		有克的课 = append(有克的课, 四课.gyang)
	}
	if 四课.gying.Ke(四课.gyang) {
		有克的课 = append(有克的课, 四课.gying)
	}
	if 四课.zyang.Ke(四课.zhi) {
		有克的课 = append(有克的课, 四课.zyang)
	}
	if 四课.zying.Ke(四课.zyang) {
		有克的课 = append(有克的课, 四课.zying)
	}

	if len(有克的课) == 0 {

		return 有克的课, false
	}

	//删除重复课
	listTmp := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range 有克的课 {
		has := false
		for _, v1 := range listTmp {
			if v.Eq(v1) {
				has = true
			}
		}

		if !has {
			listTmp = append(listTmp, v)
		}
	}

	return listTmp, true
}

func (sc *SanChuan) get贼克() bool {
	if 贼课, found := sc.has贼(); found {
		if len(贼课) == 1 {
			sc.chu = 贼课[0]
			sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
			sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
			return true
		} else {
			return sc.get比用(贼课, true)
		}
	}
	if 克课, found := sc.has克(); found {
		if len(克课) == 1 {
			sc.chu = 克课[0]
			sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
			sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
			return true
		} else {
			return sc.get比用(克课, false)
		}
	}

	return false
}

func (sc *SanChuan) get比用(d []ganzhiwuxin.DiZhi, is贼 bool) bool {
	/* len(results) ==1, >1, ==0
	   results中保存 取比用後的结果
	*/
	results := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range d {
		if v.Is阳() && sc.sike.gan.Is阳() {
			results = append(results, v)
		}
		if !v.Is阳() && !sc.sike.gan.Is阳() {
			results = append(results, v)
		}
	}

	if len(results) == 1 {
		sc.chu = results[0]
		sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
		sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
		return true
	}
	// 俱不比
	if len(results) == 0 {
		return sc.get涉害(d, is贼)
	}
	//多个俱比 len(result) >1
	return sc.get涉害(results, is贼)

}

func (sc *SanChuan) get涉害(d []ganzhiwuxin.DiZhi, is贼 bool) bool {
	var 寅, 巳, 申, 亥, 卯, 午, 酉, 子 ganzhiwuxin.DiZhi
	寅.Init("寅")
	巳.Init("巳")
	申.Init("申")
	亥.Init("亥")
	卯.Init("卯")
	午.Init("午")
	酉.Init("酉")
	子.Init("子")
	四孟 := [...]ganzhiwuxin.DiZhi{寅, 巳, 申, 亥}
	四仲 := [...]ganzhiwuxin.DiZhi{卯, 午, 酉, 子}

	//使用涉归，不用孟仲季取法
	// ke中的key存储地支的num，value存储涉害深度
	ke := make(map[int]int)
	for _, v := range d {
		临地盘 := sc.sike.GetTianPan().Get临(v)
		count := 0
		for i := 临地盘; !i.Eq(v.Add(1)); i = i.Add(1) {
			天干Slice := Get干From寄宫(i)
			if is贼 {
				if i.Ke(v) {
					count++
				}
				for _, g := range 天干Slice {
					if g.Ke(v) {
						count++
					}
				}

			} else {
				if v.Ke(i) {
					count++
				}
				for _, g := range 天干Slice {
					if v.Ke(g) {
						count++
					}
				}

			}
		}
		ke[v.GetNum()] = count
	}

	最大涉害深度 := 0
	var 有最大涉害深度的支组 []ganzhiwuxin.DiZhi
	for _, v := range ke {
		if v > 最大涉害深度 {
			最大涉害深度 = v
		}
	}
	for k, v := range ke {
		if v == 最大涉害深度 {
			var 有最大涉害深度的支 ganzhiwuxin.DiZhi
			有最大涉害深度的支.Init(k)
			有最大涉害深度的支组 = append(有最大涉害深度的支组, 有最大涉害深度的支)
		}
	}

	if len(有最大涉害深度的支组) == 1 {
		sc.chu = 有最大涉害深度的支组[0]
		sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
		sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
		return true
	}
	//涉害深度相同
	// 从孟发用
	for _, v := range 有最大涉害深度的支组 {
		临地盘 := sc.sike.GetTianPan().Get临(v)
		for _, v1 := range 四孟 {
			if 临地盘.Eq(v1) {
				sc.chu = v
				sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
				sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
				return true
			}
		}
	}
	// 从仲发用
	for _, v := range 有最大涉害深度的支组 {
		临地盘 := sc.sike.GetTianPan().Get临(v)
		for _, v1 := range 四仲 {
			if 临地盘.Eq(v1) {
				sc.chu = v
				sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
				sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
				return true
			}
		}
	}

	if sc.sike.gan.Is阳() {
		sc.chu = sc.sike.gyang
		sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
		sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
		return true
	} else {
		sc.chu = sc.sike.gying
		sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
		sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)
		return true

	}

	// 不能取涉害
	return false
}

func (sc *SanChuan) get遥克() bool {
	// 干阳与支阳相同，课中又无贼克，则取八专，因，此时，仅有两课
	if sc.sike.gyang.Eq(sc.sike.zyang) {
		return sc.get八专()
	}

	// 取遥克干的课
	有遥克的课 := make([]ganzhiwuxin.DiZhi, 0)
	if sc.sike.gying.Ke(sc.sike.gan) {
		有遥克的课 = append(有遥克的课, sc.sike.gying)
	}
	if sc.sike.zyang.Ke(sc.sike.gan) {
		有遥克的课 = append(有遥克的课, sc.sike.zyang)
	}
	if sc.sike.zying.Ke(sc.sike.gan) {
		有遥克的课 = append(有遥克的课, sc.sike.zying)
	}
	//如果无遥克干的课，取干遥克的课
	if len(有遥克的课) == 0 {
		if sc.sike.gan.Ke(sc.sike.gying) {
			有遥克的课 = append(有遥克的课, sc.sike.gying)
		}
		if sc.sike.gan.Ke(sc.sike.zyang) {
			有遥克的课 = append(有遥克的课, sc.sike.zyang)
		}
		if sc.sike.gan.Ke(sc.sike.zying) {
			有遥克的课 = append(有遥克的课, sc.sike.zying)
		}
	}

	if len(有遥克的课) == 0 {
		//无遥克课
		return false
	}

	无重复的克 := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range 有遥克的课 {
		has := false
		for _, v1 := range 无重复的克 {
			if v.Eq(v1) {
				has = true
			}
		}

		if !has {
			无重复的克 = append(无重复的克, v)
		}
	}

	if len(无重复的克) == 1 {

		sc.chu = 无重复的克[0]
		sc.zhong = sc.sike.GetTianPan().GetTianSheng(sc.chu)
		sc.mo = sc.sike.GetTianPan().GetTianSheng(sc.zhong)

		return true
	}
	// 如果有两课遥克，取比用，遥克课取比用之後，不会出现涉害（需要验证）
	return sc.get比用(无重复的克, false)
}

func (sc *SanChuan) get昂星() bool {
	//确认课备
	四课上神组 := make([]ganzhiwuxin.DiZhi, 0)
	四课上神组 = append(四课上神组, sc.sike.gyang)
	四课上神组 = append(四课上神组, sc.sike.gying)
	四课上神组 = append(四课上神组, sc.sike.zyang)
	四课上神组 = append(四课上神组, sc.sike.zying)

	无重复的课 := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range 四课上神组 {
		has := false
		for _, v1 := range 无重复的课 {
			if v.Eq(v1) {
				has = true
			}
		}

		if !has {
			无重复的课 = append(无重复的课, v)
		}
	}

	if len(无重复的课) != 4 {
		return false
	}

	var 酉 ganzhiwuxin.DiZhi
	酉.Init("酉")
	if sc.sike.gan.Is阳() {
		sc.chu = sc.sike.tp.GetTianSheng(酉)
		sc.zhong = sc.sike.zyang
		sc.mo = sc.sike.gyang
		return true
	}

	// 干为阴，取酉所临发用
	sc.chu = sc.sike.tp.Get临(酉)
	sc.zhong = sc.sike.gyang
	sc.mo = sc.sike.zyang
	return true
}

func (sc *SanChuan) get别责() bool {
	//确认课不备
	四课上神组 := make([]ganzhiwuxin.DiZhi, 0)
	四课上神组 = append(四课上神组, sc.sike.gyang)
	四课上神组 = append(四课上神组, sc.sike.gying)
	四课上神组 = append(四课上神组, sc.sike.zyang)
	四课上神组 = append(四课上神组, sc.sike.zying)

	无重复的课 := make([]ganzhiwuxin.DiZhi, 0)
	for _, v := range 四课上神组 {
		has := false
		for _, v1 := range 无重复的课 {
			if v.Eq(v1) {
				has = true
			}
		}

		if !has {
			无重复的课 = append(无重复的课, v)
		}
	}

	//四课全备，不能用别责取三传
	if len(无重复的课) == 4 {
		return false
	}

	//用别责用于三课备取三传
	if len(无重复的课) != 3 {
		return false
	}

	if sc.sike.gan.Is阳() {
		sc.chu = sc.sike.tp.GetTianSheng(Get寄宫(sc.sike.gan.Get五合()))
	} else {
		//干为阴的情况
		sc.chu = sc.sike.zhi.Add(4)
	}
	//中末皆取干上神
	sc.zhong = sc.sike.gyang
	sc.mo = sc.sike.gyang
	return true
}

func (sc *SanChuan) get八专() bool {
	if !sc.sike.gyang.Eq(sc.sike.zyang) {
		return false
	}

	if sc.sike.gan.Is阳() {
		sc.chu = sc.sike.gyang.Add(2)
	} else {
		sc.chu = sc.sike.zying.Add(-2)
	}
	sc.zhong = sc.sike.gyang
	sc.mo = sc.sike.gyang
	return true
}

func (sc *SanChuan) get伏呤() bool {
	if !sc.sike.zhi.Eq(sc.sike.zyang) {
		return false
	}

	var 乙, 癸 ganzhiwuxin.TianGan
	乙.Init(2)
	癸.Init(10)
	//	         六乙、六癸日 无克，阳日取干上神发用
	if sc.sike.gan.Eq(乙) || sc.sike.gan.Eq(癸) || sc.sike.gan.Is阳() {
		sc.chu = sc.sike.gyang
	} else {
		// 阴日，非六乙日、六癸
		sc.chu = sc.sike.zyang
	}

	//求得中传
	sc.zhong = sc.chu.Get刑()
	//	       初为自刑，阳日、六乙日、六癸日取支上神为中传
	if sc.chu.Eq(sc.zhong) {
		if sc.sike.gan.Eq(乙) || sc.sike.gan.Eq(癸) || sc.sike.gan.Is阳() {
			sc.zhong = sc.sike.zyang //六乙、六癸、阳日，初传传自刑，取支上神为中传
		} else {
			sc.zhong = sc.sike.gyang
		}
	}

	//求末传
	sc.mo = sc.zhong.Get刑()
	//中传自刑，取中所冲之神
	if sc.zhong.Eq(sc.mo) {
		sc.mo = sc.mo.Get六冲()
	}

	// 初、中互刑，如：子、卯，末取中所冲之神
	if sc.zhong.Is刑(sc.chu) {
		sc.mo = sc.zhong.Get六冲()
	}
	return true
}

func (sc *SanChuan) get反呤() bool {
	if !sc.sike.zhi.Is六冲(sc.sike.zyang) {
		return false
	}
	if sc.get贼克() {
		return true
	}

	for _, v := range [...]int{1, 4, 7, 10} {
		if sc.sike.zhi.GetNum() == v {
			sc.chu = sc.sike.zhi.Add(6)
			sc.zhong = sc.sike.zyang
			sc.mo = sc.sike.gyang
			return true
		}
	}
	for _, v := range [...]int{2, 5, 8, 11} {
		if sc.sike.zhi.GetNum() == v {
			sc.chu = sc.sike.zhi.Add(-4).Add(6)
			sc.zhong = sc.sike.zyang
			sc.mo = sc.sike.gyang
			return true
		}
	}

	// 支为四季
	sc.chu = sc.sike.zhi.Add(4).Add(6)
	sc.zhong = sc.sike.zyang
	sc.mo = sc.sike.gyang
	return true
}
