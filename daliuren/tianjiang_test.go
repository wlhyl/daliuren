package daliuren

import (
	"encoding/json"
	"testing"

	"github.com/wlhyl/ganzhiwuxin"
)

var tianJiangName = [...]string{"贵", "蛇", "雀", "合", "勾", "龙", "空", "虎", "常", "玄", "阴", "后"}

func TestNewTianJiang(t *testing.T) {
	t.Log("NewTianJiang()方法")
	t.Log("正确创建TianJian")
	for _, v := range tianJiangName {
		w, err := NewTianJiang(v)
		if err != nil {
			t.Error(err)
		}
		if w.Name() != v {
			t.Error("Name()方法测试失败")
		}
	}

	t.Log("不正确天将创建")
	if _, err := NewTianJiang("否"); err == nil {
		t.Error("`否`不是天将")
	}
}

func TestTianJiangEquals(t *testing.T) {
	t.Log("天将相等")
	for _, name := range tianJiangName {
		tg0, _ := NewTianJiang(name)
		tg1, _ := NewTianJiang(name)
		if !tg0.Equals(tg1) {
			t.Errorf("%s 与 %s 相等", tg0.Name(), tg1.Name())
		}
	}
}

func TesttianJiangNotEqual(t *testing.T) {
	t.Log("天将不相等")
	for i, v := range tianJiangName {
		g0, _ := NewTianJiang(v)
		for j := 1; j < len(tianJiangName); j++ {
			n := (i + j) % len(tianJiangName)
			g1, _ := NewTianJiang(tianJiangName[n])
			if g0.Equals(g1) {
				t.Errorf("%s != %s", g0.Name(), g1.Name())
			}
		}
	}
}

//用数学归纳法
func TestTianJiangPlus(t *testing.T) {
	t.Log("天将 + 整数")
	for i, v := range tianJiangName {
		tg0, _ := NewTianJiang(tianJiangName[(i+1)%len(tianJiangName)])
		tg1, _ := NewTianJiang(v)
		tg1 = tg1.Plus(1)
		if !tg0.Equals(tg1) {
			t.Errorf("%s + 1 = %s", tg1.Name(), tg0.Name())
		}
	}

	for i, v := range tianJiangName {
		tg0, _ := NewTianJiang(tianJiangName[(i-1+len(tianJiangName))%len(tianJiangName)])
		tg1, _ := NewTianJiang(v)
		tg1 = tg1.Plus(-1)
		if !tg0.Equals(tg1) {
			t.Errorf("%s + (-1) = %s", tg1.Name(), tg0.Name())
		}
	}

	for _, it := range tianJiangName {
		g, _ := NewTianJiang(it)
		if !(g.Plus(99)).Plus(1).Equals(g.Plus(100)) {
			t.Errorf("%s + 99 + 1 = %s + 100", g.Name(), g.Name())
		}
		if !g.Plus(-99).Plus(-1).Equals(g.Plus(-100)) {
			t.Errorf("%s + (-99) + (-1) = %s + (-100)", g.Name(), g.Name())
		}
	}
}

func TestTianJiangMinus(t *testing.T) {
	t.Log("天将 - 地支 = 正整数")
	for _, v := range tianJiangName {
		g0, _ := NewTianJiang(v)
		for j := 0; j < len(tianJiangName); j++ {
			g1 := g0.Plus(j)
			if j != g1.Minus(g0) {
				t.Errorf("%s - %s = %v", g1.Name(), g0.Name(), j)
			}
		}
	}
}

