use std::vec;

use ganzhiwuxing::{
    DiZhi::{self, *},
    TianGan::{self, *},
};

// 天干的寄宫
pub fn ji_gong(tian_gan: &TianGan) -> DiZhi {
    let d = [寅, 辰, 巳, 未, 巳, 未, 申, 戌, 亥, 丑];
    let n: usize = tian_gan.minus(&甲).into();
    d[n].clone()
}

// 获取某地支所寄的天干
pub fn ji_gong_gan(d: &DiZhi) -> Vec<TianGan> {
    match d {
        丑 => vec![癸],
        寅 => vec![甲],
        辰 => vec![乙],
        巳 => vec![丙, 戊],
        未 => vec![丁, 己],
        申 => vec![庚],
        戌 => vec![辛],
        亥 => vec![壬],
        _ => vec![],
    }
}

#[cfg(test)]
mod tests {
    use ganzhiwuxing::{DiZhi::*, TianGan::*};

    use super::{ji_gong, ji_gong_gan};

    #[test]
    fn test_ji_gong() {
        // 测试天干的寄宫
        let jigong = [寅, 辰, 巳, 未, 巳, 未, 申, 戌, 亥, 丑];

        for i in 0..10 {
            let tg = 甲.plus(i);
            let d = ji_gong(&tg);
            assert_eq!(
                d, jigong[i as usize],
                "`{}`的寄宫是`{}`，非是`{}`",
                tg, jigong[i as usize], d
            )
        }
    }

    #[test]
    fn test_ji_gong_gan() {
        // 测试从地支获取所寄天干

        for i in 0..12 {
            let d = 子.plus(i);
            let g = ji_gong_gan(&d);
            match d {
                子 => assert_eq!(g.len(), 0, "子无所寄天干，非`%{:?}`", g),
                丑 => assert!(g.len() == 1 && g[0] == 癸, "丑所寄天干是`癸`，非`{:?}`", g),
                寅 => assert!(g.len() == 1 && g[0] == 甲, "寅所寄天干是`甲`，非`{:?}`", g),
                卯 => assert_eq!(g.len(), 0, "卯无所寄天干，非`{:?}`", g),
                辰 => assert!(g.len() == 1 && g[0] == 乙, "辰所寄天干是`乙`，非`{:?}`", g),
                巳 => assert!(
                    g.len() == 2 && g[0] == 丙 && g[1] == 戊,
                    "巳所寄天干是`丙、戊`，非`{:?}`",
                    g
                ),
                午 => assert_eq!(g.len(), 0, "午无所寄天干，非`{:?}`", g),
                未 => assert!(
                    g.len() == 2 && g[0] == 丁 && g[1] == 己,
                    "未所寄天干是`丁、己`，非`{:?}`",
                    g
                ),
                申 => assert!(g.len() == 1 && g[0] == 庚, "申所寄天干是`庚`，非`{:?}`", g),
                酉 => assert_eq!(g.len(), 0, "酉无所寄天干，非`{:?}`", g),
                戌 => assert!(g.len() == 1 && g[0] == 辛, "戌所寄天干是`辛`，非`{:?}`", g),
                _ => assert!(g.len() == 1 && g[0] == 壬, "亥所寄天干是`壬`，非`{:?}`", g),
            }
        }
    }
}
