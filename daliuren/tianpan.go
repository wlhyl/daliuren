package daliuren

import (
	"encoding/json"

	"github.com/wlhyl/ganzhiwuxin"
)

type TianPan struct {
	yueJiang       ganzhiwuxin.DiZhi
	divinationTime ganzhiwuxin.DiZhi
}

//取得地盘上神
func (tp TianPan) up(d ganzhiwuxin.DiZhi) ganzhiwuxin.DiZhi {
	return tp.yueJiang.Plus(d.Minus(tp.divinationTime))
}

//取得天神所临地盘
func (tp TianPan) down(d ganzhiwuxin.DiZhi) ganzhiwuxin.DiZhi {
	return tp.divinationTime.Plus(d.Minus(tp.yueJiang))
}

func (tp TianPan) MarshalJSON() ([]byte, error) {
	var t [12]string
	zi, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		return nil, err
	}
	for i := 0; i < 12; i++ {
		d := zi.Plus(i)
		t[i] = tp.up(d).Name()
	}
	return json.Marshal(t)
}
