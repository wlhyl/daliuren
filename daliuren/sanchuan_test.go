package daliuren

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlhyl/ganzhiwuxin"
)

func Test伏吟(t *testing.T) {
	t.Log("测试伏吟")
	甲, _ := ganzhiwuxin.NewTianGan("甲")
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		t.Fatal(err)
	}
	tp := TianPan{子, 子}
	t.Log("六甲日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2 * i)
		gz, _ := ganzhiwuxin.NewGanZhi(甲, z)
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)
		寅 := 子.Plus(2)
		巳 := 寅.Plus(3)
		申 := 巳.Plus(3)
		if !sc.chu.Equals(寅) || !sc.zhong.Equals(巳) || !sc.mo.Equals(申) {
			t.Fatalf("六甲日伏吟三传`寅，巳，申`，非是`%s`，`%s`，`%s`", sc.chu.Name(), sc.zhong.Name(), sc.mo.Name())
		}
	}
	t.Log("六丙日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2 * i)
		gz, err := ganzhiwuxin.NewGanZhi(甲.Plus(2), z)
		if err != nil {
			t.Fatal(err)
		}
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)
		chu, zhong, mo := sc.chu.Name(), sc.zhong.Name(), sc.mo.Name()
		if chu != "巳" || zhong != "申" || mo != "寅" {
			t.Fatalf("六丙日伏吟三传`巳，申，寅`，非是`%s`，`%s`，`%s`", chu, zhong, mo)
		}
	}
	t.Log("六戊日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2 * i)
		gz, err := ganzhiwuxin.NewGanZhi(甲.Plus(4), z)
		if err != nil {
			t.Fatal(err)
		}
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)
		chu, zhong, mo := sc.chu.Name(), sc.zhong.Name(), sc.mo.Name()
		if chu != "巳" || zhong != "申" || mo != "寅" {
			t.Fatalf("六戊日伏吟三传`巳，申，寅`，非是`%s`，`%s`，`%s`", chu, zhong, mo)
		}
	}
	t.Log("六庚日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2 * i)
		gz, err := ganzhiwuxin.NewGanZhi(甲.Plus(6), z)
		if err != nil {
			t.Fatal(err)
		}
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)
		chu, zhong, mo := sc.chu.Name(), sc.zhong.Name(), sc.mo.Name()
		if chu != "申" || zhong != "寅" || mo != "巳" {
			t.Fatalf("六庚日伏吟三传`申，寅，巳`，非是`%s`，`%s`，`%s`", chu, zhong, mo)
		}
	}

	t.Log("六壬日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2 * i)
		gz, err := ganzhiwuxin.NewGanZhi(甲.Plus(8), z)
		if err != nil {
			t.Fatal(err)
		}
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)
		亥 := 子.Plus(-1)

		if !sc.chu.Equals(亥) || !sc.zhong.Equals(z) {
			t.Fatalf("六壬日伏吟初传`亥`，中传`%s`, 非是`%s`, `%s`", z.Name(), sc.chu.Name(), sc.zhong.Name())
		}
		if 卯 := 子.Plus(3); z.Equals(子) && !sc.mo.Equals(卯) {
			t.Fatalf("壬子日伏呤，末传`卯`，非是`%s`", 卯.Name())
		}
		if 巳 := 子.Plus(5); z.Equals(子.Plus(2)) && !sc.mo.Equals(巳) {
			t.Fatalf("壬寅日伏呤，末传`巳`，非是`%s`", 巳.Name())
		}
		if 戌 := 子.Plus(-2); z.Equals(子.Plus(4)) && !sc.mo.Equals(戌) {
			t.Fatalf("壬辰日伏呤，末传`戌`，非是`%s`", 戌.Name())
		}
		if z.Equals(子.Plus(6)) && !sc.mo.Equals(子) {
			t.Fatalf("壬午日伏呤，末传`子`，非是`%s`", 子.Name())
		}
		if 寅 := 子.Plus(2); z.Equals(子.Plus(8)) && !sc.mo.Equals(寅) {
			t.Fatalf("壬申日伏呤，末传`寅`，非是`%s`", 寅.Name())
		}
		if 未 := 子.Plus(7); z.Equals(子.Plus(10)) && !sc.mo.Equals(未) {
			t.Fatalf("壬戌日伏呤，末传`未`，非是`%s`", 未.Name())
		}
	}

	t.Log("六乙日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2*i + 1)
		gz, err := ganzhiwuxin.NewGanZhi(甲.Plus(1), z)
		if err != nil {
			t.Fatal(err)
		}
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)
		辰 := 子.Plus(4)

		if !sc.chu.Equals(辰) || !sc.zhong.Equals(z) {
			t.Fatalf("乙%s日伏吟初传`辰`，中传`%s`, 非是`%s`, `%s`", z.Name(), z.Name(), sc.chu.Name(), sc.zhong.Name())
		}
		if zhi := 子.Plus(-2); z.Equals(子.Plus(1)) && !sc.mo.Equals(zhi) {
			t.Fatalf("乙丑日伏呤，末传`戌`，非是`%s`", zhi.Name())
		}
		if z.Equals(子.Plus(3)) && !sc.mo.Equals(子) {
			t.Fatalf("乙卯日伏呤，末传`子`，非是`%s`", 子.Name())
		}
		if zhi := 子.Plus(8); z.Equals(子.Plus(5)) && !sc.mo.Equals(zhi) {
			t.Fatalf("乙巳日伏呤，末传`申`，非是`%s`", zhi.Name())
		}
		if zhi := 子.Plus(1); z.Equals(子.Plus(7)) && !sc.mo.Equals(zhi) {
			t.Fatalf("乙未日伏呤，末传`丑`，非是`%s`", zhi.Name())
		}
		if zhi := 子.Plus(3); z.Equals(子.Plus(9)) && !sc.mo.Equals(zhi) {
			t.Fatalf("乙酉日伏呤，末传`卯`，非是`%s`", zhi.Name())
		}
		if zhi := 子.Plus(5); z.Equals(子.Plus(11)) && !sc.mo.Equals(zhi) {
			t.Fatalf("乙亥日伏呤，末传`巳`，非是`%s`", zhi.Name())
		}
	}

	t.Log("六丁、己、辛日")
	//  丁=甲 + 3
	// 己 = 甲 + 5
	// 辛 = 甲 + 7
	for _, g := range [...]ganzhiwuxin.TianGan{甲.Plus(3), 甲.Plus(5), 甲.Plus(7)} {
		for i := 0; i < 6; i++ {
			z := 子.Plus(2*i + 1)
			gz, err := ganzhiwuxin.NewGanZhi(g, z)
			if err != nil {
				t.Fatal(err)
			}
			sike := NewSiKe(tp, gz)
			sc := NewSanChuan(tp, sike)

			// 丑
			if z.Equals(子.Plus(1)) && (!sc.chu.Equals(子.Plus(1)) || !sc.zhong.Equals(子.Plus(-2)) || !sc.mo.Equals(子.Plus(7))) {
				t.Fatalf("%s日伏吟三传`丑，戌，未`，非是`%s，%s，%s`", gz.Name(), sc.chu.Name(), sc.zhong.Name(), sc.mo.Name())
			}

			// 卯
			if z.Equals(子.Plus(3)) && (!sc.chu.Equals(子.Plus(3)) || !sc.zhong.Equals(子) || !sc.mo.Equals(子.Plus(6))) {
				t.Fatalf("%s日伏吟三传`卯，子，午`，非是`%s，%s，%s`", gz.Name(), sc.chu.Name(), sc.zhong.Name(), sc.mo.Name())
			}

			// 巳
			if z.Equals(子.Plus(5)) && (!sc.chu.Equals(子.Plus(5)) || !sc.zhong.Equals(子.Plus(8)) || !sc.mo.Equals(子.Plus(2))) {
				t.Fatalf("%s日伏吟三传`巳，申，寅`，非是`%s，%s，%s`", gz.Name(), sc.chu.Name(), sc.zhong.Name(), sc.mo.Name())
			}
			// 未
			if z.Equals(子.Plus(7)) && (!sc.chu.Equals(子.Plus(7)) || !sc.zhong.Equals(子.Plus(1)) || !sc.mo.Equals(子.Plus(-2))) {
				t.Fatalf("%s日伏吟三传`未，丑，戌`，非是`%s，%s，%s`", gz.Name(), sc.chu.Name(), sc.zhong.Name(), sc.mo.Name())
			}

			// 酉、亥日
			if z.Equals(子.Plus(9)) || z.Equals(子.Plus(11)) {
				if !sc.chu.Equals(z) && !sc.zhong.Equals(sike.ganYang) {
					t.Fatalf("%s日伏吟，初传为`%s`，中传为`%s`, 非是`%s`，`%s`", gz.Name(), z.Name(), sike.ganYang.Name(), sc.chu.Name(), sc.zhong.Name())
				}

				// 丁、己日
				if (g.Equals(甲.Plus(3)) || g.Equals(甲.Plus(5))) && !sc.mo.Equals(子.Plus(1)) {
					t.Fatalf("%s日伏吟，末传为`丑`，非是`%s`", gz.Name(), sc.mo.Name())
				}
				// 辛日
				if g.Equals(甲.Plus(7)) && !sc.mo.Equals(子.Plus(7)) {
					t.Fatalf("%s日伏吟，末传为`未`，非是`%s`", gz.Name(), sc.mo.Name())
				}

			}
		}
	}
	// t.Log("六癸日")
	for i := 0; i < 6; i++ {
		z := 子.Plus(2*i + 1)
		gz, err := ganzhiwuxin.NewGanZhi(甲.Plus(-1), z)
		if err != nil {
			t.Fatal(err)
		}
		sike := NewSiKe(tp, gz)
		sc := NewSanChuan(tp, sike)

		chu, zhong, mo := sc.chu.Name(), sc.zhong.Name(), sc.mo.Name()
		if chu != "丑" || zhong != "戌" || mo != "未" {
			t.Fatalf("六癸日伏吟，三传`丑、戌、未`，非是`%s, %s, %s`", chu, zhong, mo)
		}
	}
}

