use std::vec;

use ganzhiwuxin::{DiZhi, TianGan};

// 天干的寄宫
pub fn ji_gong(tian_gan: &TianGan) -> DiZhi {
    let 子 = DiZhi::default();
    let 甲 = TianGan::default();

    let 寅 = 子.plus(2);
    let 辰 = 寅.plus(2);
    let 巳 = 辰.plus(1);
    let 未 = 巳.plus(2);
    let 申 = 未.plus(1);
    let 戌 = 申.plus(2);
    let 亥 = 戌.plus(1);
    let 丑 = 亥.plus(2);
    let d = [&寅, &辰, &巳, &未, &巳, &未, &申, &戌, &亥, &丑];
    let n: usize = tian_gan.minus(&甲).into();
    d[n].clone()
}

// 获取某地支所寄的天干
pub fn ji_gong_gan(d: &DiZhi) -> Vec<TianGan> {
    let 甲 = TianGan::default();

    let d: &str = &d.name;

    match d {
        "丑" => vec![甲.plus(-1)],
        "寅" => vec![甲],
        "辰" => vec![甲.plus(1)],
        "巳" => vec![甲.plus(2), 甲.plus(4)],
        "未" => vec![甲.plus(3), 甲.plus(5)],
        "申" => vec![甲.plus(6)],
        "戌" => vec![甲.plus(7)],
        "亥" => vec![甲.plus(-2)],
        _ => vec![],
    }
}

#[cfg(test)]
mod tests {
    use ganzhiwuxin::{DiZhi, TianGan};

    use super::{ji_gong, ji_gong_gan};

    #[test]
    fn test_ji_gong() {
        // 测试天干的寄宫
        let jigong = ["寅", "辰", "巳", "未", "巳", "未", "申", "戌", "亥", "丑"];

        let 甲 = TianGan::new("甲").unwrap();
        for i in 0..10 {
            let tg = 甲.plus(i);
            let d = ji_gong(&tg);
            assert_eq!(
                d.name, jigong[i as usize],
                "`{}`的寄宫是`{}`，非是`{}`",
                tg, jigong[i as usize], d
            )
        }
    }

    #[test]
    fn test_ji_gong_gan() {
        // 测试从地支获取所寄天干
        let 子 = DiZhi::new("子").unwrap();

        for i in 0..12 {
            let d = &子.plus(i);
            let g = ji_gong_gan(&d);
            let name: &str = &d.name;
            match name {
                "子" => assert_eq!(g.len(), 0, "子无所寄天干，非`%{:?}`", g),
                "丑" => assert!(
                    g.len() == 1 && g[0].name == "癸",
                    "丑所寄天干是`癸`，非`{:?}`",
                    g
                ),
                "寅" => assert!(
                    g.len() == 1 && g[0].name == "甲",
                    "寅所寄天干是`甲`，非`{:?}`",
                    g
                ),
                "卯" => assert_eq!(g.len(), 0, "卯无所寄天干，非`{:?}`", g),
                "辰" => assert!(
                    g.len() == 1 && g[0].name == "乙",
                    "辰所寄天干是`乙`，非`{:?}`",
                    g
                ),
                "巳" => assert!(
                    g.len() == 2 && g[0].name == "丙" && g[1].name == "戊",
                    "巳所寄天干是`丙、戊`，非`{:?}`",
                    g
                ),
                "午" => assert_eq!(g.len(), 0, "午无所寄天干，非`{:?}`", g),
                "未" => assert!(
                    g.len() == 2 && g[0].name == "丁" && g[1].name == "己",
                    "未所寄天干是`丁、己`，非`{:?}`",
                    g
                ),
                "申" => assert!(
                    g.len() == 1 && g[0].name == "庚",
                    "申所寄天干是`庚`，非`{:?}`",
                    g
                ),
                "酉" => assert_eq!(g.len(), 0, "酉无所寄天干，非`{:?}`", g),
                "戌" => assert!(
                    g.len() == 1 && g[0].name == "辛",
                    "戌所寄天干是`辛`，非`{:?}`",
                    g
                ),
                _ => assert!(
                    g.len() == 1 && g[0].name == "壬",
                    "亥所寄天干是`壬`，非`{:?}`",
                    g
                ),
            }
        }
    }
}
