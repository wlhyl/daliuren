package daliuren

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlhyl/ganzhiwuxin"
)

func TestNewDaliurenShiPan(t *testing.T) {
	t.Log("测试方法：NewDaliurenShiPan")

	assert := assert.New(t)

	t.Log("2022-7-7 9:42:0")
	year := 2022
	month := 7
	day := 7
	hour := 9
	minute := 42
	second := 0
	yueJian := ""
	divinationTime := "巳"
	diurnal := true
	masculine := true
	yearOfBirth := 2000

	ephePath := os.Getenv("EPHE_PATH")
	if ephePath == "" {
		fmt.Println("EPHE_PATH must be specified")
		return
	}

	t.Log("男命")
	p, err := NewDaliurenShiPan(year, month, day, hour, minute, second, yueJian, divinationTime, diurnal, yearOfBirth, masculine, ephePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	jsonPan, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}
	pan := make(map[string]interface{})
	err = json.Unmarshal(jsonPan, &pan)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}

	// pan["lunarCalendar"]
	lunarCalendar := pan["lunarCalendar"].(map[string]interface{})
	assert.Equal("壬寅", fmt.Sprintf("%v", lunarCalendar["year"]))
	assert.Equal("六月", fmt.Sprintf("%v", lunarCalendar["month"]))
	assert.Equal("初九", fmt.Sprintf("%v", lunarCalendar["day"]))
	assert.Equal("巳", fmt.Sprintf("%v", lunarCalendar["timeGanZhi"]))

	solarTermFirst := pan["solarTermFirst"].(map[string]interface{})
	assert.Equal("芒种", fmt.Sprintf("%v", solarTermFirst["name"]))
	assert.Equal("2022", fmt.Sprintf("%v", solarTermFirst["year"]))
	assert.Equal("6", fmt.Sprintf("%v", solarTermFirst["month"]))
	assert.Equal("6", fmt.Sprintf("%v", solarTermFirst["day"]))
	assert.Equal("0", fmt.Sprintf("%v", solarTermFirst["hour"]))

	solarTermSecond := pan["solarTermSecond"].(map[string]interface{})
	assert.Equal("夏至", fmt.Sprintf("%v", solarTermSecond["name"]))
	assert.Equal("2022", fmt.Sprintf("%v", solarTermSecond["year"]))
	assert.Equal("6", fmt.Sprintf("%v", solarTermSecond["month"]))
	assert.Equal("21", fmt.Sprintf("%v", solarTermSecond["day"]))
	assert.Equal("17", fmt.Sprintf("%v", solarTermSecond["hour"]))

	ganZhi := pan["ganZhi"].(map[string]interface{})
	assert.Equal("壬寅", fmt.Sprintf("%v", ganZhi["year"]))
	assert.Equal("丙午", fmt.Sprintf("%v", ganZhi["month"]))
	assert.Equal("辛酉", fmt.Sprintf("%v", ganZhi["day"]))
	assert.Equal("癸巳", fmt.Sprintf("%v", ganZhi["time"]))

	kongWang := pan["kongWang"].([]interface{})
	assert.Equal("子", fmt.Sprintf("%v", kongWang[0]))
	assert.Equal("丑", fmt.Sprintf("%v", kongWang[1]))

	tianPan := pan["tianPan"].([]interface{})
	寅, err := ganzhiwuxin.NewDiZhi("寅")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(12, len(tianPan))
	for i := 0; i < 12; i++ {
		assert.Equal(寅.Plus(i).Name(), tianPan[i])
	}

	tianJiangPan := pan["tianJiangPan"].([]interface{})
	阴, err := NewTianJiang("阴")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(12, len(tianJiangPan))
	for i := 0; i < 12; i++ {
		assert.Equal(阴.Plus(i).Name(), tianJiangPan[i])
	}

	sike := pan["sike"].(map[string]interface{})
	assert.Equal("辛", fmt.Sprintf("%v", sike["gan"]))
	assert.Equal("子", fmt.Sprintf("%v", sike["ganYang"]))
	assert.Equal("寅", fmt.Sprintf("%v", sike["ganYing"]))
	assert.Equal("酉", fmt.Sprintf("%v", sike["zhi"]))
	assert.Equal("亥", fmt.Sprintf("%v", sike["zhiYang"]))
	assert.Equal("丑", fmt.Sprintf("%v", sike["zhiYing"]))

	sanChuan := pan["sanChuan"].(map[string]interface{})
	assert.Equal("丑", fmt.Sprintf("%v", sanChuan["chu"]))
	assert.Equal("卯", fmt.Sprintf("%v", sanChuan["zhong"]))
	assert.Equal("巳", fmt.Sprintf("%v", sanChuan["mo"]))

	dunGan := sanChuan["dunGan"].([]interface{})
	assert.Equal("", fmt.Sprintf("%v", dunGan[0]))
	assert.Equal("乙", fmt.Sprintf("%v", dunGan[1]))
	assert.Equal("丁", fmt.Sprintf("%v", dunGan[2]))

	liuQing := sanChuan["liuQing"].([]interface{})
	assert.Equal("父", fmt.Sprintf("%v", liuQing[0]))
	assert.Equal("财", fmt.Sprintf("%v", liuQing[1]))
	assert.Equal("官", fmt.Sprintf("%v", liuQing[2]))

	assert.Equal("庚辰", fmt.Sprintf("%v", pan["yearOfBirthGanZhi"]))
	assert.Equal("戊子", fmt.Sprintf("%v", pan["xingNian"]))

	// 测试女命
	masculine = false
	p, err = NewDaliurenShiPan(year, month, day, hour, minute, second, yueJian, divinationTime, diurnal, yearOfBirth, masculine, ephePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	jsonPan, err = json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}
	pan = make(map[string]interface{})
	err = json.Unmarshal(jsonPan, &pan)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}

	assert.Equal("庚戌", fmt.Sprintf("%v", pan["xingNian"]))

	// 自定义月将
	// 只需测试天盘是否正确即可，因为前已经测试可以正确排盘
	t.Log("测试自定义月将")
	yueJian = "申"
	p, err = NewDaliurenShiPan(year, month, day, hour, minute, second, yueJian, divinationTime, diurnal, yearOfBirth, masculine, ephePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	jsonPan, err = json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}
	pan = make(map[string]interface{})
	err = json.Unmarshal(jsonPan, &pan)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}

	// 地盘子，上神是卯
	tianPan = pan["tianPan"].([]interface{})
	assert.Equal(12, len(tianPan))
	for i := 0; i < 12; i++ {
		assert.Equal(寅.Plus(i+1).Name(), tianPan[i])
	}

	// 不正确的时间
	t.Log("测试不正确的时间")
	day = -1
	p, err = NewDaliurenShiPan(year, month, day, hour, minute, second, yueJian, divinationTime, diurnal, yearOfBirth, masculine, ephePath)
	assert.Error(err)
}
