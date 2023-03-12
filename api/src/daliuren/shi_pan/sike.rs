use ganzhiwuxin::{DiZhi, GanZhi, TianGan};
use serde::{ser::SerializeStruct, Serialize};

use super::{jigong::ji_gong, tianpan::TianPan};

#[derive(Debug)]
pub struct SiKe {
    /*
        gan 日干
        zhi 日支
        ganYang 干阳
        ganYing 干阴
        zhiYang 支阳
        zhiYing 支阴
    */
    pub gan: TianGan,
    pub zhi: DiZhi,
    pub gan_yang: DiZhi,
    pub gan_ying: DiZhi,
    pub zhi_yang: DiZhi,
    pub zhi_ying: DiZhi,
}

impl SiKe {
    pub fn new(tian_pan: &TianPan, divination_day: &GanZhi) -> SiKe {
        let gan = divination_day.gan.clone();
        let gan_yang = tian_pan.up(&ji_gong(&gan));
        let gan_ying = tian_pan.up(&gan_yang);

        let zhi = divination_day.zhi.clone();
        let zhi_yang = tian_pan.up(&zhi);
        let zhi_ying = tian_pan.up(&zhi_yang);
        Self {
            gan,
            zhi,
            gan_yang,
            gan_ying,
            zhi_yang,
            zhi_ying,
        }
    }
}

impl Serialize for SiKe {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let mut state = serializer.serialize_struct("SiKe", 6)?;
        state.serialize_field("gan", &self.gan.to_string())?;
        state.serialize_field("gan_yang", &self.gan_yang.to_string())?;
        state.serialize_field("gan_ying", &self.gan_ying.to_string())?;

        state.serialize_field("zhi", &self.zhi.to_string())?;
        state.serialize_field("zhi_yang", &self.zhi_yang.to_string())?;
        state.serialize_field("zhi_ying", &self.zhi_ying.to_string())?;
        state.end()
    }
}

#[cfg(test)]
mod tests {
    use crate::daliuren::shi_pan::tianpan::TianPan;

    use super::SiKe;
    use ganzhiwuxin::{DiZhi, GanZhi, TianGan};

    #[test]
    fn test_new_si_ke() {
        // 测试NewSiKe()
        let 子 = DiZhi::new("子").unwrap();
        let 甲 = TianGan::new("甲").unwrap();
        let 丑 = 子.plus(1);
        let 甲子 = GanZhi::new(甲, 子.clone()).unwrap();

        let tp = TianPan::new(&子, &丑);
        // 甲子日，子将丑
        let sike = SiKe::new(&tp, &甲子);
        assert_eq!(
            sike.gan.to_string(),
            "甲",
            "甲子日，子将丑时，干是：`甲`，非是`{}`",
            sike.gan
        );

        assert_eq!(
            sike.gan_yang.to_string(),
            "丑",
            "甲子日，子将丑时，干阳是：`丑`，非是`{}`",
            sike.gan_yang
        );

        assert_eq!(
            sike.gan_ying.to_string(),
            "子",
            "甲子日，子将丑时，干阴是：`子`，非是`{}`",
            sike.gan_ying
        );

        assert_eq!(
            sike.zhi.to_string(),
            "子",
            "甲子日，子将丑时，支是：`子`，非是`{}`",
            sike.zhi
        );

        assert_eq!(
            sike.zhi_yang.to_string(),
            "亥",
            "甲子日，子将丑时，支阳是：`亥`，非是`{}`",
            sike.zhi_yang
        );

        assert_eq!(
            sike.zhi_ying.to_string(),
            "戌",
            "甲子日，子将丑时，支阴是：`戌`，非是`{}`",
            sike.zhi_ying
        );

        let json = serde_json::to_string(&sike);
        assert!(json.is_ok());
        let json = json.unwrap();
        // let m = HashMap::from([
        //     ("gan", "甲"),
        //     ("gan_yang", "丑"),
        //     ("gan_ying", "子"),
        //     ("zhi", "子"),
        //     ("zhi_yang", "亥"),
        //     ("zhi_ying", "戌"),
        // ]);
        // let s = serde_json::to_string(&m).unwrap();
        let s=   "{\"gan\":\"甲\",\"gan_yang\":\"丑\",\"gan_ying\":\"子\",\"zhi\":\"子\",\"zhi_yang\":\"亥\",\"zhi_ying\":\"戌\"}";
        assert_eq!(json, s, "四课序列化失败");
    }
}