func TestHas贼(t *testing.T) {
	t.Log("测试函数has贼()")

	t.Log("没有下克上的课")
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		t.Fatal(err)
	}
	甲, err := ganzhiwuxin.NewTianGan("甲")
	if err != nil {
		t.Fatal(err)
	}
	甲子, err := ganzhiwuxin.NewGanZhi(甲, 子)
	if err != nil {
		t.Fatal(err)
	}
	tp := TianPan{子, 子}
	sk := NewSiKe(tp, 甲子)
	ze := has贼(sk)
	if len(ze) != 0 {
		t.Fatal("甲子日伏吟，四课中无下克上")
	}

	t.Log("有下克上的课")
	// 课中有重复
	// 辰将寅时，甲子日，四课：甲、辰、午、子、寅、辰
	// 返回值:[1]
	tp = TianPan{子.Plus(4), 子.Plus(2)}
	sk = NewSiKe(tp, 甲子)
	ze = has贼(sk)
	if len(ze) != 1 || ze[0] != 1 {
		t.Fatal("甲子日辰将寅时，取第一课，下克上")
	}
}

func TestHas克(t *testing.T) {
	t.Log("测试函数has克()")

	t.Log("没有上克下的课")
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		t.Fatal(err)
	}
	甲, err := ganzhiwuxin.NewTianGan("甲")
	if err != nil {
		t.Fatal(err)
	}
	甲子, err := ganzhiwuxin.NewGanZhi(甲, 子)
	if err != nil {
		t.Fatal(err)
	}
	tp := TianPan{子, 子}
	sk := NewSiKe(tp, 甲子)
	ze := has克(sk)
	if len(ze) != 0 {
		t.Fatal("甲子日伏吟，四课中无上克下")
	}

	t.Log("有上克下的课")
	// 课中无重复
	// 甲子日，酉将寅时，四课：甲、酉、辰、子、未、寅
	// 返回值:[1, 3 , 4]
	酉 := 子.Plus(-3)
	寅 := 子.Plus(2)
	tp = TianPan{酉, 寅}
	sk = NewSiKe(tp, 甲子)
	ze = has克(sk)
	if len(ze) != 3 || ze[0] != 1 || ze[1] != 3 || ze[2] != 4 {
		t.Fatalf("甲子日酉将寅时，第一、三、四课有下克上，非是`%v`", ze)
	}

	// 课中有重复
	// 甲寅日，酉将寅时，四课：甲、酉、辰、寅、酉、辰
	// 返回值:[1, 3 , 4]
	甲寅, err := ganzhiwuxin.NewGanZhi(甲, 寅)
	if err != nil {
		t.Fatal(err)
	}
	tp = TianPan{酉, 寅}
	sk = NewSiKe(tp, 甲寅)
	ze = has克(sk)

	if len(ze) != 1 || ze[0] != 1 {
		skJson, err := json.Marshal(sk)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(skJson))
		t.Fatalf("甲寅日酉将寅时，取第一课的下克上，非是`%v`", ze)
	}
}

