use ganzhiwuxing::{
    DiZhi::{self, *},
    GanZhi::{self, *},
    TianGan::{self, *},
};
use serde::Serialize;

use crate::daliuren::shi_pan::{
    sanchuan::{
        get八专, get别责, get反呤, get昴星, get比用, get涉害, get贼克, get遥克, has克, has贼,
        san_chuan, SanChuan,
    },
    sike::SiKe,
    tianpan::TianPan,
};

#[test]
fn test伏吟() {
    // 测试伏吟

    let tp = TianPan::new(子, 子);
    // 六甲日
    for i in 0..6 {
        let z = 子.plus(2 * i);
        let gz = GanZhi::new(&甲, &z).unwrap();
        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);

        assert!(
            sc.chu == 寅 && sc.zhong == 巳 && sc.mo == 申,
            "六甲日伏吟三传`寅，巳，申`，非是`{}`，`{}`，`{}`",
            sc.chu,
            sc.zhong,
            sc.mo
        );
    }

    // 六丙日
    for i in 0..6 {
        let z = 子.plus(2 * i);
        let gz = GanZhi::new(&甲.plus(2), &z).unwrap();

        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);
        let (chu, zhong, mo) = (sc.chu.to_string(), sc.zhong.to_string(), sc.mo.to_string());
        assert!(
            chu == "巳" && zhong == "申" && mo == "寅",
            "六丙日伏吟三传`巳，申，寅`，非是`{}`，`{}`，`{}`",
            chu,
            zhong,
            mo
        );
    }

    // info!("六戊日");
    for i in 0..6 {
        let z = 子.plus(2 * i);
        let gz = GanZhi::new(&甲.plus(4), &z).unwrap();

        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);
        let (chu, zhong, mo) = (sc.chu.to_string(), sc.zhong.to_string(), sc.mo.to_string());
        assert!(
            chu == "巳" && zhong == "申" && mo == "寅",
            "六戊日伏吟三传`巳，申，寅`，非是`{}`，`{}`，`{}`",
            chu,
            zhong,
            mo
        );
    }

    // t.Log("六庚日")
    for i in 0..6 {
        let z = 子.plus(2 * i);
        let gz = GanZhi::new(&甲.plus(6), &z).unwrap();

        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);
        let (chu, zhong, mo) = (sc.chu.to_string(), sc.zhong.to_string(), sc.mo.to_string());
        assert!(
            chu == "申" && zhong == "寅" && mo == "巳",
            "六庚日伏吟三传`申，寅，巳`，非是`{}`，`{}`，`{}`",
            chu,
            zhong,
            mo
        );
    }

    // t.Log("六壬日")
    for i in 0..6 {
        let z = 子.plus(2 * i);
        let gz = GanZhi::new(&甲.plus(8), &z).unwrap();

        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);

        assert!(
            sc.chu == 亥 && sc.zhong == z,
            "六壬日伏吟初传`亥`，中传`{}`, 非是`{}`, `{}`",
            z,
            sc.chu,
            sc.zhong
        );

        if z == 子 {
            assert_eq!(sc.mo, 卯, "壬子日伏呤，末传`卯`，非是`{}`", sc.mo);
        }

        if z == 寅 {
            assert_eq!(sc.mo, 巳, "壬寅日伏呤，末传`巳`，非是`{}`", sc.mo);
        }

        if z == 辰 {
            assert_eq!(sc.mo, 戌, "壬辰日伏呤，末传`戌`，非是`{}`", sc.mo);
        }
        if z == 午 {
            assert_eq!(sc.mo, 子, "壬午日伏呤，末传`子`，非是`{}`", sc.mo);
        }

        if z == 申 {
            assert_eq!(sc.mo, 寅, "壬申日伏呤，末传`寅`，非是`{}`", sc.mo);
        }

        if z == 戌 {
            assert_eq!(sc.mo, 未, "壬戌日伏呤，末传`未`，非是`{}`", sc.mo);
        }
    }

    // t.Log("六乙日")
    for i in 0..6 {
        let z = 子.plus(2 * i + 1);
        let gz = GanZhi::new(&甲.plus(1), &z).unwrap();

        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);

        assert!(
            sc.chu == 辰 && sc.zhong == z,
            "乙{}日伏吟初传`辰`，中传`{}`, 非是`{}`, `{}`",
            z,
            z,
            sc.chu,
            sc.zhong
        );

        let zhi = 子.plus(-2);
        if z == 丑 {
            assert_eq!(sc.mo, zhi, "乙丑日伏呤，末传`戌`，非是`{}`", sc.mo);
        }
        if z == 卯 {
            assert_eq!(sc.mo, 子, "乙卯日伏呤，末传`子`，非是`{}`", sc.mo);
        }
        let zhi = 子.plus(8);
        if z == 巳 {
            assert_eq!(sc.mo, zhi, "乙巳日伏呤，末传`申`，非是`{}`", sc.mo);
        }
        let zhi = 子.plus(1);
        if z == 未 {
            assert_eq!(sc.mo, zhi, "乙未日伏呤，末传`丑`，非是`{}`", sc.mo);
        }
        let zhi = 子.plus(3);
        if z == 酉 {
            assert_eq!(sc.mo, zhi, "乙酉日伏呤，末传`卯`，非是`{}`", sc.mo);
        }
        let zhi = 子.plus(5);
        if z == 亥 {
            assert_eq!(sc.mo, zhi, "乙亥日伏呤，末传`巳`，非是`{}`", sc.mo);
        }
    }

    // t.Log("六丁、己、辛日")
    //  丁=甲 + 3
    // 己 = 甲 + 5
    // 辛 = 甲 + 7
    for g in [甲.plus(3), 甲.plus(5), 甲.plus(7)] {
        for i in 0..6 {
            let z = 子.plus(2 * i + 1);
            let gz = GanZhi::new(&g, &z).unwrap();

            let sike = SiKe::new(&tp, &gz);
            let sc = SanChuan::new(&tp, &sike);

            // 丑
            if z == 子.plus(1) {
                assert!(
                    sc.chu == 子.plus(1) && sc.zhong == 子.plus(-2) && sc.mo == 子.plus(7),
                    "{}日伏吟三传`丑，戌，未`，非是`{}，{}，{}`",
                    gz,
                    sc.chu,
                    sc.zhong,
                    sc.mo
                );
            }

            // 卯
            if z == 子.plus(3) {
                assert!(
                    sc.chu == 子.plus(3) && sc.zhong == 子 && sc.mo == 子.plus(6),
                    "{}日伏吟三传`卯，子，午`，非是`{}，{}，{}`",
                    gz,
                    sc.chu,
                    sc.zhong,
                    sc.mo
                );
            }

            // 巳
            if z == 子.plus(5) {
                assert!(
                    sc.chu == 子.plus(5) && sc.zhong == 子.plus(8) && sc.mo == 子.plus(2),
                    "{}日伏吟三传`巳，申，寅`，非是`{}，{}，{}`",
                    gz,
                    sc.chu,
                    sc.zhong,
                    sc.mo
                );
            }

            // 未
            if z == 子.plus(7) {
                assert!(
                    sc.chu == 子.plus(7) && sc.zhong == 子.plus(1) && sc.mo == 子.plus(-2),
                    "{}日伏吟三传`未，丑，戌`，非是`{}，{}，{}`",
                    gz,
                    sc.chu,
                    sc.zhong,
                    sc.mo
                );
            }

            // 酉、亥日
            if z == 子.plus(9) || z == 子.plus(11) {
                assert!(
                    sc.chu == z && sc.zhong == sike.gan_yang,
                    "{}日伏吟，初传为`{}`，中传为`{}`, 非是`{}`，`{}`",
                    gz,
                    z,
                    sike.gan_yang,
                    sc.chu,
                    sc.zhong
                );

                // 丁、己日
                if g == 甲.plus(3) || g == 甲.plus(5) {
                    assert_eq!(
                        sc.mo,
                        子.plus(1),
                        "{}日伏吟，末传为`丑`，非是`{}`",
                        gz,
                        sc.mo
                    );
                }
                // 辛日
                if g == 甲.plus(7) {
                    assert_eq!(
                        sc.mo,
                        子.plus(7),
                        "{}日伏吟，末传为`未`，非是`{}`",
                        gz,
                        sc.mo
                    );
                }
            }
        }
    }

    // // t.Log("六癸日")
    for i in 0..6 {
        let z = 子.plus(2 * i + 1);
        let gz = GanZhi::new(&癸, &z).unwrap();

        let sike = SiKe::new(&tp, &gz);
        let sc = SanChuan::new(&tp, &sike);

        let (chu, zhong, mo) = (sc.chu.to_string(), sc.zhong.to_string(), sc.mo.to_string());
        assert!(
            chu == "丑" && zhong == "戌" && mo == "未",
            "六癸日伏吟，三传`丑、戌、未`，非是`{}, {}, {}`",
            chu,
            zhong,
            mo
        );
    }
}

