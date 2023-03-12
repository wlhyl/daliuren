use super::tianpan::TianPan;
use ganzhiwuxin::{DiZhi, TianGan};
use itertools::Itertools;
use serde::{ser::SerializeSeq, Serialize};
use std::fmt::Display;

const TIAN_JIANG_NUM_TO_NAME: [&str; 12] = [
    "贵", "蛇", "雀", "合", "勾", "龙", "空", "虎", "常", "玄", "阴", "后",
];

// 天将
#[derive(Eq, Debug)]
pub struct TianJiang {
    num: u8,
    pub name: String,
}

impl TianJiang {
    pub fn new(name: &str) -> Result<TianJiang, String> {
        if let Some(num) = TIAN_JIANG_NUM_TO_NAME.iter().position(|&s| s == name) {
            Ok(Self {
                name: name.to_string(),
                num: num as u8 + 1,
            })
        } else {
            Err(format!("{}不是正确的天将", name))
        }
    }

    pub fn plus(&self, other: isize) -> Self {
        let tmp = if other < 0 {
            other - other / 12 * 12 + 12
        } else {
            other
        };
        let mut tmp = (usize::from(self.num) + tmp as usize) % 12;
        if tmp == 0 {
            tmp = 12
        }
        Self::new(TIAN_JIANG_NUM_TO_NAME[tmp - 1]).unwrap()
    }

    // pub fn minus(&self, other: &TianJiang) -> u8 {
    //     // 返回整数值
    //     //  (self.num - other.num + 12) % 12
    //     (12 + self.num - other.num) % 12
    // }
}

impl Display for TianJiang {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.name)
    }
}

