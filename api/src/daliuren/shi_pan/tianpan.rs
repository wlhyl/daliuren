use ganzhiwuxin::DiZhi;
use itertools::Itertools;
use serde::{ser::SerializeSeq, Serialize};

#[derive(Debug)]
pub struct TianPan {
    pub yue_jiang: DiZhi,
    pub divination_time: DiZhi,
}

impl TianPan {
    pub fn new(yue_jiang: &DiZhi, divination_time: &DiZhi) -> Self {
        Self {
            yue_jiang: yue_jiang.clone(),
            divination_time: divination_time.clone(),
        }
    }
    //取得地盘上神
    pub fn up(&self, d: &DiZhi) -> DiZhi {
        self.yue_jiang.plus(d.minus(&self.divination_time) as isize)
    }

    //取得天神所临地盘
    pub fn down(&self, d: &DiZhi) -> DiZhi {
        self.divination_time.plus(d.minus(&self.yue_jiang) as isize)
    }
}

impl Serialize for TianPan {
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

#[cfg(test)]
mod tests {
    use ganzhiwuxin::DiZhi;

    use super::TianPan;

    #[test]
    fn test_new() {
        let 子 = DiZhi::new("子").unwrap();
        let 丑 = 子.plus(1);
        let tp = TianPan::new(&子, &丑);
        assert_eq!(tp.yue_jiang, 子);
        assert_eq!(tp.divination_time, 丑);
    }

    #[test]
    fn test_tianpan() {
        // 测试天盘
        let yue_jiang = DiZhi::new("申").unwrap();

        let divination_time = DiZhi::new("辰").unwrap();

        let zi = DiZhi::new("子").unwrap();

        let tp = TianPan {
            yue_jiang: yue_jiang.clone(),
            divination_time: divination_time.clone(),
        };
        for i in 0..12 {
            let d = zi.plus(i);
            let up = tp.up(&d);
            assert_eq!(
                up,
                zi.plus(4 + i),
                "{}加{}，{}上神是{}，而非{}",
                yue_jiang,
                divination_time,
                d,
                zi.plus(4 + i),
                up
            );

            let down = tp.down(&d);
            assert_eq!(
                down,
                zi.plus(8 + i),
                "{}加{}，{}临{}，而非{}",
                yue_jiang,
                divination_time,
                d,
                zi.plus(8 + i),
                down
            );
        }

        let json = serde_json::to_string(&tp);
        assert!(json.is_ok());
        let json = json.unwrap();
        let s =
            "[\"辰\",\"巳\",\"午\",\"未\",\"申\",\"酉\",\"戌\",\"亥\",\"子\",\"丑\",\"寅\",\"卯\"]";
        assert_eq!(json, s, "天盘序列化失败");
    }
}