#[test]
fn test_has贼() {
    // t.Log("测试函数has贼()")

    // t.Log("没有下克上的课")

    let tp = TianPan::new(子, 子);
    let sk = SiKe::new(&tp, &甲子);
    let ze = has贼(&sk);
    assert!(ze.is_empty(), "甲子日伏吟，四课中无下克上");

    // t.Log("有下克上的课")
    // 课中有重复
    // 辰将寅时，甲子日，四课：甲、辰、午、子、寅、辰
    // 返回值:[1]
    let tp = TianPan::new(子.plus(4), 子.plus(2));
    let sk = SiKe::new(&tp, &甲子);
    let ze = has贼(&sk);
    assert!(
        ze.len() == 1 && ze[0] == 1,
        "甲子日辰将寅时，取第一课，下克上"
    );
}

#[test]
fn test_has克() {
    // t.Log("测试函数has克()")

    // t.Log("没有上克下的课")

    let tp = TianPan::new(子, 子);
    let sk = SiKe::new(&tp, &甲子);
    let ze = has克(&sk);
    assert!(ze.is_empty(), "甲子日伏吟，四课中无上克下");

    // t.Log("有上克下的课")
    // 课中无重复
    // 甲子日，酉将寅时，四课：甲、酉、辰、子、未、寅
    // 返回值:[1, 3 , 4]
    let tp = TianPan::new(酉, 寅);
    let sk = SiKe::new(&tp, &甲子);
    let ze = has克(&sk);
    assert!(
        ze.len() == 3 && ze[0] == 1 && ze[1] == 3 && ze[2] == 4,
        "甲子日酉将寅时，第一、三、四课有下克上，非是`{:?}`",
        ze
    );

    // 课中有重复
    // 甲寅日，酉将寅时，四课：甲、酉、辰、寅、酉、辰
    // 返回值:[1, 3 , 4]

    let tp = TianPan::new(酉, 寅);
    let sk = SiKe::new(&tp, &甲寅);
    let ze = has克(&sk);

    assert!(
        ze.len() == 1 && ze[0] == 1,
        // skJson, err := json.Marshal(sk)
        // if err != nil {
        // 	t.Fatal(err)
        // }
        // t.Log(string(skJson))
        "甲寅日酉将寅时，取第一课的下克上，非是`{:?}`",
        ze
    );
}

