use super::tianpan::TianPan;
use ganzhiwuxing::{
    DiZhi::{self, *},
    TianGan::{self, *},
};
use itertools::Itertools;
use serde::{ser::SerializeSeq, Serialize};
use std::fmt::Display;
use TianJiang::*;

// 天将
#[derive(Eq, Debug, PartialEq, Clone)]
pub enum TianJiang {
    贵,
    蛇,
    雀,
    合,
    勾,
    龙,
    空,
    虎,
    常,
    玄,
    阴,
    后,
}

impl TianJiang {
    pub fn plus(&self, other: isize) -> Self {
        let n = if other < 0 {
            other - other / 12 * 12 + 12
        } else {
            other
        };
        let n = (self.clone() as isize + n) % 12;

        match n {
            0 => TianJiang::贵,
            1 => TianJiang::蛇,
            2 => TianJiang::雀,
            3 => TianJiang::合,
            4 => TianJiang::勾,
            5 => TianJiang::龙,
            6 => TianJiang::空,
            7 => TianJiang::虎,
            8 => TianJiang::常,
            9 => TianJiang::玄,
            10 => TianJiang::阴,
            _ => TianJiang::后,
        }
    }

    // pub fn minus(&self, other: &TianJiang) -> u8 {
    //     // 返回整数值
    //     //  (self.num - other.num + 12) % 12
    //     (12 + self.num - other.num) % 12
    // }
}

impl Display for TianJiang {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            TianJiang::贵 => write!(f, "{}", "贵"),
            TianJiang::蛇 => write!(f, "{}", "蛇"),
            TianJiang::雀 => write!(f, "{}", "雀"),
            TianJiang::合 => write!(f, "{}", "合"),
            TianJiang::勾 => write!(f, "{}", "勾"),
            TianJiang::龙 => write!(f, "{}", "龙"),
            TianJiang::空 => write!(f, "{}", "空"),
            TianJiang::虎 => write!(f, "{}", "虎"),
            TianJiang::常 => write!(f, "{}", "常"),
            TianJiang::玄 => write!(f, "{}", "玄"),
            TianJiang::阴 => write!(f, "{}", "阴"),
            TianJiang::后 => write!(f, "{}", "后"),
        }
    }
}

#[derive(Debug)]
pub struct TianJiangPan {
    tian_yi_di_zhi: DiZhi, // 天乙所在的地支
    inverse: bool,         // 逆布为true
}

impl TianJiangPan {
    // tianPan:天盘
    // gan:日天干
    // diurnal:昼占为true，夜占为false
    pub fn new(tian_pan: &TianPan, gan: &TianGan, diurnal: bool) -> TianJiangPan {
        // 天乙所在的地支
        let tian_yi_di_zhi = tian_yi_di_zhi(gan, diurnal);

        let inverse = is_inverse(tian_pan, &tian_yi_di_zhi);

        TianJiangPan {
            tian_yi_di_zhi,
            inverse,
        }
    }
    // 获取某地支的天将
    pub fn up(&self, di_zhi: &DiZhi) -> TianJiang {
        let n = di_zhi.minus(&self.tian_yi_di_zhi);
        if self.inverse {
            贵.plus(isize::from(n) * -1)
        } else {
            贵.plus(n.into())
        }
    }

    // // 获取某天将所临地支
    // pub fn down(&self, tian_jiang: &TianJiang) -> DiZhi {
    //     let 贵 = TianJiang::new("贵").unwrap();
    //     let n = tian_jiang.minus(&贵);
    //     if self.inverse {
    //         self.tian_yi_di_zhi.plus(-1_isize * isize::from(n))
    //     } else {
    //         self.tian_yi_di_zhi.plus(n.into())
    //     }
    // }
}

impl Serialize for TianJiangPan {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let t: Vec<_> = (0isize..12)
            .map(|n| {
                let d = 子.plus(n);
                self.up(&d).to_string()
            })
            .collect_vec();

        let mut seq = serializer.serialize_seq(Some(t.len()))?;
        for e in t {
            seq.serialize_element(&e)?;
        }