func TestNewTianJiangPan(t *testing.T) {
	t.Log("测试NewTianJiangPan")
	申, err := ganzhiwuxin.NewDiZhi("申")
	if err != nil {
		t.Fatalf(err.Error())
	}
	辰, err := ganzhiwuxin.NewDiZhi("辰")
	if err != nil {
		t.Fatalf(err.Error())
	}
	甲, err := ganzhiwuxin.NewTianGan("甲")
	if err != nil {
		t.Fatalf(err.Error())
	}
	tp := TianPan{申, 辰}
	tjpan, err := NewTianJiangPan(tp, 甲, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// 甲日昼贵在未，申将辰时
	if !tjpan.tianYiDiZhi.Equals(申.Plus(-1)) {
		t.Errorf("甲日昼贵人在未，非是%s", tjpan.tianYiDiZhi.Name())
	}

	if tjpan.inverse {
		t.Error("甲日昼占，申将辰时，天将顺布")
	}

	tjpan, err = NewTianJiangPan(tp, 甲, false)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// 甲日夜贵在未，申将辰时
	if !tjpan.tianYiDiZhi.Equals(申.Plus(5)) {
		t.Errorf("甲日夜贵人在丑，非是%s", tjpan.tianYiDiZhi.Name())
	}

	if !tjpan.inverse {
		t.Error("甲日夜占，申将辰时，天将逆布")
	}

	// 顺逆临界条件测试
	t.Log("测试甲日贵人在辰")
	未 := 申.Plus(-1)
	tp = TianPan{未, 辰}
	tjpan, err = NewTianJiangPan(tp, 甲, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if tjpan.inverse {
		t.Error("甲日昼占，未将辰时，天将顺布")
	}
	tjpan, err = NewTianJiangPan(tp, 甲, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !tjpan.inverse {
		t.Error("甲日夜占，未将辰时，天将逆布")
	}

	t.Log("测试甲日贵人在巳")
	巳 := 申.Plus(-3)
	tp = TianPan{未, 巳}
	tjpan, err = NewTianJiangPan(tp, 甲, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !tjpan.inverse {
		t.Error("甲日昼占，未将巳时，天将逆布")
	}
	tjpan, err = NewTianJiangPan(tp, 甲, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if tjpan.inverse {
		t.Error("甲日夜占，未将巳时，天将顺布")
	}

	t.Log("测试甲日贵人在戌")
	戌 := 申.Plus(2)
	tp = TianPan{未, 戌}
	tjpan, err = NewTianJiangPan(tp, 甲, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !tjpan.inverse {
		t.Error("甲日昼占，未将戌时，天将逆布")
	}
	tjpan, err = NewTianJiangPan(tp, 甲, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if tjpan.inverse {
		t.Error("甲日夜占，未将戌时，天将顺布")
	}

	t.Log("测试甲日贵人在亥")
	亥 := 申.Plus(3)
	tp = TianPan{未, 亥}
	tjpan, err = NewTianJiangPan(tp, 甲, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if tjpan.inverse {
		t.Error("甲日昼占，未将亥时，天将顺布")
	}
	tjpan, err = NewTianJiangPan(tp, 甲, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !tjpan.inverse {
		t.Error("甲日夜占，未将亥时，天将逆布")
	}
}

func TestTianJiangPanUpAndDown(t *testing.T) {
	t.Log("测试获取地支所乘天将和天将所临地支")
	子, _ := ganzhiwuxin.NewDiZhi("子")
	甲, _ := ganzhiwuxin.NewTianGan("甲")
	// 甲日昼占，子将丑时
	tp := TianPan{子, 子.Plus(1)}
	tjPan, err := NewTianJiangPan(tp, 甲, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	// 子上所乘天将是 白虎，天将逆布
	虎, err := NewTianJiang("虎")
	if err != nil {
		t.Fatalf(err.Error())
	}
	for i := 0; i < 12; i++ {
		d := 子.Plus(i)
		tj := 虎.Plus(-i)
		if tj0 := tjPan.up(d); !tj0.Equals(tj) {
			t.Fatalf("`%s`所乘天将是`%s`，非是`%s`", d.Name(), tj.name, tj0.name)
		}

		if d0 := tjPan.down(tj); !d0.Equals(d) {
			t.Fatalf("`%s`所临地支是`%s`，非是`%s`", tj.name, d.Name(), d0.Name())
		}
	}
}

func TestTianJianPanJson(t *testing.T) {
	t.Log("测试TianJiangPan序列化")
	子, _ := ganzhiwuxin.NewDiZhi("子")
	丑 := 子.Plus(1)
	甲, _ := ganzhiwuxin.NewTianGan("甲")
	tp := TianPan{子, 丑}
	tjPan, _ := NewTianJiangPan(tp, 甲, true)
	json, err := json.Marshal(tjPan)
	if err != nil {
		t.Fatalf(err.Error())
	}
	s := "[\"虎\",\"空\",\"龙\",\"勾\",\"合\",\"雀\",\"蛇\",\"贵\",\"后\",\"阴\",\"玄\",\"常\"]"
	if string(json) != s {
		t.Fatalf("天将盘序列化失败")
	}
}