#[test]
fn test_get涉害() {
    // t.Log("涉害课取三传")

    // t.Log("丁卯日，亥将丑时")

    let gz = 丁卯;

    let tp = TianPan::new(亥, 丑);
    let sk = SiKe::new(&tp, &gz);
    //三、四课涉害
    let ke_list = [3, 4];
    let sc = get涉害(&tp, &sk, &ke_list);
    let chu = sc[0].to_string();
    let zhong = sc[1].to_string();
    let mo = sc[2].to_string();
    assert!(
        chu == "亥" && zhong == "酉" && mo == "未",
        "丁卯日，亥将丑时，三传：亥、酉、未，非：`{}、{}、{}`",
        chu,
        zhong,
        mo
    );

    // t.Log("庚子日，申将戌时")

    let gz = 庚子;

    let tp = TianPan::new(申, 戌);
    let sk = SiKe::new(&tp, &gz);
    //一、三课涉害
    let ke_list = [1, 3];
    let sc = get涉害(&tp, &sk, &ke_list);
    let chu = sc[0].to_string();
    let zhong = sc[1].to_string();
    let mo = sc[2].to_string();
    assert!(
        chu == "午" && zhong == "辰" && mo == "寅",
        "{}日，申将戌时，三传：午、辰、寅，非：`{}、{}、{}`",
        gz,
        chu,
        zhong,
        mo
    );

    // t.Log("丙子日，亥将辰时，取孟")

    let gz = 丙子;

    let tp = TianPan::new(亥, 辰);
    let sk = SiKe::new(&tp, &gz);
    //一、四课涉害
    let ke_list = [1, 4];
    let sc = get涉害(&tp, &sk, &ke_list);
    let chu = sc[0].to_string();
    let zhong = sc[1].to_string();
    let mo = sc[2].to_string();
    assert!(
        chu == "子" && zhong == "未" && mo == "寅",
        "{}日，亥将辰时，三传：子、未、寅，非：`{}、{}、{}`",
        gz,
        chu,
        zhong,
        mo
    );

    // t.Log("庚午日，未将卯时，取仲")
    let gz = 庚午;

    let tp = TianPan::new(未, 卯);
    let sk = SiKe::new(&tp, &gz);
    //二、四课涉害
    let ke_list = [2, 4];
    let sc = get涉害(&tp, &sk, &ke_list);
    let chu = sc[0].to_string();
    let zhong = sc[1].to_string();
    let mo = sc[2].to_string();
    assert!(
        chu == "辰" && zhong == "申" && mo == "子",
        "{}日，未将卯时，三传：辰、申、子，非：`{}、{}、{}`",
        gz,
        chu,
        zhong,
        mo
    );

    // t.Log("丁卯日，未将午时，俱不比，取仲")

    let gz = 丁卯;

    let tp = TianPan::new(未, 午);
    let sk = SiKe::new(&tp, &gz);
    //二、四课涉害
    let ke_list = [1, 3];
    let sc = get涉害(&tp, &sk, &ke_list);
    let chu = sc[0].to_string();
    let zhong = sc[1].to_string();
    let mo = sc[2].to_string();
    assert!(
        chu == "辰" && zhong == "巳" && mo == "午",
        "{}日，未将午时，三传：辰、巳、午，非：`{}、{}、{}`",
        gz,
        chu,
        zhong,
        mo
    );

    // t.Log("戊辰日，丑将午时，复等课")
    let gz = 戊辰;

    let tp = TianPan::new(丑, 午);
    let sk = SiKe::new(&tp, &gz);
    //一、四课涉害
    let ke_list = [1, 4];
    let sc = get涉害(&tp, &sk, &ke_list);
    let chu = sc[0].to_string();
    let zhong = sc[1].to_string();
    let mo = sc[2].to_string();
    assert!(
        chu == "子" && zhong == "未" && mo == "寅",
        "{}日，丑将午时，三传：辰、申、子，非：`{}、{}、{}`",
        gz,
        chu,
        zhong,
        mo
    );
}

