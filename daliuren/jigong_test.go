package daliuren

import (
	"testing"

	"github.com/wlhyl/ganzhiwuxin"
)

func TestJiGong(t *testing.T) {
	t.Log("测试天干的寄宫")
	jigong := [...]string{"寅", "辰", "巳", "未", "巳", "未", "申", "戌", "亥", "丑"}

	甲, _ := ganzhiwuxin.NewTianGan("甲")
	for i := 0; i < 10; i++ {
		tg := 甲.Plus(i)
		if s := jiGong(tg).Name(); s != jigong[i] {
			t.Fatalf("`%s`的寄宫是`%s`，非是`%s`", tg.Name(), jigong[i], s)
		}
	}
}

func TestJiGongGan(t *testing.T) {
	t.Log("测试从地支获取所寄天干")
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 12; i++ {
		d := 子.Plus(i)
		g := jiGongGan(d)
		switch d.Name() {
		case "子":
			if len(g) != 0 {
				t.Fatalf("子无所寄天干，非`%v`", g)
			}
		case "丑":
			if len(g) != 1 || g[0].Name() != "癸" {
				t.Fatalf("丑所寄天干是`癸`，非`%v`", g)
			}

		case "寅":
			if len(g) != 1 || g[0].Name() != "甲" {
				t.Fatalf("丑所寄天干是`甲`，非`%v`", g)
			}
		case "卯":
			if len(g) != 0 {
				t.Fatalf("卯无所寄天干，非`%v`", g)
			}
		case "辰":
			if len(g) != 1 || g[0].Name() != "乙" {
				t.Fatalf("辰所寄天干是`乙`，非`%v`", g)
			}
		case "巳":
			if len(g) != 2 || g[0].Name() != "丙" || g[1].Name() != "戊" {
				t.Fatalf("巳所寄天干是`丙、戊`，非`%v`", g)
			}

		case "午":
			if len(g) != 0 {
				t.Fatalf("午无所寄天干，非`%v`", g)
			}
		case "未":
			if len(g) != 2 || g[0].Name() != "丁" || g[1].Name() != "己" {
				t.Fatalf("未所寄天干是`丁、己`，非`%v`", g)
			}
		case "申":
			if len(g) != 1 || g[0].Name() != "庚" {
				t.Fatalf("申所寄天干是`庚`，非`%v`", g)
			}
		case "酉":
			if len(g) != 0 {
				t.Fatalf("酉无所寄天干，非`%v`", g)
			}
		case "戌":
			if len(g) != 1 || g[0].Name() != "辛" {
				t.Fatalf("戌所寄天干是`辛`，非`%v`", g)
			}
		default:
			if len(g) != 1 || g[0].Name() != "壬" {
				t.Fatalf("亥所寄天干是`壬`，非`%v`", g)
			}
		}
	}
}
