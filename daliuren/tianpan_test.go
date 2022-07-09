package daliuren

import (
	"encoding/json"
	"testing"

	"github.com/wlhyl/ganzhiwuxin"
)

func TestTianpan(t *testing.T) {
	t.Log("测试天盘")
	yueJiang, err := ganzhiwuxin.NewDiZhi("申")
	if err != nil {
		t.Fatalf(err.Error())
	}
	divinationTime, err := ganzhiwuxin.NewDiZhi("辰")
	if err != nil {
		t.Fatalf(err.Error())
	}
	zi, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		t.Fatalf(err.Error())
	}
	tp := TianPan{yueJiang, divinationTime}
	for i := 0; i < 12; i++ {
		d := zi.Plus(i)
		up := tp.up(d)
		if !up.Equals(zi.Plus(4 + i)) {
			t.Fatalf("%s加%s，%s上神是%s，而非%s", yueJiang.Name(), divinationTime.Name(), d.Name(), zi.Plus(4+i).Name(), up.Name())
		}
		down := tp.down(d)
		if !down.Equals(zi.Plus(8 + i)) {
			t.Fatalf("%s加%s，%s临%s，而非%s", yueJiang.Name(), divinationTime.Name(), d.Name(), zi.Plus(8+i).Name(), down.Name())
		}
	}

	json, err := json.Marshal(tp)
	if err != nil {
		t.Fatalf(err.Error())
	}
	s := "[\"辰\",\"巳\",\"午\",\"未\",\"申\",\"酉\",\"戌\",\"亥\",\"子\",\"丑\",\"寅\",\"卯\"]"
	if string(json) != s {
		t.Fatalf("天盘序列化失败")
	}
}
