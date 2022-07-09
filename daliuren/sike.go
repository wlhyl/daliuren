package daliuren

import (
	"encoding/json"

	"github.com/wlhyl/ganzhiwuxin"
)

type SiKe struct {
	/*
		gan 日干
		zhi 日支
		ganYang 干阳
		ganYing 干阴
		zhiYang 支阳
		zhiYing 支阴
	*/
	gan                                     ganzhiwuxin.TianGan
	zhi, ganYang, ganYing, zhiYang, zhiYing ganzhiwuxin.DiZhi
}

func NewSiKe(tianPan TianPan, divinationDay ganzhiwuxin.GanZhi) (sike SiKe) {

	sike.gan = divinationDay.Gan()
	sike.ganYang = tianPan.up(jiGong(sike.gan))
	sike.ganYing = tianPan.up(sike.ganYang)

	sike.zhi = divinationDay.Zhi()
	sike.zhiYang = tianPan.up(sike.zhi)
	sike.zhiYing = tianPan.up(sike.zhiYang)
	return
}

func (sike SiKe) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"gan":     sike.gan.Name(),
		"ganYang": sike.ganYang.Name(),
		"ganYing": sike.ganYing.Name(),
		"zhi":     sike.zhi.Name(),
		"zhiYang": sike.zhiYang.Name(),
		"zhiYing": sike.zhiYing.Name(),
	}
	return json.Marshal(m)
}