func TestGet涉害(t *testing.T) {
	t.Log("涉害谭取三传")
	甲, err := ganzhiwuxin.NewTianGan("甲")
	if err != nil {
		t.Fatal(err)
	}
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("丁卯日，亥将丑时")
	丁 := 甲.Plus(3)
	卯 := 子.Plus(3)
	亥 := 子.Plus(-1)
	丑 := 子.Plus(1)
	gz, err := ganzhiwuxin.NewGanZhi(丁, 卯)
	if err != nil {
		t.Fatal(err)
	}
	tp := TianPan{亥, 丑}
	sk := NewSiKe(tp, gz)
	//三、四课涉害
	keList := []int{3, 4}
	sc := get涉害(tp, sk, keList)
	chu := sc[0].Name()
	zhong := sc[1].Name()
	mo := sc[2].Name()
	if chu != "亥" || zhong != "酉" || mo != "未" {
		t.Fatalf("丁卯日，亥将丑时，三传：亥、酉、未，非：`%s、%s、%s`", chu, zhong, mo)
	}

	t.Log("庚子日，申将戌时")
	庚 := 甲.Plus(6)
	申 := 子.Plus(8)
	戌 := 子.Plus(-2)
	gz, err = ganzhiwuxin.NewGanZhi(庚, 子)
	if err != nil {
		t.Fatal(err)
	}
	tp = TianPan{申, 戌}
	sk = NewSiKe(tp, gz)
	//一、三课涉害
	keList = []int{1, 3}
	sc = get涉害(tp, sk, keList)
	chu = sc[0].Name()
	zhong = sc[1].Name()
	mo = sc[2].Name()
	if chu != "午" || zhong != "辰" || mo != "寅" {
		t.Fatalf("%s日，申将戌时，三传：午、辰、寅，非：`%s、%s、%s`", gz.Name(), chu, zhong, mo)
	}

	t.Log("丙子日，亥将辰时，取孟")
	丙 := 甲.Plus(2)
	亥 = 子.Plus(-1)
	辰 := 子.Plus(4)
	gz, err = ganzhiwuxin.NewGanZhi(丙, 子)
	if err != nil {
		t.Fatal(err)
	}
	tp = TianPan{亥, 辰}
	sk = NewSiKe(tp, gz)
	//一、四课涉害
	keList = []int{1, 4}
	sc = get涉害(tp, sk, keList)
	chu = sc[0].Name()
	zhong = sc[1].Name()
	mo = sc[2].Name()
	if chu != "子" || zhong != "未" || mo != "寅" {
		t.Fatalf("%s日，亥将辰时，三传：子、未、寅，非：`%s、%s、%s`", gz.Name(), chu, zhong, mo)
	}

	t.Log("庚午日，未将卯时，取仲")
	庚 = 甲.Plus(6)
	午 := 子.Plus(6)
	未 := 子.Plus(7)
	卯 = 子.Plus(3)
	gz, err = ganzhiwuxin.NewGanZhi(庚, 午)
	if err != nil {
		t.Fatal(err)
	}
	tp = TianPan{未, 卯}
	sk = NewSiKe(tp, gz)
	//二、四课涉害
	keList = []int{2, 4}
	sc = get涉害(tp, sk, keList)
	chu = sc[0].Name()
	zhong = sc[1].Name()
	mo = sc[2].Name()
	if chu != "辰" || zhong != "申" || mo != "子" {
		t.Fatalf("%s日，未将卯时，三传：辰、申、子，非：`%s、%s、%s`", gz.Name(), chu, zhong, mo)
	}

	t.Log("戊辰日，丑将午时，复等课")
	戊 := 甲.Plus(4)
	辰 = 子.Plus(4)
	丑 = 子.Plus(1)
	午 = 子.Plus(6)
	gz, err = ganzhiwuxin.NewGanZhi(戊, 辰)
	if err != nil {
		t.Fatal(err)
	}
	tp = TianPan{丑, 午}
	sk = NewSiKe(tp, gz)
	//一、四课涉害
	keList = []int{1, 4}
	sc = get涉害(tp, sk, keList)
	chu = sc[0].Name()
	zhong = sc[1].Name()
	mo = sc[2].Name()
	if chu != "子" || zhong != "未" || mo != "寅" {
		t.Fatalf("%s日，丑将午时，三传：辰、申、子，非：`%s、%s、%s`", gz.Name(), chu, zhong, mo)
	}
}