impl PartialEq for TianJiang {
    fn eq(&self, other: &Self) -> bool {
        self.num == other.num && self.name == other.name
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
        let 贵 = TianJiang::new("贵").unwrap();
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
        let zi = DiZhi::new("子").unwrap();
        let t: Vec<_> = (0isize..12)
            .map(|n| {
                let d = zi.plus(n);
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
    let 甲 = TianGan::default();
    let 子 = DiZhi::default();

    let 丑 = 子.plus(1);
    let 寅 = 丑.plus(1);
    let 卯 = 寅.plus(1);
    let 辰 = 卯.plus(1);
    let 巳 = 辰.plus(1);
    let 午 = 巳.plus(1);
    let 未 = 午.plus(1);
    let 申 = 未.plus(1);
    let 酉 = 申.plus(1);
    let 戌 = 酉.plus(1);
    let 亥 = 戌.plus(1);
    let 昼贵 = [&未, &申, &酉, &亥, &丑, &子, &丑, &寅, &卯, &巳];
    let 夜贵 = [&丑, &子, &亥, &酉, &未, &申, &未, &午, &巳, &卯];
    let n = gan.minus(&甲);
    if diurnal {
        昼贵[n as usize].clone()
    } else {
        夜贵[n as usize].clone()
    }
}

fn is_inverse(tian_pan: &TianPan, tian_yi_di_zhi: &DiZhi) -> bool {
    let 子 = DiZhi::new("子").unwrap();

    // 贵人地盘之支
    let gui_ren_di_pan = 子.plus(tian_yi_di_zhi.minus(&tian_pan.up(&子)).into());

    let 巳 = 子.plus(5);
    // 支 - 巳 必定 >=0，支 - 巳 <=5，则 巳=< 支 <= 戌
    if gui_ren_di_pan.minus(&巳) <= 5 {
        true
    } else {
        false
    }
}

#[cfg(test)]
mod tests {
    use super::{TianJiang, TianJiangPan};
    use crate::daliuren::shi_pan::tianpan::TianPan;
    use ganzhiwuxin::{DiZhi, TianGan};

    const TIAN_JIANG_NAME: [&str; 12] = [
        "贵", "蛇", "雀", "合", "勾", "龙", "空", "虎", "常", "玄", "阴", "后",
    ];

    #[test]
    fn test_new_tian_jiang() {
        // NewTianJiang()方法"
        // 正确创建TianJian

        for v in TIAN_JIANG_NAME {
            let w = TianJiang::new(v);
            assert!(w.is_ok());
            let w = w.unwrap();
            assert_eq!(format!("{}", w), v);
        }

        // 不正确天将创建
        let w = TianJiang::new("否");
        assert!(w.is_err(), "`否`不是天将");
    }

    #[test]
    fn test_tian_jiang_equals() {
        // 天将相等
        for name in TIAN_JIANG_NAME {
            let tg0 = TianJiang::new(name).unwrap();
            let tg1 = TianJiang::new(name).unwrap();
            assert_eq!(tg0, tg1, "{} 与 {} 相等", tg0, tg1);
        }
    }

    #[test]
    fn testtian_jiang_not_equal() {
        // 天将不相等
        for (i, v) in TIAN_JIANG_NAME.iter().enumerate() {
            let g0 = TianJiang::new(v).unwrap();
            for j in 1..TIAN_JIANG_NAME.len() {
                let n = (i + j) % TIAN_JIANG_NAME.len();
                let g1 = TianJiang::new(TIAN_JIANG_NAME[n]).unwrap();
                assert_ne!(g0, g1, "{} != {}", g0, g1);
            }
        }
    }

    #[test]
    //用数学归纳法
    fn test_tian_jiang_plus() {
        // 天将 + 整数
        for (i, v) in TIAN_JIANG_NAME.iter().enumerate() {
            let tg0 = TianJiang::new(TIAN_JIANG_NAME[(i + 1) % TIAN_JIANG_NAME.len()]).unwrap();
            let tg1 = TianJiang::new(v).unwrap();
            let tg1 = tg1.plus(1);
            assert_eq!(tg0, tg1, "{} + 1 = {}", tg1, tg0);
        }

        for (i, v) in TIAN_JIANG_NAME.iter().enumerate() {
            let tg0 = TianJiang::new(
                TIAN_JIANG_NAME[(i + TIAN_JIANG_NAME.len() - 1) % TIAN_JIANG_NAME.len()],
            )
            .unwrap();
            let tg1 = TianJiang::new(v).unwrap();
            let tg1 = tg1.plus(-1);
            assert_eq!(tg0, tg1, "{} + (-1) = {}", tg1, tg0);
        }

        for it in TIAN_JIANG_NAME {
            let g = TianJiang::new(it).unwrap();
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
        let 申 = DiZhi::new("申").unwrap();

        let 辰 = DiZhi::new("辰").unwrap();

        let 甲 = TianGan::new("甲").unwrap();

        let tp = TianPan {
            yue_jiang: 申.clone(),
            divination_time: 辰.clone(),
        };
        let tjpan = TianJiangPan::new(&tp, &甲, true);

        // 甲日昼贵在未，申将辰时
        assert_eq!(
            tjpan.tian_yi_di_zhi,
            申.plus(-1),
            "甲日昼贵人在未，非是{}",
            tjpan.tian_yi_di_zhi
        );

        assert!(!tjpan.inverse, "甲日昼占，申将辰时，天将顺布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);

        // 甲日夜贵在未，申将辰时
        assert_eq!(
            tjpan.tian_yi_di_zhi,
            申.plus(5),
            "甲日夜贵人在丑，非是{}",
            tjpan.tian_yi_di_zhi
        );

        assert!(tjpan.inverse, "甲日夜占，申将辰时，天将逆布");

        // 顺逆临界条件测试
        // 测试甲日贵人在辰
        let 未 = 申.plus(-1);
        let tp = TianPan {
            yue_jiang: 未.clone(),
            divination_time: 辰,
        };
        let tjpan = TianJiangPan::new(&tp, &甲, true);

        assert!(!tjpan.inverse, "甲日昼占，未将辰时，天将顺布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(tjpan.inverse, "甲日夜占，未将辰时，天将逆布");

        // 测试甲日贵人在巳
        let 巳 = 申.plus(-3);
        let tp = TianPan {
            yue_jiang: 未.clone(),
            divination_time: 巳,
        };
        let tjpan = TianJiangPan::new(&tp, &甲, true);
        assert!(tjpan.inverse, "甲日昼占，未将巳时，天将逆布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(!tjpan.inverse, "甲日夜占，未将巳时，天将顺布");

        // 测试甲日贵人在戌
        let 戌 = 申.plus(2);
        let tp = TianPan {
            yue_jiang: 未.clone(),
            divination_time: 戌,
        };
        let tjpan = TianJiangPan::new(&tp, &甲, true);
        assert!(tjpan.inverse, "甲日昼占，未将戌时，天将逆布");

        let tjpan = TianJiangPan::new(&tp, &甲, false);
        assert!(!tjpan.inverse, "甲日夜占，未将戌时，天将顺布");

        // 测试甲日贵人在亥
        let 亥 = 申.plus(3);
        let tp = TianPan {
            yue_jiang: 未,
            divination_time: 亥,
        };
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
        let 子 = DiZhi::new("子").unwrap();
        let 丑 = 子.plus(1);
        let 甲 = TianGan::new("甲").unwrap();
        let tp = TianPan::new(&子, &丑);
        let tj_pan = TianJiangPan::new(&tp, &甲, true);
        let json = serde_json::to_string(&tj_pan);
        assert!(json.is_ok());
        let json = json.unwrap();
        let s =
            "[\"虎\",\"空\",\"龙\",\"勾\",\"合\",\"雀\",\"蛇\",\"贵\",\"后\",\"阴\",\"玄\",\"常\"]";
        assert_eq!(json, s, "天将盘序列化失败");
    }
}