        seq.end()
    }
}

fn tian_yi_di_zhi(gan: &TianGan, diurnal: bool) -> DiZhi {
    let 昼贵 = [未, 申, 酉, 亥, 丑, 子, 丑, 寅, 卯, 巳];
    let 夜贵 = [丑, 子, 亥, 酉, 未, 申, 未, 午, 巳, 卯];
    let n = gan.minus(&甲);
    if diurnal {
        昼贵[n as usize].clone()
    } else {
        夜贵[n as usize].clone()
    }
}

fn is_inverse(tian_pan: &TianPan, tian_yi_di_zhi: &DiZhi) -> bool {
    // 贵人地盘之支
    let gui_ren_di_pan = 子.plus(tian_yi_di_zhi.minus(&tian_pan.up(&子)).into());

    // 支 - 巳 必定 >=0，支 - 巳 <=5，则 巳=< 支 <= 戌
    if gui_ren_di_pan.minus(&巳) <= 5 {
        true
    } else {
        false
    }
}

#[cfg(test)]
mod tests {
    use super::{
        TianJiang::{self, *},
        TianJiangPan,
    };
    use crate::daliuren::shi_pan::tianpan::TianPan;
    use ganzhiwuxing::{DiZhi::*, TianGan::*};

    const TIAN_JIANG: [TianJiang; 12] = [贵, 蛇, 雀, 合, 勾, 龙, 空, 虎, 常, 玄, 阴, 后];

    #[test]
    fn test_tian_jiang_equals() {
        // 天将相等

        assert_eq!(贵, 贵);
        assert_eq!(蛇, 蛇);
        assert_eq!(雀, 雀);
        assert_eq!(合, 合);
        assert_eq!(勾, 勾);
        assert_eq!(龙, 龙);
        assert_eq!(空, 空);
        assert_eq!(虎, 虎);
        assert_eq!(常, 常);
        assert_eq!(玄, 玄);
        assert_eq!(阴, 阴);
        assert_eq!(后, 后);
    }

    #[test]
    fn testtian_jiang_not_equal() {
        // 天将不相等
        for (i, tj) in TIAN_JIANG.iter().enumerate() {
            let g0 = tj.clone();
            for j in 1..TIAN_JIANG.len() {
                let n = (i + j) % TIAN_JIANG.len();
                let g1 = TIAN_JIANG[n].clone();
                assert_ne!(g0, g1, "{} != {}", g0, g1);
            }
        }
    }

    #[test]
    //用数学归纳法
    fn test_tian_jiang_plus() {
        // 天将 + 整数
        for (i, v) in TIAN_JIANG.iter().enumerate() {
            let tg0 = TIAN_JIANG[(i + 1) % TIAN_JIANG.len()].clone();
            let tg1 = v.plus(1);
            assert_eq!(tg0, tg1, "{} + 1 = {}", v, tg0);
        }

        for (i, v) in TIAN_JIANG.iter().enumerate() {
            let tg0 = TIAN_JIANG[(i + TIAN_JIANG.len() - 1) % TIAN_JIANG.len()].clone();

            let tg1 = v.plus(-1);
            assert_eq!(tg0, tg1, "{} + (-1) = {}", v, tg0);
        }

        for g in TIAN_JIANG {
            assert_eq!(
                g.plus(99).plus(1),
                g.plus(100),
                "{} + 99 + 1 = {} + 100",
                g,
                g
            );

            assert_eq!(
                g.plus(-99).plus(-1),
                g.plus(-100),
                "{} + (-99) + (-1) = {} + (-100)",
                g,
                g
            );
        }
    }

    // #[test]
    // fn test_tian_jiang_minus() {
    //     // 天将 - 天将 = 正整数
    //     for v in TIAN_JIANG_NAME {
    //         let g0 = TianJiang::new(v).unwrap();
    //         for j in 0..TIAN_JIANG_NAME.len() {
    //             let g1 = g0.plus(j.try_into().unwrap());
    //             assert_eq!(j, g1.minus(&g0) as usize, "{} - {} = {}", g1, g0, j);
    //         }
    //     }
    // }

