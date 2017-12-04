package shipan

import "ganzhiwuxin"

func Get寄宫(g ganzhiwuxin.TianGan) ganzhiwuxin.DiZhi {
	to寄宫映射 := [...]int{0, 1, 3, 4, 6, 4, 6, 7, 9, 10, 12}
	var d ganzhiwuxin.DiZhi
	d.Init(to寄宫映射[g.GetNum()])
	return d
}

func Get干From寄宫(d ganzhiwuxin.DiZhi) []ganzhiwuxin.TianGan {
	gSlice := make([]ganzhiwuxin.TianGan, 0)
	for i := 1; i < 11; i++ {
		var g ganzhiwuxin.TianGan
		var d1 ganzhiwuxin.DiZhi
		g.Init(i)
		d1 = Get寄宫(g)
		if d1.Eq(d) {
			gSlice = append(gSlice, g)
		}
	}
	return gSlice
}

type TianPan struct {
	yueJiang ganzhiwuxin.DiZhi
	zhanShi  ganzhiwuxin.DiZhi
}

func (tp *TianPan) Init(y, z ganzhiwuxin.DiZhi) {
	tp.yueJiang = y
	tp.zhanShi = z
}

//取得地盘上神
func (tp TianPan) GetTianSheng(d ganzhiwuxin.DiZhi) ganzhiwuxin.DiZhi {
	return tp.yueJiang.Add(d.Sub(tp.zhanShi))
}

//取得天神所临地盘
func (tp TianPan) Get临(d ganzhiwuxin.DiZhi) ganzhiwuxin.DiZhi {
	return tp.zhanShi.Add(d.Sub(tp.yueJiang))
}

//取得月将
func (tp TianPan) GetYueJiang() ganzhiwuxin.DiZhi {
	return tp.yueJiang
}

//取得占时
func (tp TianPan) GetZhanShi() ganzhiwuxin.DiZhi {
	return tp.zhanShi
}

//判断两个天盘格局是否一样
func (tp TianPan) Eq(tp2 TianPan) bool {
	var 子 ganzhiwuxin.DiZhi
	子.Init(1)
	子上神 := tp.yueJiang.Add((tp.zhanShi.Sub(子)))
	子上神2 := tp2.yueJiang.Add((tp2.zhanShi.Sub(子)))
	return 子上神.Eq(子上神2)
}
