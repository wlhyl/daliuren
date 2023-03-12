mod jigong;
mod sanchuan;
mod sike;
mod tianjiang;
mod tianpan;

use ganzhiwuxin::{DiZhi, GanZhi, TianGan};
use lunar_calendar::lunar_calendar;
use serde::Serialize;
use swe::{swe_calc_ut, swe_close, swe_julday, swe_set_ephe_path, swe_utc_time_zone};

use self::{sanchuan::SanChuan, sike::SiKe, tianjiang::TianJiangPan, tianpan::TianPan};

#[derive(Debug, Serialize)]
pub struct DaliurenShiPan {
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
    lunar_calendar: LunarCalendar, //农历日期
    solar_term_first: SolarTerm,   //节气
    solar_term_second: SolarTerm,  //节气
    si_zhu: SiZhu,                 //四柱
    kong: (String, String),        //空亡
    tian_pan: TianPan,             //天盘
    tian_jiang_pan: TianJiangPan,  // 天将盘
    sike: SiKe,                    // 四课
    san_chuan: SanChuan,           // 三传
    year_of_birth_gan_zhi: String, //本命
    xing_nian: String,             // 行年
    yue_jiang: String,
}

impl DaliurenShiPan {
    //起课时间： year, month, day, hour, minute, second
    // 出生年：yearOfBirth
    // 月将：yueJian,非空为自定义月将
    // 起课时辰： divinationTime
    // 昼夜： diurnal, 昼占:true
    // 性别： masculine， 男：true
    // 星历表位置：ephePath
    pub fn new(
        year: i32,
        month: u8,
        day: u8,
        hour: u8,
        minute: u8,
        second: u8,
        yue_jiang: Option<String>,
        divination_time: &str,
        diurnal: bool,
        year_of_birth: i32,
        masculine: bool,
        ephe_path: &str,
    ) -> Result<DaliurenShiPan, String> {
        // var p DaliurenShiPan

        // 将公历转换成农历
        let lunar_calendar = lunar_calendar(year, month, day, hour, minute, second, ephe_path)?;

        // 得到月将地支
        let yue_jian_zhi = if let Some(yue_jiang) = yue_jiang {
            DiZhi::new(&yue_jiang)?
        } else {
            let (y_utc, m_utc, d_utc, h_utc, mi_utc, sec_utc) = swe_utc_time_zone(
                year,
                month.into(),
                day.into(),
                hour.into(),
                minute.into(),
                second.into(),
                8.0,
            );
            let jd = swe_julday(
                y_utc,
                m_utc,
                d_utc,
                h_utc as f64 + mi_utc as f64 / 60.0 + sec_utc / 3600.0,
                swe::Calendar::Gregorian,
            );

            swe_set_ephe_path(ephe_path);

            //计算太阳黄道经度

            let xx = match swe_calc_ut(jd, swe::Planet::SUN, Default::default()) {
                Ok(xx) => xx,
                Err(error) => {
                    swe_close();
                    return Err(error);
                }
            };

            let sun_posi = xx[0];
            let n = (sun_posi / 30.0).floor() as isize;
            let 戌 = DiZhi::new("戌").unwrap();

            戌.plus(-n)
        };

        // 得到占时地支
        let divination_time_zhi =
            DiZhi::new(divination_time).map_err(|error| format!("占时地支不正确,{}", error))?;

        // 天盘
        let tian_pan = TianPan::new(&yue_jian_zhi, &divination_time_zhi);
        // 天将盘
        let tian_jiang_pan =
            TianJiangPan::new(&tian_pan, &lunar_calendar.lunar_day_gan_zhi.gan, diurnal);

        // 四课
        let sike = SiKe::new(&tian_pan, &lunar_calendar.lunar_day_gan_zhi);
        // 三传
        let san_chuan = SanChuan::new(&tian_pan, &sike);

        let lunar_calendar_struct = LunarCalendar {
            year: lunar_calendar.lunar_year.to_string(),
            month: lunar_calendar.lunar_month,
            day: lunar_calendar.lunar_day,
            time_gan_zhi: lunar_calendar.time_gan_zhi.to_string(),
        };

        let solar_term_first = SolarTerm {
            name: lunar_calendar.solar_term_first.name,
            year: lunar_calendar.solar_term_first.year,
            month: lunar_calendar.solar_term_first.month,
            day: lunar_calendar.solar_term_first.day,
            hour: lunar_calendar.solar_term_first.hour,
            minute: lunar_calendar.solar_term_first.minute,
            second: lunar_calendar.solar_term_first.second,
        };

        let solar_term_second = SolarTerm {
            name: lunar_calendar.solar_term_second.name,
            year: lunar_calendar.solar_term_second.year,
            month: lunar_calendar.solar_term_second.month,
            day: lunar_calendar.solar_term_second.day,
            hour: lunar_calendar.solar_term_second.hour,
            minute: lunar_calendar.solar_term_second.minute,
            second: lunar_calendar.solar_term_second.second,
        };

        let si_zhu = SiZhu {
            year: lunar_calendar.lunar_year_gan_zhi.to_string(),
            month: lunar_calendar.lunar_month_gan_zhi.to_string(),
            day: lunar_calendar.lunar_day_gan_zhi.to_string(),
            time: lunar_calendar.time_gan_zhi.to_string(),
        };

        // 计算空亡
        let kong = kong_wang(&lunar_calendar.lunar_day_gan_zhi);

        // 计算本命
        let 甲子 = GanZhi::default();

        let year_of_birth_gan_zhi = 甲子
            .plus((year_of_birth - 1984).try_into().unwrap())
            .to_string();
        // 计算行年
        let xing_nian = if masculine {
            let 丙寅 = 甲子.plus(2);
            丙寅
                .plus((year - year_of_birth).try_into().unwrap())
                .to_string()
        } else {
            let 壬申 = 甲子.plus(8);
            壬申
                .plus((year_of_birth - year).try_into().unwrap())
                .to_string()
        };
        // return p, nil
        Ok(Self {
            lunar_calendar: lunar_calendar_struct,
            solar_term_first,
            solar_term_second,
            si_zhu,
            kong,
            tian_pan,
            tian_jiang_pan,
            sike,
            san_chuan,
            year_of_birth_gan_zhi,
            xing_nian,
            yue_jiang: yue_jian_zhi.to_string(),
        })
    }
}

fn kong_wang(day: &GanZhi) -> (String, String) {
    let gan = &day.gan;
    let zhi = &day.zhi;

    let jia = TianGan::new("甲").unwrap();

    let delta: isize = gan.minus(&jia).into();

    let xun_shou = zhi.plus(-delta);

    (xun_shou.plus(-2).to_string(), xun_shou.plus(-1).to_string())
}
#[derive(Debug, Serialize)]
struct LunarCalendar {
    year: String,
    month: String,
    day: String,
    time_gan_zhi: String,
}

#[derive(Debug, Serialize)]
struct SolarTerm {
    name: String,
    year: i32,
    month: u8,
    day: u8,
    hour: u8,
    minute: u8,
    second: u8,
}

#[derive(Debug, Serialize)]
struct SiZhu {
    year: String,
    month: String,
    day: String,
    time: String,
}