func TestGet比用(t *testing.T) {
	t.Log("测试比用课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("甲寅日，寅将酉时")
	yueJing := 子.Plus(2)
	shi := 子.Plus(-3)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(甲, yueJing)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	keList := []int{1, 2}
	sc := get比用(tp, sk, keList)
	assert.Equalf("子", sc[0].Name(), "%s日，%s将%s，初传`子`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("巳", sc[1].Name(), "%s日，%s将%s，中传`子`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("戌", sc[2].Name(), "%s日，%s将%s，末传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("壬辰日，辰将巳时")
	yueJing = 子.Plus(4)
	shi = 子.Plus(5)
	g := 甲.Plus(-2)
	z := 子.Plus(4)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	keList = []int{1, 3}
	sc = get比用(tp, sk, keList)
	assert.Equalf("戌", sc[0].Name(), "%s日，%s将%s，初传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("酉", sc[1].Name(), "%s日，%s将%s，中传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("申", sc[2].Name(), "%s日，%s将%s，末传`申`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("辛卯日，亥将未时，涉害")
	yueJing = 子.Plus(-1)
	shi = 子.Plus(7)
	g = 甲.Plus(7)
	z = 子.Plus(3)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	keList = []int{3, 4}
	sc = get比用(tp, sk, keList)
	assert.Equalf("亥", sc[0].Name(), "%s日，%s将%s，初传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("卯", sc[1].Name(), "%s日，%s将%s，中传`卯`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("未", sc[2].Name(), "%s日，%s将%s，末传`未`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestGet贼克(t *testing.T) {
	t.Log("测试贼克课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	// 没有贼克
	t.Log("乙未日，亥将子时，无贼克")
	yueJing := 子.Plus(-1)
	shi := 子.Plus(0)
	g := 甲.Plus(1)
	z := 子.Plus(7)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := get贼克(tp, sk)
	assert.Equalf(0, len(sc), "%s日，%s将%s，非贼克课传。", gz.Name(), yueJing.Name(), shi.Name())

	t.Log("癸亥日，亥将午时，重审")
	yueJing = 子.Plus(-1)
	shi = 子.Plus(6)
	g = 甲.Plus(-1)
	z = 子.Plus(-1)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get贼克(tp, sk)
	assert.Equalf("午", sc[0].Name(), "%s日，%s将%s，初传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("亥", sc[1].Name(), "%s日，%s将%s，中传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("辰", sc[2].Name(), "%s日，%s将%s，末传`辰`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("丁丑，子将申时，元首课")
	g = 甲.Plus(3)
	z = 子.Plus(1)
	yueJing = 子.Plus(8)
	shi = 子.Plus(0)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get贼克(tp, sk)
	assert.Equalf("巳", sc[0].Name(), "%s日，%s将%s，初传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("丑", sc[1].Name(), "%s日，%s将%s，中传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("酉", sc[2].Name(), "%s日，%s将%s，末传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("甲戌，亥将戌时，比用课")
	g = 甲.Plus(0)
	z = 子.Plus(-2)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(-2)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get贼克(tp, sk)
	assert.Equalf("辰", sc[0].Name(), "%s日，%s将%s，初传`辰`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("巳", sc[1].Name(), "%s日，%s将%s，中传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("午", sc[2].Name(), "%s日，%s将%s，末传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestGet反吟(t *testing.T) {
	t.Log("测试反吟课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("甲寅，亥将巳时，有贼克反吟课")
	g := 甲.Plus(0)
	z := 子.Plus(2)
	yueJing := 子.Plus(-1)
	shi := 子.Plus(5)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := get反呤(tp, sk)
	assert.Equalf("寅", sc[0].Name(), "%s日，%s将%s，初传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("申", sc[1].Name(), "%s日，%s将%s，中传`申`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("寅", sc[2].Name(), "%s日，%s将%s，末传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("辛丑日，亥将巳时，无贼克反吟课")
	g = 甲.Plus(7)
	z = 子.Plus(1)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(5)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get反呤(tp, sk)
	assert.Equalf("亥", sc[0].Name(), "%s日，%s将%s，初传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("未", sc[1].Name(), "%s日，%s将%s，中传`未`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("辰", sc[2].Name(), "%s日，%s将%s，末传`辰`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestGet八专(t *testing.T) {
	t.Log("测试八专课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("丁未日，亥将申时，阴日八专课")
	g := 甲.Plus(3)
	z := 子.Plus(7)
	yueJing := 子.Plus(-1)
	shi := 子.Plus(8)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := get八专(sk)
	assert.Equalf("亥", sc[0].Name(), "%s日，%s将%s，初传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("戌", sc[1].Name(), "%s日，%s将%s，中传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("戌", sc[2].Name(), "%s日，%s将%s，末传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("甲寅日，丑将辰时，阳日八专")
	g = 甲.Plus(0)
	z = 子.Plus(2)
	yueJing = 子.Plus(1)
	shi = 子.Plus(4)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get八专(sk)
	assert.Equalf("丑", sc[0].Name(), "%s日，%s将%s，初传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("亥", sc[1].Name(), "%s日，%s将%s，中传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("亥", sc[2].Name(), "%s日，%s将%s，末传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("己未日，酉将未时，独足格")
	g = 甲.Plus(5)
	z = 子.Plus(7)
	yueJing = 子.Plus(9)
	shi = 子.Plus(7)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get八专(sk)
	assert.Equalf("酉", sc[0].Name(), "%s日，%s将%s，初传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("酉", sc[1].Name(), "%s日，%s将%s，中传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("酉", sc[2].Name(), "%s日，%s将%s，末传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestGet遥克(t *testing.T) {
	t.Log("测试遥克课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("壬辰日，巳将寅时，比用遥克课")
	g := 甲.Plus(-2)
	z := 子.Plus(4)
	yueJing := 子.Plus(5)
	shi := 子.Plus(2)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := get遥克(tp, sk)
	assert.Equalf("戌", sc[0].Name(), "%s日，%s将%s，初传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("丑", sc[1].Name(), "%s日，%s将%s，中传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("辰", sc[2].Name(), "%s日，%s将%s，末传`辰`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("壬申日，亥将申时，干遥克课")
	g = 甲.Plus(-2)
	z = 子.Plus(8)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(8)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get遥克(tp, sk)
	assert.Equalf("巳", sc[0].Name(), "%s日，%s将%s，初传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("申", sc[1].Name(), "%s日，%s将%s，中传`申`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("亥", sc[2].Name(), "%s日，%s将%s，末传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestGet昴星(t *testing.T) {
	t.Log("测试昴星课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("戊寅日，辰将子时，阳昴星课")
	g := 甲.Plus(4)
	z := 子.Plus(2)
	yueJing := 子.Plus(4)
	shi := 子.Plus(0)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := get昴星(tp, sk)
	assert.Equalf("丑", sc[0].Name(), "%s日，%s将%s，初传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("午", sc[1].Name(), "%s日，%s将%s，中传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("酉", sc[2].Name(), "%s日，%s将%s，末传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("丁亥日，巳将寅时，阴日昴星")
	g = 甲.Plus(3)
	z = 子.Plus(-1)
	yueJing = 子.Plus(5)
	shi = 子.Plus(2)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get昴星(tp, sk)
	assert.Equalf("午", sc[0].Name(), "%s日，%s将%s，初传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("戌", sc[1].Name(), "%s日，%s将%s，中传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("寅", sc[2].Name(), "%s日，%s将%s，末传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestGet别责(t *testing.T) {
	t.Log("测试别责课")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("戊辰日，亥将戌时，阳别责课")
	g := 甲.Plus(4)
	z := 子.Plus(4)
	yueJing := 子.Plus(-1)
	shi := 子.Plus(-2)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := get别责(tp, sk)
	assert.Equalf("寅", sc[0].Name(), "%s日，%s将%s，初传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("午", sc[1].Name(), "%s日，%s将%s，中传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("午", sc[2].Name(), "%s日，%s将%s，末传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())

	t.Log("辛酉日，子将丑时，阴别责")
	g = 甲.Plus(7)
	z = 子.Plus(-3)
	yueJing = 子.Plus(0)
	shi = 子.Plus(1)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	sc = get别责(tp, sk)
	assert.Equalf("子", sc[0].Name(), "%s日，%s将%s，初传`子`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[0].Name())
	assert.Equalf("酉", sc[1].Name(), "%s日，%s将%s，中传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[1].Name())
	assert.Equalf("酉", sc[2].Name(), "%s日，%s将%s，末传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc[2].Name())
}

func TestSanChuan(t *testing.T) {
	t.Log("测试获取三传函数：sanChuan()")
	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("下克上")
	t.Log("丙戌日，申将巳时")
	g := 甲.Plus(2)
	z := 子.Plus(-2)
	yueJing := 子.Plus(8)
	shi := 子.Plus(5)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	chu, zhong, mo := sanChuan(tp, sk)
	assert.Equalf("申", chu.Name(), "%s日，%s将%s，初传`申`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("亥", zhong.Name(), "%s日，%s将%s，中传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("寅", mo.Name(), "%s日，%s将%s，末传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("上克下")
	t.Log("丁丑日，申将子时")
	g = 甲.Plus(3)
	z = 子.Plus(1)
	yueJing = 子.Plus(8)
	shi = 子.Plus(0)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("巳", chu.Name(), "%s日，%s将%s，初传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("丑", zhong.Name(), "%s日，%s将%s，中传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("酉", mo.Name(), "%s日，%s将%s，末传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("比用")
	t.Log("甲戌日，亥将戌时")
	g = 甲.Plus(0)
	z = 子.Plus(-2)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(-2)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("辰", chu.Name(), "%s日，%s将%s，初传`辰`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("巳", zhong.Name(), "%s日，%s将%s，中传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("午", mo.Name(), "%s日，%s将%s，末传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("涉害")
	t.Log("辛卯日，亥将未时")
	g = 甲.Plus(7)
	z = 子.Plus(3)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(7)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("亥", chu.Name(), "%s日，%s将%s，初传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("卯", zhong.Name(), "%s日，%s将%s，中传`卯`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("未", mo.Name(), "%s日，%s将%s，末传`未`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("遥克")
	t.Log("戊辰日，亥将未时")
	g = 甲.Plus(4)
	z = 子.Plus(4)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(7)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("子", chu.Name(), "%s日，%s将%s，初传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("辰", zhong.Name(), "%s日，%s将%s，中传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("申", mo.Name(), "%s日，%s将%s，末传`酉`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("昴星")
	t.Log("乙未日，亥将子时")
	g = 甲.Plus(1)
	z = 子.Plus(7)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(0)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("戌", chu.Name(), "%s日，%s将%s，初传`戌`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("卯", zhong.Name(), "%s日，%s将%s，中传`卯`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("午", mo.Name(), "%s日，%s将%s，末传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("伏吟")
	t.Log("丙子日，亥将亥时")
	g = 甲.Plus(2)
	z = 子.Plus(0)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(-1)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("巳", chu.Name(), "%s日，%s将%s，初传`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("申", zhong.Name(), "%s日，%s将%s，中传`申`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("寅", mo.Name(), "%s日，%s将%s，末传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("反吟")
	t.Log("丁丑日，亥将巳时")
	g = 甲.Plus(3)
	z = 子.Plus(1)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(5)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("亥", chu.Name(), "%s日，%s将%s，初传`亥`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("未", zhong.Name(), "%s日，%s将%s，中传`未`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("丑", mo.Name(), "%s日，%s将%s，末传`丑`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())

	t.Log("别责")
	t.Log("戊辰日，亥将戌时")
	g = 甲.Plus(4)
	z = 子.Plus(4)
	yueJing = 子.Plus(-1)
	shi = 子.Plus(-2)
	tp = TianPan{yueJing, shi}
	gz, err = ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk = NewSiKe(tp, gz)
	chu, zhong, mo = sanChuan(tp, sk)
	assert.Equalf("寅", chu.Name(), "%s日，%s将%s，初传`寅`，非%s", gz.Name(), yueJing.Name(), shi.Name(), chu.Name())
	assert.Equalf("午", zhong.Name(), "%s日，%s将%s，中传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), zhong.Name())
	assert.Equalf("午", mo.Name(), "%s日，%s将%s，末传`午`，非%s", gz.Name(), yueJing.Name(), shi.Name(), mo.Name())
}

func TestNewSanChuan(t *testing.T) {
	t.Log("测试方法：NewSanChuan")

	assert := assert.New(t)
	甲, err := ganzhiwuxin.NewTianGan("甲")
	assert.NoError(err)
	子, err := ganzhiwuxin.NewDiZhi("子")
	assert.NoError(err)

	t.Log("辛丑日，未将寅时")
	g := 甲.Plus(7)
	z := 子.Plus(-3)
	yueJing := 子.Plus(7)
	shi := 子.Plus(2)
	tp := TianPan{yueJing, shi}
	gz, err := ganzhiwuxin.NewGanZhi(g, z)
	assert.NoError(err)
	sk := NewSiKe(tp, gz)
	sc := NewSanChuan(tp, sk)

	assert.Equalf("未", sc.chu.Name(), "%s日，%s将%s，初传`未`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.chu.Name())
	assert.Equalf("子", sc.zhong.Name(), "%s日，%s将%s，中传`子`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.zhong.Name())
	assert.Equalf("巳", sc.mo.Name(), "%s日，%s将%s，末传遁干`巳`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.mo.Name())

	assert.Equalf("己", sc.dunGan[0], "%s日，%s将%s，初传遁干`己`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.dunGan[0])
	assert.Equalf("", sc.dunGan[1], "%s日，%s将%s，中传遁干`无遁干`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.dunGan[1])
	assert.Equalf("丁", sc.dunGan[2], "%s日，%s将%s，末传`丁`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.dunGan[2])

	assert.Equalf("父", sc.liuQing[0], "%s日，%s将%s，初传六亲`父`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.liuQing[0])
	assert.Equalf("子", sc.liuQing[1], "%s日，%s将%s，中传六亲`子`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.liuQing[1])
	assert.Equalf("官", sc.liuQing[2], "%s日，%s将%s，末传六亲`官`，非%s", gz.Name(), yueJing.Name(), shi.Name(), sc.liuQing[2])

	scJson, err := json.Marshal(sc)
	if err != nil {
		t.Fatalf(err.Error())
	}
	m := map[string]interface{}{
		"chu":     "未",
		"zhong":   "子",
		"mo":      "巳",
		"dunGan":  [3]string{"己", "", "丁"},
		"liuQing": [3]string{"父", "子", "官"},
	}
	s, err := json.Marshal(m)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(string(s), string(scJson), "三传序列化失败")
}