    #[test]
    fn test_new_tian_jiang_pan() {
        // "测试NewTianJiangPan

        let tp = TianPan::new(申, 辰);
        let tjpan = TianJiangPan::new(&tp, &甲, true);

        // 甲日昼贵在未，申将辰时
        assert_eq!(
            tjpan.tian_yi_di_zhi, 未,
            "甲日昼贵人在未，非是{}",
            tjpan.tian_yi_di_zhi
        );

        assert!(!tjpan.inverse, "甲日昼占，申将辰时，天将顺布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);

        // 甲日夜贵在未，申将辰时
        assert_eq!(
            tjpan.tian_yi_di_zhi, 丑,
            "甲日夜贵人在丑，非是{}",
            tjpan.tian_yi_di_zhi
        );

        assert!(tjpan.inverse, "甲日夜占，申将辰时，天将逆布");

        // 顺逆临界条件测试
        // 测试甲日贵人在辰
        let tp = TianPan::new(未, 辰);
        let tjpan = TianJiangPan::new(&tp, &甲, true);

        assert!(!tjpan.inverse, "甲日昼占，未将辰时，天将顺布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(tjpan.inverse, "甲日夜占，未将辰时，天将逆布");

        // 测试甲日贵人在巳
        let tp = TianPan::new(未, 巳);
        let tjpan = TianJiangPan::new(&tp, &甲, true);
        assert!(tjpan.inverse, "甲日昼占，未将巳时，天将逆布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(!tjpan.inverse, "甲日夜占，未将巳时，天将顺布");

        // 测试甲日贵人在戌
        let tp = TianPan::new(未, 戌);
        let tjpan = TianJiangPan::new(&tp, &甲, true);
        assert!(tjpan.inverse, "甲日昼占，未将戌时，天将逆布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(!tjpan.inverse, "甲日夜占，未将戌时，天将顺布");

        // 测试甲日贵人在亥
        let tp = TianPan::new(未, 亥);
        let tjpan = TianJiangPan::new(&tp, &甲, true);
        assert!(!tjpan.inverse, "甲日昼占，未将亥时，天将顺布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(tjpan.inverse, "甲日夜占，未将亥时，天将逆布");
    }

    // #[test]
    // fn test_tian_jiang_pan_up_and_down() {
    //     // 测试获取地支所乘天将和天将所临地支
    //     let 子 = DiZhi::new("子").unwrap();
    //     let 甲 = TianGan::new("甲").unwrap();
    //     // 甲日昼占，子将丑时
    //     let tp = TianPan {
    //         yue_jiang: 子.clone(),
    //         divination_time: 子.plus(1),
    //     };
    //     let tj_pan = TianJiangPan::new(&tp, &甲, true);

    //     // 子上所乘天将是 白虎，天将逆布
    //     let 虎 = TianJiang::new("虎").unwrap();

    //     for i in 0..12 {
    //         let d = 子.plus(i);
    //         let tj = 虎.plus(-i);
    //         let tj0 = tj_pan.up(&d);
    //         assert_eq!(tj0, tj, "`{}`所乘天将是`{}`，非是`{}`", d, tj, tj0);

    //         let d0 = tj_pan.down(&tj);
    //         assert_eq!(d0, d, "`{}`所临地支是`{}`，非是`{}`", tj.name, d, d0);
    //     }
    // }

    #[test]
    fn test_tian_jian_pan_json() {
        // t.Log("测试TianJiangPan序列化")
        let tp = TianPan::new(子, 丑);
        let tj_pan = TianJiangPan::new(&tp, &甲, true);
        let json = serde_json::to_string(&tj_pan);
        assert!(json.is_ok());
        let json = json.unwrap();
        let s =
            "[\"虎\",\"空\",\"龙\",\"勾\",\"合\",\"雀\",\"蛇\",\"贵\",\"后\",\"阴\",\"玄\",\"常\"]";
        assert_eq!(json, s, "天将盘序列化失败");
    }
}
