package daliuren

import (
	"encoding/json"
	"fmt"

	"github.com/wlhyl/ganzhiwuxin"
)

var tianJiangNumToName = [...]string{"贵", "蛇", "雀", "合", "勾", "龙", "空", "虎", "常", "玄", "阴", "后"}

// 天将
type TianJiang struct {
	num  int
	name string
	// good bool
}

func NewTianJiang(name string) (TianJiang, error) {
	var tj TianJiang
	for i, v := range tianJiangNumToName {
		if v == name {
			tj.num = i + 1
			tj.name = v
			// switch tj.num {
			// case 1, 4, 6, 9, 11, 12:
			// 	tj.good = true
			// default:
			// 	tj.good = false
			// }
			return tj, nil
		}
	}
	return tj, fmt.Errorf("%s不是正确的天将", name)

}
func (tj TianJiang) Name() string {
	return tj.name
}
func (tj TianJiang) Equals(other TianJiang) bool {
	return tj.num == other.num
}

func (tj TianJiang) Plus(other int) TianJiang {
	var tmp = other
	for tmp < 0 {
		tmp += 12
	}
	tmp = (tj.num + tmp) % 12
	if tmp == 0 {
		tmp = 12
	}
	t, _ := NewTianJiang(tianJiangNumToName[tmp-1])
	return t
}

func (tj TianJiang) Minus(other TianJiang) int {
	// 返回整数值
	return (tj.num - other.num + 12) % 12
}

type TianJiangPan struct {
	tianYiDiZhi ganzhiwuxin.DiZhi // 天乙所在的地支
	inverse     bool              // 逆布为true
}

// tianPan:天盘
// gan:日天干
// diurnal:昼占为true，夜占为false
func NewTianJiangPan(tianPan TianPan, gan ganzhiwuxin.TianGan, diurnal bool) (TianJiangPan, error) {
	var tjpan TianJiangPan
	// 天乙所在的地支
	if tianYiDiZhi, err := tianYiDiZhi(gan, diurnal); err != nil {
		return tjpan, err
	} else {
		tjpan.tianYiDiZhi = tianYiDiZhi
	}
	if inverse, err := isInverse(tianPan, tjpan.tianYiDiZhi); err != nil {
		return tjpan, err
	} else {
		tjpan.inverse = inverse
	}
	return tjpan, nil
}

func tianYiDiZhi(gan ganzhiwuxin.TianGan, diurnal bool) (ganzhiwuxin.DiZhi, error) {
	甲, err := ganzhiwuxin.NewTianGan("甲")
	if err != nil {
		return ganzhiwuxin.DiZhi{}, nil
	}
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		return ganzhiwuxin.DiZhi{}, nil
	}
	丑 := 子.Plus(1)
	寅 := 丑.Plus(1)
	卯 := 寅.Plus(1)
	辰 := 卯.Plus(1)
	巳 := 辰.Plus(1)
	午 := 巳.Plus(1)
	未 := 午.Plus(1)
	申 := 未.Plus(1)
	酉 := 申.Plus(1)
	戌 := 酉.Plus(1)
	亥 := 戌.Plus(1)
	昼贵 := [...]ganzhiwuxin.DiZhi{未, 申, 酉, 亥, 丑, 子, 丑, 寅, 卯, 巳}
	夜贵 := [...]ganzhiwuxin.DiZhi{丑, 子, 亥, 酉, 未, 申, 未, 午, 巳, 卯}
	n := gan.Minus(甲)
	if diurnal {
		return 昼贵[n], nil
	} else {
		return 夜贵[n], nil
	}
}

func isInverse(tianPan TianPan, tianYiDiZhi ganzhiwuxin.DiZhi) (bool, error) {
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		return false, err
	}
	// 贵人地盘之支
	guiRenDiPan := 子.Plus(tianYiDiZhi.Minus(tianPan.up(子)))

	巳 := 子.Plus(5)
	// 支 - 巳 必定 >=0，支 - 巳 <=5，则 巳=< 支 <= 戌
	if guiRenDiPan.Minus(巳) <= 5 {
		return true, nil
	}
	return false, nil
}

// 获取某地支的天将
func (tjPan TianJiangPan) up(diZhi ganzhiwuxin.DiZhi) TianJiang {
	贵, _ := NewTianJiang("贵")
	n := diZhi.Minus(tjPan.tianYiDiZhi)
	if tjPan.inverse {
		return 贵.Plus(-n)
	} else {
		return 贵.Plus(n)
	}
}

// 获取某天将所临地支
func (tjPan TianJiangPan) down(tianJiang TianJiang) ganzhiwuxin.DiZhi {
	贵, _ := NewTianJiang("贵")
	n := tianJiang.Minus(贵)
	if tjPan.inverse {
		return tjPan.tianYiDiZhi.Plus(-n)
	} else {
		return tjPan.tianYiDiZhi.Plus(n)
	}
}

func (tjPan TianJiangPan) MarshalJSON() ([]byte, error) {
	var t [12]string
	zi, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		return nil, err
	}
	for i := 0; i < 12; i++ {
		d := zi.Plus(i)
		t[i] = tjPan.up(d).Name()
	}
	return json.Marshal(t)
}
