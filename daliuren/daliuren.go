package daliuren

import (
	"fmt"
	"math"
	"strings"

	"github.com/mshafiee/swephgo"
	"github.com/wlhyl/ganzhiwuxin"
	lunarCalendar "github.com/wlhyl/lunar-calendar"
)

type DaliurenShiPan struct {
	// Year            int
	// Month           int
	// Day             int
	// Hour            int
	// Minute          int
	// Second          int
	// YueJian        string
	// DivinationTime string
	// Diurnal         bool //昼占：true,夜占:false
	// Describe        string
	// Masculine       bool
	// YearOfBirth     int
	LunarCalendar     map[string]string `json:"lunarCalendar"`     //农历日期
	SolarTermFirst    map[string]string `json:"solarTermFirst"`    //节气
	SolarTermSecond   map[string]string `json:"solarTermSecond"`   //节气
	GanZhi            map[string]string `json:"ganZhi"`            //四柱
	Kong              [2]string         `json:"kongWang"`          //空亡
	TianPan           TianPan           `json:"tianPan"`           //天盘
	TianJiangPan      TianJiangPan      `json:"tianJiangPan"`      // 天将盘
	Sike              SiKe              `json:"sike"`              // 四课
	SanChuan          SanChuan          `json:"sanChuan"`          // 三传
	YearOfBirthGanZhi string            `json:"yearOfBirthGanZhi"` //本命
	XingNian          string            `json:"xingNian"`          // 行年
}