#[test]
fn test_get比用() {
    // t.Log("测试比用课")

    // t.Log("甲寅日，寅将酉时")
    let yue_jing = 寅;
    let shi = 酉;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 甲寅;

    let sk = SiKe::new(&tp, &gz);
    let ke_list = [1, 2];
    let sc = get比用(&tp, &sk, &ke_list);
    assert_eq!(
        "子",
        sc[0].to_string(),
        "{}日，{}将{}，初传`子`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "巳",
        sc[1].to_string(),
        "{}日，{}将{}，中传`子`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "戌",
        sc[2].to_string(),
        "{}日，{}将{}，末传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("壬辰日，辰将巳时")
    let yue_jing = 辰;
    let shi = 巳;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 壬辰;

    let sk = SiKe::new(&tp, &gz);
    let ke_list = [1, 3];
    let sc = get比用(&tp, &sk, &ke_list);
    assert_eq!(
        "戌",
        sc[0].to_string(),
        "{}日，{}将{}，初传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "酉",
        sc[1].to_string(),
        "{}日，{}将{}，中传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "申",
        sc[2].to_string(),
        "{}日，{}将{}，末传`申`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("辛卯日，亥将未时，涉害")
    let yue_jing = 亥;
    let shi = 未;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 辛卯;

    let sk = SiKe::new(&tp, &gz);
    let ke_list = [3, 4];
    let sc = get比用(&tp, &sk, &ke_list);
    assert_eq!(
        "亥",
        sc[0].to_string(),
        "{}日，{}将{}，初传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "卯",
        sc[1].to_string(),
        "{}日，{}将{}，中传`卯`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "未",
        sc[2].to_string(),
        "{}日，{}将{}，末传`未`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("庚午日，未将寅时")
    let yue_jing = 未;
    let shi = 寅;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 庚午;

    let sk = SiKe::new(&tp, &gz);
    let ke_list = [3, 4];
    let sc = get比用(&tp, &sk, &ke_list);
    assert_eq!(
        "辰",
        sc[0].to_string(),
        "{}日，{}将{}，初传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "酉",
        sc[1].to_string(),
        "{}日，{}将{}，中传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "寅",
        sc[2].to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_get贼克() {
    // t.Log("测试贼克课")

    // 没有贼克
    // t.Log("乙未日，亥将子时，无贼克")
    let yue_jing = 亥;
    let shi = 子;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 乙未;

    let sk = SiKe::new(&tp, &gz);
    let sc = get贼克(&tp, &sk);
    assert!(
        sc.is_empty(),
        "{}日，{}将{}，非贼克课传。",
        gz,
        yue_jing,
        shi
    );

    // t.Log("癸亥日，亥将午时，重审")
    let yue_jing = 亥;
    let shi = 午;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 癸亥;

    let sk = SiKe::new(&tp, &gz);
    let sc = get贼克(&tp, &sk);
    assert_eq!(
        "午",
        sc[0].to_string(),
        "{}日，{}将{}，初传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "亥",
        sc[1].to_string(),
        "{}日，{}将{}，中传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "辰",
        sc[2].to_string(),
        "{}日，{}将{}，末传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("丁丑，申将子时，元首课")
    let yue_jing = 申;
    let shi = 子;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 丁丑;

    let sk = SiKe::new(&tp, &gz);
    let sc = get贼克(&tp, &sk);
    assert_eq!(
        "巳",
        sc[0].to_string(),
        "{}日，{}将{}，初传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "丑",
        sc[1].to_string(),
        "{}日，{}将{}，中传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "酉",
        sc[2].to_string(),
        "{}日，{}将{}，末传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("甲戌，亥将戌时，比用课")
    let yue_jing = 亥;
    let shi = 戌;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 甲戌;

    let sk = SiKe::new(&tp, &gz);
    let sc = get贼克(&tp, &sk);
    assert_eq!(
        "辰",
        sc[0].to_string(),
        "{}日，{}将{}，初传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "巳",
        sc[1].to_string(),
        "{}日，{}将{}，中传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "午",
        sc[2].to_string(),
        "{}日，{}将{}，末传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_get反吟() {
    // t.Log("测试反吟课")

    // t.Log("甲寅，亥将巳时，有贼克反吟课")
    let yue_jing = 亥;
    let shi = 巳;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 甲寅;

    let sk = SiKe::new(&tp, &gz);
    let sc = get反呤(&tp, &sk);
    assert_eq!(
        "寅",
        sc[0].to_string(),
        "{}日，{}将{}，初传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "申",
        sc[1].to_string(),
        "{}日，{}将{}，中传`申`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "寅",
        sc[2].to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("辛丑日，亥将巳时，无贼克反吟课")
    let yue_jing = 亥;
    let shi = 巳;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 辛丑;

    let sk = SiKe::new(&tp, &gz);
    let sc = get反呤(&tp, &sk);
    assert_eq!(
        "亥",
        sc[0].to_string(),
        "{}日，{}将{}，初传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "未",
        sc[1].to_string(),
        "{}日，{}将{}，中传`未`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "辰",
        sc[2].to_string(),
        "{}日，{}将{}，末传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_get八专() {
    // t.Log("测试八专课")

    // t.Log("丁未日，亥将申时，阴日八专课")
    let yue_jing = 亥;
    let shi = 申;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 丁未;

    let sk = SiKe::new(&tp, &gz);
    let sc = get八专(&sk);
    assert_eq!(
        "亥",
        sc[0].to_string(),
        "{}日，{}将{}，初传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "戌",
        sc[1].to_string(),
        "{}日，{}将{}，中传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "戌",
        sc[2].to_string(),
        "{}日，{}将{}，末传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("甲寅日，丑将辰时，阳日八专")
    let yue_jing = 丑;
    let shi = 辰;
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = 甲寅;

    let sk = SiKe::new(&tp, &gz);
    let sc = get八专(&sk);
    assert_eq!(
        "丑",
        sc[0].to_string(),
        "{}日，{}将{}，初传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "亥",
        sc[1].to_string(),
        "{}日，{}将{}，中传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "亥",
        sc[2].to_string(),
        "{}日，{}将{}，末传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("己未日，酉将未时，独足格")
    let g = TianGan::default().plus(5);
    let z = DiZhi::default().plus(7);
    let yue_jing = DiZhi::default().plus(9);
    let shi = DiZhi::default().plus(7);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get八专(&sk);
    assert_eq!(
        "酉",
        sc[0].to_string(),
        "{}日，{}将{}，初传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "酉",
        sc[1].to_string(),
        "{}日，{}将{}，中传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "酉",
        sc[2].to_string(),
        "{}日，{}将{}，末传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_get遥克() {
    // t.Log("测试遥克课")

    // t.Log("壬辰日，巳将寅时，比用遥克课")
    let g = TianGan::default().plus(-2);
    let z = DiZhi::default().plus(4);
    let yue_jing = DiZhi::default().plus(5);
    let shi = DiZhi::default().plus(2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get遥克(&tp, &sk);
    assert_eq!(
        "戌",
        sc[0].to_string(),
        "{}日，{}将{}，初传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "丑",
        sc[1].to_string(),
        "{}日，{}将{}，中传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "辰",
        sc[2].to_string(),
        "{}日，{}将{}，末传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("壬申日，亥将申时，干遥克课")
    let g = TianGan::default().plus(-2);
    let z = DiZhi::default().plus(8);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(8);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get遥克(&tp, &sk);
    assert_eq!(
        "巳",
        sc[0].to_string(),
        "{}日，{}将{}，初传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "申",
        sc[1].to_string(),
        "{}日，{}将{}，中传`申`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "亥",
        sc[2].to_string(),
        "{}日，{}将{}，末传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_get昴星() {
    // t.Log("测试昴星课")

    // t.Log("戊寅日，辰将子时，阳昴星课")
    let g = TianGan::default().plus(4);
    let z = DiZhi::default().plus(2);
    let yue_jing = DiZhi::default().plus(4);
    let shi = DiZhi::default().plus(0);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get昴星(&tp, &sk);
    assert_eq!(
        "丑",
        sc[0].to_string(),
        "{}日，{}将{}，初传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "午",
        sc[1].to_string(),
        "{}日，{}将{}，中传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "酉",
        sc[2].to_string(),
        "{}日，{}将{}，末传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("丁亥日，巳将寅时，阴日昴星")
    let g = TianGan::default().plus(3);
    let z = DiZhi::default().plus(-1);
    let yue_jing = DiZhi::default().plus(5);
    let shi = DiZhi::default().plus(2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get昴星(&tp, &sk);
    assert_eq!(
        "午",
        sc[0].to_string(),
        "{}日，{}将{}，初传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "戌",
        sc[1].to_string(),
        "{}日，{}将{}，中传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "寅",
        sc[2].to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_get别责() {
    // t.Log("测试别责课")

    // t.Log("戊辰日，亥将戌时，阳别责课")
    let g = TianGan::default().plus(4);
    let z = DiZhi::default().plus(4);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(-2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get别责(&tp, &sk);
    assert_eq!(
        "寅",
        sc[0].to_string(),
        "{}日，{}将{}，初传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "午",
        sc[1].to_string(),
        "{}日，{}将{}，中传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "午",
        sc[2].to_string(),
        "{}日，{}将{}，末传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("辛酉日，子将丑时，阴别责")
    let g = TianGan::default().plus(7);
    let z = DiZhi::default().plus(-3);
    let yue_jing = DiZhi::default().plus(0);
    let shi = DiZhi::default().plus(1);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = get别责(&tp, &sk);
    assert_eq!(
        "子",
        sc[0].to_string(),
        "{}日，{}将{}，初传`子`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "酉",
        sc[1].to_string(),
        "{}日，{}将{}，中传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "酉",
        sc[2].to_string(),
        "{}日，{}将{}，末传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_san_chuan() {
    // t.Log("测试获取三传函数：sanChuan()")

    // t.Log("下克上")
    // t.Log("丙戌日，申将巳时")
    let g = TianGan::default().plus(2);
    let z = DiZhi::default().plus(-2);
    let yue_jing = DiZhi::default().plus(8);
    let shi = DiZhi::default().plus(5);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "申",
        sc[0].to_string(),
        "{}日，{}将{}，初传`申`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "亥",
        sc[1].to_string(),
        "{}日，{}将{}，中传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "寅",
        sc[2].to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // 	t.Log("上克下")
    // t.Log("丁丑日，申将子时")
    let g = TianGan::default().plus(3);
    let z = DiZhi::default().plus(1);
    let yue_jing = DiZhi::default().plus(8);
    let shi = DiZhi::default().plus(0);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "巳",
        sc[0].to_string(),
        "{}日，{}将{}，初传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "丑",
        sc[1].to_string(),
        "{}日，{}将{}，中传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "酉",
        sc[2].to_string(),
        "{}日，{}将{}，末传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // 	t.Log("比用")
    // 	t.Log("甲戌日，亥将戌时")
    let g = TianGan::default();
    let z = DiZhi::default().plus(-2);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(-2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "辰",
        sc[0].to_string(),
        "{}日，{}将{}，初传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "巳",
        sc[1].to_string(),
        "{}日，{}将{}，中传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "午",
        sc[2].to_string(),
        "{}日，{}将{}，末传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
    // 	t.Log("庚午日，未将寅时，比用")
    let g = TianGan::default().plus(6);
    let z = DiZhi::default().plus(6);
    let yue_jing = DiZhi::default().plus(7);
    let shi = DiZhi::default().plus(2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "辰",
        sc[0].to_string(),
        "{}日，{}将{}，初传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "酉",
        sc[1].to_string(),
        "{}日，{}将{}，中传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "寅",
        sc[2].to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // 	t.Log("涉害")
    // t.Log("辛卯日，亥将未时")
    let g = TianGan::default().plus(7);
    let z = DiZhi::default().plus(3);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(7);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "亥",
        sc[0].to_string(),
        "{}日，{}将{}，初传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "卯",
        sc[1].to_string(),
        "{}日，{}将{}，中传`卯`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "未",
        sc[2].to_string(),
        "{}日，{}将{}，末传`未`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("遥克")
    // t.Log("戊辰日，亥将未时")
    let g = TianGan::default().plus(4);
    let z = DiZhi::default().plus(4);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(7);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "子",
        sc[0].to_string(),
        "{}日，{}将{}，初传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "辰",
        sc[1].to_string(),
        "{}日，{}将{}，中传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "申",
        sc[2].to_string(),
        "{}日，{}将{}，末传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // 	t.Log("昴星")
    // 	t.Log("乙未日，亥将子时")
    let g = TianGan::default().plus(1);
    let z = DiZhi::default().plus(7);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(0);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "戌",
        sc[0].to_string(),
        "{}日，{}将{}，初传`戌`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "卯",
        sc[1].to_string(),
        "{}日，{}将{}，中传`卯`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "午",
        sc[2].to_string(),
        "{}日，{}将{}，末传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // 	t.Log("伏吟")
    // 	t.Log("丙子日，亥将亥时")
    let g = TianGan::default().plus(2);
    let z = DiZhi::default();
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(-1);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "巳",
        sc[0].to_string(),
        "{}日，{}将{}，初传`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "申",
        sc[1].to_string(),
        "{}日，{}将{}，中传`申`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "寅",
        sc[2].to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // 	t.Log("反吟")
    // 	t.Log("丁丑日，亥将巳时")
    let g = TianGan::default().plus(3);
    let z = DiZhi::default().plus(1);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(5);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "亥",
        sc[0].to_string(),
        "{}日，{}将{}，初传`亥`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "未",
        sc[1].to_string(),
        "{}日，{}将{}，中传`未`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "丑",
        sc[2].to_string(),
        "{}日，{}将{}，末传`丑`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );

    // t.Log("别责")
    // t.Log("戊辰日，亥将戌时")
    let g = TianGan::default().plus(4);
    let z = DiZhi::default().plus(4);
    let yue_jing = DiZhi::default().plus(-1);
    let shi = DiZhi::default().plus(-2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = san_chuan(&tp, &sk);
    assert_eq!(
        "寅",
        sc[0].to_string(),
        "{}日，{}将{}，初传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc[0]
    );
    assert_eq!(
        "午",
        sc[1].to_string(),
        "{}日，{}将{}，中传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[1]
    );
    assert_eq!(
        "午",
        sc[2].to_string(),
        "{}日，{}将{}，末传`午`，非{}",
        gz,
        yue_jing,
        shi,
        sc[2]
    );
}

#[test]
fn test_new_san_chuan() {
    // t.Log("测试方法：NewSanChuan")

    // t.Log("辛丑日，未将寅时")
    let g = TianGan::default().plus(7);
    let z = DiZhi::default().plus(-3);
    let yue_jing = DiZhi::default().plus(7);
    let shi = DiZhi::default().plus(2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = SanChuan::new(&tp, &sk);

    assert_eq!(
        "未",
        sc.chu.to_string(),
        "{}日，{}将{}，初传`未`，非{}",
        gz,
        yue_jing,
        shi,
        sc.chu
    );
    assert_eq!(
        "子",
        sc.zhong.to_string(),
        "{}日，{}将{}，中传`子`，非{}",
        gz,
        yue_jing,
        shi,
        sc.zhong
    );
    assert_eq!(
        "巳",
        sc.mo.to_string(),
        "{}日，{}将{}，末传遁干`巳`，非{}",
        gz,
        yue_jing,
        shi,
        sc.mo
    );

    assert_eq!(
        "己", sc.dun_gan[0],
        "{}日，{}将{}，初传遁干`己`，非{}",
        gz, yue_jing, shi, sc.dun_gan[0]
    );
    assert_eq!(
        "", sc.dun_gan[1],
        "{}日，{}将{}，中传遁干`无遁干`，非{}",
        gz, yue_jing, shi, sc.dun_gan[1]
    );
    assert_eq!(
        "丁", sc.dun_gan[2],
        "{}日，{}将{}，末传遁干`丁`，非{}",
        gz, yue_jing, shi, sc.dun_gan[2]
    );

    assert_eq!(
        "父", sc.liu_qing[0],
        "{}日，{}将{}，初传六亲`父`，非{}",
        gz, yue_jing, shi, sc.liu_qing[0]
    );
    assert_eq!(
        "子", sc.liu_qing[1],
        "{}日，{}将{}，中传六亲`子`，非{}",
        gz, yue_jing, shi, sc.liu_qing[1]
    );
    assert_eq!(
        "官", sc.liu_qing[2],
        "{}日，{}将{}，末传六亲`官`，非{}",
        gz, yue_jing, shi, sc.liu_qing[2]
    );

    let sc_json = serde_json::to_string(&sc);
    assert!(sc_json.is_ok());
    let sc_json = sc_json.unwrap();

    let m = SanChuanJson {
        chu: "未",
        zhong: "子",
        mo: "巳",
        dun_gan: ["己", "", "丁"],
        liu_qing: ["父", "子", "官"],
    };

    let s = serde_json::to_string(&m).unwrap();

    assert_eq!(s, sc_json, "三传序列化失败");

    // t.Log("庚午日，未将寅时")
    let g = TianGan::default().plus(6);
    let z = DiZhi::default().plus(6);
    let yue_jing = DiZhi::default().plus(7);
    let shi = DiZhi::default().plus(2);
    let tp = TianPan::new(yue_jing.clone(), shi.clone());
    let gz = GanZhi::new(&g, &z).unwrap();

    let sk = SiKe::new(&tp, &gz);
    let sc = SanChuan::new(&tp, &sk);

    assert_eq!(
        "辰",
        sc.chu.to_string(),
        "{}日，{}将{}，初传`辰`，非{}",
        gz,
        yue_jing,
        shi,
        sc.chu
    );
    assert_eq!(
        "酉",
        sc.zhong.to_string(),
        "{}日，{}将{}，中传`酉`，非{}",
        gz,
        yue_jing,
        shi,
        sc.zhong
    );
    assert_eq!(
        "寅",
        sc.mo.to_string(),
        "{}日，{}将{}，末传`寅`，非{}",
        gz,
        yue_jing,
        shi,
        sc.mo
    );

    assert_eq!(
        "戊", sc.dun_gan[0],
        "{}日，{}将{}，初传遁干`戊`，非{}",
        gz, yue_jing, shi, sc.dun_gan[0]
    );
    assert_eq!(
        "癸", sc.dun_gan[1],
        "{}日，{}将{}，中传遁干`癸`，非{}",
        gz, yue_jing, shi, sc.dun_gan[1]
    );
    assert_eq!(
        "丙", sc.dun_gan[2],
        "{}日，{}将{}，末传遁干`丙`，非{}",
        gz, yue_jing, shi, sc.dun_gan[2]
    );

    assert_eq!(
        "父", sc.liu_qing[0],
        "{}日，{}将{}，初传六亲`父`，非{}",
        gz, yue_jing, shi, sc.liu_qing[0]
    );
    assert_eq!(
        "兄", sc.liu_qing[1],
        "{}日，{}将{}，中传六亲`兄`，非{}",
        gz, yue_jing, shi, sc.liu_qing[1]
    );
    assert_eq!(
        "财", sc.liu_qing[2],
        "{}日，{}将{}，末传六亲`财`，非{}",
        gz, yue_jing, shi, sc.liu_qing[2]
    );

    let sc_json = serde_json::to_string(&sc);
    assert!(sc_json.is_ok());
    let sc_json = sc_json.unwrap();

    let m = SanChuanJson {
        chu: "辰",
        zhong: "酉",
        mo: "寅",
        dun_gan: ["戊", "癸", "丙"],
        liu_qing: ["父", "兄", "财"],
    };
    let s = serde_json::to_string(&m).unwrap();

    assert_eq!(s, sc_json, "三传序列化失败");
}

#[derive(Serialize)]
struct SanChuanJson<'a> {
    chu: &'a str,
    zhong: &'a str,
    mo: &'a str,
    dun_gan: [&'a str; 3],
    liu_qing: [&'a str; 3],
}
