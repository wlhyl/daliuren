package shipan

import "ganzhiwuxin"

type SiKe struct {
	/*
		tp 天盘
		gan 日干
		zhi 日支
		gyang 干阳
		gying 干阴
		zyang 支阳
		zying 支阴
	*/
	tp                              TianPan
	gan                             ganzhiwuxin.TianGan
	zhi, gyang, gying, zyang, zying ganzhiwuxin.DiZhi
}

func (sk *SiKe) Init(tp TianPan, jz ganzhiwuxin.JiaZi) {
	sk.tp = tp

	sk.gan = jz.GetGan()
	sk.zhi = jz.GetDiZhi()

	sk.gyang = tp.GetTianSheng(Get寄宫(sk.gan))
	sk.gying = tp.GetTianSheng(sk.gyang)

	sk.zyang = tp.GetTianSheng(sk.zhi)
	sk.zying = tp.GetTianSheng(sk.zyang)
}

//得到天盘
func (sk SiKe) GetTianPan() TianPan {
	return sk.tp
}

//得到日干
func (sk SiKe) GetGan() ganzhiwuxin.TianGan {
	return sk.gan
}

//得到日支
func (sk SiKe) GetZhi() ganzhiwuxin.DiZhi {
	return sk.zhi
}

//得到干阳
func (sk SiKe) GetGanYang() ganzhiwuxin.DiZhi {
	return sk.gyang
}

//得到干阴
func (sk SiKe) GetGanYing() ganzhiwuxin.DiZhi {
	return sk.gying
}

//得到支阳
func (sk SiKe) GetZhiYang() ganzhiwuxin.DiZhi {
	return sk.zyang
}

//得到支阴
func (sk SiKe) GetZhiYing() ganzhiwuxin.DiZhi {
	return sk.zying
}

//判断两个四课格局是否一样
func (sk SiKe) Eq(sk2 SiKe) bool {
	return sk.tp.Eq(sk2.tp) && sk.gan.Eq(sk2.gan) && sk.zhi.Eq(sk2.zhi)
}

func (sc SiKe) Get遁干(d ganzhiwuxin.DiZhi) (t ganzhiwuxin.TianGan, hasg bool) {
	var 甲, 遁干 ganzhiwuxin.TianGan
	甲.Init("甲")
	旬首 := sc.zhi.Add(甲.Sub(sc.gan))
	delta := (d.Sub(旬首) + 12) % 12
	if delta >= 10 {
		return
	}
	遁干 = 甲.Add(delta)
	return 遁干, true
}