//起课时间： year, month, day, hour, minute, second
// 出生年：yearOfBirth
// 月将：yueJian,非空为自定义月将
// 起课时辰： divinationTime
// 昼夜： diurnal, 昼占:true
// 性别： masculine， 男：true
// 星历表位置：ephePath
func NewDaliurenShiPan(year, month, day, hour, minute, second int, yueJian, divinationTime string, diurnal bool, yearOfBirth int, masculine bool, ephePath string) (DaliurenShiPan, error) {
	var p DaliurenShiPan

	// 将公历转换成农历
	lunarCalendar, err := lunarCalendar.ConvertToLunarCalendar(year, month, day, hour, minute, second)
	if err != nil {
		return p, err
	}

	// 得到月将地支
	var yueJianZhi ganzhiwuxin.DiZhi
	if yueJian != "" {
		yueJianZhi, err = ganzhiwuxin.NewDiZhi(yueJian)
		if err != nil {
			return p, fmt.Errorf("月将地支不正确,%v", err.Error())
		}
	} else {
		// 计算月将
		yUtc := make([]int, 1)
		mUtc := make([]int, 1)
		dUtc := make([]int, 1)
		hUtc := make([]int, 1)
		miUtc := make([]int, 1)
		secUtc := make([]float64, 1)
		swephgo.UtcTimeZone(year, month, day, hour, minute, float64(second), 8.0, yUtc, mUtc, dUtc, hUtc, miUtc, secUtc)
		jd := swephgo.Julday(yUtc[0], mUtc[0], dUtc[0], float64(hUtc[0])+float64(miUtc[0])/60.0+secUtc[0]/3600.0, swephgo.SeGregCal)

		swephgo.SetEphePath([]byte(ephePath))

		//计算太阳黄道经度
		xx := make([]float64, 6)
		serr := make([]byte, 256)
		iflgret := swephgo.CalcUt(jd, swephgo.SeSun, swephgo.SeflgSwieph, xx, serr)

		serrString := strings.TrimRight(string(serr), string(byte(0)))
		if len(serrString) > 0 || iflgret < 0 {
			swephgo.Close()
			return p, fmt.Errorf("swe_calc_ut()错误。%s", serrString)
		}

		sunPosi := xx[0]
		n := int(math.Floor(sunPosi / 30))
		戌, _ := ganzhiwuxin.NewDiZhi("戌")
		if err != nil {
			return p, err
		}
		yueJianZhi = 戌.Plus(-n)
	}

	// 得到占时地支
	divinationTimeZhi, err := ganzhiwuxin.NewDiZhi(divinationTime)
	if err != nil {
		return p, fmt.Errorf("占时地支不正确,%v", err.Error())
	}

	// 天盘
	p.TianPan = TianPan{yueJianZhi, divinationTimeZhi}
	// 天将盘
	p.TianJiangPan, err = NewTianJiangPan(p.TianPan, lunarCalendar.LunarDayGanZhi.Gan(), diurnal)
	if err != nil {
		return p, err
	}
	// 四课
	p.Sike = NewSiKe(p.TianPan, lunarCalendar.LunarDayGanZhi)
	// 三传
	p.SanChuan = NewSanChuan(p.TianPan, p.Sike)

	p.LunarCalendar = make(map[string]string)
	p.LunarCalendar["year"] = lunarCalendar.LunarYear.Name()
	p.LunarCalendar["month"] = lunarCalendar.LunarMonth
	p.LunarCalendar["day"] = lunarCalendar.LunarDay
	p.LunarCalendar["timeGanZhi"] = lunarCalendar.TimeGanZhi.Zhi().Name()

	p.SolarTermFirst = make(map[string]string)
	p.SolarTermFirst["name"] = lunarCalendar.SolarTermFirst.Name
	p.SolarTermFirst["year"] = fmt.Sprint(lunarCalendar.SolarTermFirst.Year)
	p.SolarTermFirst["month"] = fmt.Sprint(lunarCalendar.SolarTermFirst.Month)
	p.SolarTermFirst["day"] = fmt.Sprint(lunarCalendar.SolarTermFirst.Day)
	p.SolarTermFirst["hour"] = fmt.Sprint(lunarCalendar.SolarTermFirst.Hour)
	p.SolarTermFirst["minute"] = fmt.Sprint(lunarCalendar.SolarTermFirst.Minute)
	p.SolarTermFirst["second"] = fmt.Sprint(lunarCalendar.SolarTermFirst.Second)

	p.SolarTermSecond = make(map[string]string)
	p.SolarTermSecond["name"] = lunarCalendar.SolarTermSecond.Name
	p.SolarTermSecond["year"] = fmt.Sprint(lunarCalendar.SolarTermSecond.Year)
	p.SolarTermSecond["month"] = fmt.Sprint(lunarCalendar.SolarTermSecond.Month)
	p.SolarTermSecond["day"] = fmt.Sprint(lunarCalendar.SolarTermSecond.Day)
	p.SolarTermSecond["hour"] = fmt.Sprint(lunarCalendar.SolarTermSecond.Hour)
	p.SolarTermSecond["minute"] = fmt.Sprint(lunarCalendar.SolarTermSecond.Minute)
	p.SolarTermSecond["second"] = fmt.Sprint(lunarCalendar.SolarTermSecond.Second)

	p.GanZhi = make(map[string]string)
	p.GanZhi["year"] = lunarCalendar.LunarYearGanZhi.Name()
	p.GanZhi["month"] = lunarCalendar.LunarMonthGanZhi.Name()
	p.GanZhi["day"] = lunarCalendar.LunarDayGanZhi.Name()
	p.GanZhi["time"] = lunarCalendar.TimeGanZhi.Name()

	// 计算空亡
	p.Kong = kongWang(lunarCalendar.LunarDayGanZhi)

	// 计算本命
	甲, err := ganzhiwuxin.NewTianGan("甲")
	if err != nil {
		return p, err
	}
	子, err := ganzhiwuxin.NewDiZhi("子")
	if err != nil {
		return p, err
	}
	甲子, err := ganzhiwuxin.NewGanZhi(甲, 子)
	if err != nil {
		return p, err
	}
	p.YearOfBirthGanZhi = 甲子.Plus(yearOfBirth - 1984).Name()
	// 计算行年
	if masculine {
		丙寅 := 甲子.Plus(2)
		p.XingNian = 丙寅.Plus(year - yearOfBirth).Name()
	} else {
		壬申 := 甲子.Plus(8)
		p.XingNian = 壬申.Plus(yearOfBirth - year).Name()
	}
	return p, nil
}

func kongWang(day ganzhiwuxin.GanZhi) [2]string {
	// d = []
	gan := day.Gan()
	zhi := day.Zhi()

	jia, _ := ganzhiwuxin.NewTianGan("甲")

	delta := gan.Minus(jia)

	xunShou := zhi.Plus(-delta)

	return [2]string{xunShou.Plus(-2).Name(), xunShou.Plus(-1).Name()}
}
