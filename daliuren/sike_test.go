package daliuren

import (
	"encoding/json"
	"testing"

	"github.com/wlhyl/ganzhiwuxin"
)

func TestNewSiKe(t *testing.T) {
	t.Log("测试NewSiKe()")
	子, _ := ganzhiwuxin.NewDiZhi("子")
	甲, _ := ganzhiwuxin.NewTianGan("甲")
	丑 := 子.Plus(1)
	甲子, _ := ganzhiwuxin.NewGanZhi(甲, 子)

	tp := TianPan{子, 丑}
	// 甲子日，子将丑
	sike := NewSiKe(tp, 甲子)
	if s := sike.gan.Name(); s != "甲" {
		t.Fatalf("甲子日，子将丑时，干是：`甲`，非是`%s`", s)
	}
	if s := sike.ganYang.Name(); s != "丑" {
		t.Fatalf("甲子日，子将丑时，干阳是：`丑`，非是`%s`", s)
	}
	if s := sike.ganYing.Name(); s != "子" {
		t.Fatalf("甲子日，子将丑时，干阴是：`子`，非是`%s`", s)
	}
	if s := sike.zhi.Name(); s != "子" {
		t.Fatalf("甲子日，子将丑时，支是：`子`，非是`%s`", s)
	}
	if s := sike.zhiYang.Name(); s != "亥" {
		t.Fatalf("甲子日，子将丑时，支阳是：`亥`，非是`%s`", s)
	}
	if s := sike.zhiYing.Name(); s != "戌" {
		t.Fatalf("甲子日，子将丑时，支阴是：`戌`，非是`%s`", s)
	}

	sikeJson, err := json.Marshal(sike)
	if err != nil {
		t.Fatalf(err.Error())
	}
	m := map[string]string{
		"gan":     "甲",
		"ganYang": "丑",
		"ganYing": "子",
		"zhi":     "子",
		"zhiYang": "亥",
		"zhiYing": "戌",
	}
	s, err := json.Marshal(m)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if string(sikeJson) != string(s) {
		t.Fatalf("四课序列化失败")
	}
}
