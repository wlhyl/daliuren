#[cfg(test)]
mod test;
use std::collections::HashMap;

use super::{
    jigong::{ji_gong, ji_gong_gan},
    sike::SiKe,
    tianpan::TianPan,
};
use ganzhiwuxing::{
    DiZhi::{self, *},
    TianGan::*,
};
use itertools::Itertools;
use serde::{ser::SerializeStruct, Serialize};

/// 三传
#[derive(Debug)]
pub struct SanChuan {
    chu: DiZhi,
    zhong: DiZhi,
    mo: DiZhi,
    dun_gan: [String; 3],
    liu_qing: [String; 3],
}

impl SanChuan {
    pub fn new(tian_pan: &TianPan, sike: &SiKe) -> Self {
        // 三传
        let sc = san_chuan(tian_pan, sike);
        let (chu, zhong, mo) = (sc[0].clone(), sc[1].clone(), sc[2].clone());

        // 遁干
        let delta = sike.gan.minus(&甲);

        let xun_shou = sike.zhi.plus(-1 * isize::from(delta));

        //         let dunGan :[DiZhi;3]= [chu, zhong, mo].iter().map(|d|{
        // let zhiDelta = d.minus(&xunShou);
        // if zhiDelta == 10 || zhiDelta == 11 {
        //     "".to_owned()
        // } else {
        //     甲.plus(zhiDelta.into()).to_string()
        // }
        // }).collect();

        let mut dun_gan: [String; 3] = Default::default();
        let mut liu_qing: [String; 3] = Default::default();
        for (index, d) in [&chu, &zhong, &mo].iter().enumerate() {
            let zhi_delta = d.minus(&xun_shou);
            if zhi_delta == 10 || zhi_delta == 11 {
                dun_gan[index] = "".to_string();
            } else {
                dun_gan[index] = 甲.plus(zhi_delta.into()).to_string();
            }

            // 六亲
            if sike.gan.wu_xing().ke(&d.wu_xing()) {
                liu_qing[index] = "财".to_owned();
            } else if d.wu_xing().ke(&sike.gan.wu_xing()) {
                liu_qing[index] = "官".to_owned();
            } else if sike.gan.wu_xing().sheng(&d.wu_xing()) {
                liu_qing[index] = "子".to_owned();
            } else if d.wu_xing().sheng(&sike.gan.wu_xing()) {
                liu_qing[index] = "父".to_owned()
            } else {
                liu_qing[index] = "兄".to_owned()
            }
        }

        Self {
            chu,
            zhong,
            mo,
            dun_gan,
            liu_qing,
        }
    }
}

impl Serialize for SanChuan {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        // 5 is the number of fields in the struct.
        let mut state = serializer.serialize_struct("SanChuan", 5)?;
        state.serialize_field("chu", &self.chu.to_string())?;
        state.serialize_field("zhong", &self.zhong.to_string())?;
        state.serialize_field("mo", &self.mo.to_string())?;
        state.serialize_field("dun_gan", &self.dun_gan)?;
        state.serialize_field("liu_qing", &self.liu_qing)?;
        state.end()
    }
}

fn san_chuan(tianpan: &TianPan, sike: &SiKe) -> Vec<DiZhi> {
    if sike.zhi_yang == sike.zhi {
        return get伏呤(sike);
    }

    if sike.zhi_yang.liu_chong(&sike.zhi) {
        return get反呤(tianpan, sike);
    }
    let sc = get贼克(tianpan, sike);
    if !sc.is_empty() {
        return sc;
    }

    if sike.gan_yang == sike.zhi_yang {
        return get八专(sike);
    }

    let sc = get遥克(tianpan, sike);
    if !sc.is_empty() {
        return sc;
    }

    //确认是昴星课
    let 四课上神组 = [
        sike.gan_yang.clone(),
        sike.gan_ying.clone(),
        sike.zhi_yang.clone(),
        sike.zhi_ying.clone(),
    ];
    let 无重复的课: Vec<_> = 四课上神组.iter().unique().collect();

    if 无重复的课.len() == 4 {
        return get昴星(tianpan, sike);
    }
    get别责(tianpan, sike)
}

//有下克上的课
// 返回有贼的课序数,1,2,3,4课
fn has贼(sike: &SiKe) -> Vec<u8> {
    let mut ke_list = Vec::new();

    if sike.gan.wu_xing().ke(&sike.gan_yang.wu_xing()) {
        ke_list.push(1);
    }
    if sike.gan_yang.ke(&sike.gan_ying) {
        ke_list.push(2);
    }
    if sike.zhi.ke(&sike.zhi_yang) {
        ke_list.push(3);
    }
    if sike.zhi_yang.ke(&sike.zhi_ying) {
        ke_list.push(4);
    }

    if ke_list.is_empty() {
        return ke_list;
    }

    //删除重复课
    let mut shang_shen_list = Vec::new();
    let mut ke_list_tmp = Vec::new();
    for v in ke_list {
        let shang_shen = get_shang_shen(sike, v);

        if !shang_shen_list.contains(shang_shen) {
            shang_shen_list.push(shang_shen.clone());
            ke_list_tmp.push(v);
        }
    }

    ke_list_tmp
}

//有上克下的课
fn has克(sike: &SiKe) -> Vec<u8> {
    let mut ke_list = vec![];

    if sike.gan_yang.wu_xing().ke(&sike.gan.wu_xing()) {
        ke_list.push(1);
    }
    if sike.gan_ying.ke(&sike.gan_yang) {
        ke_list.push(2);
    }
    if sike.zhi_yang.ke(&sike.zhi) {
        ke_list.push(3);
    }
    if sike.zhi_ying.ke(&sike.zhi_yang) {
        ke_list.push(4);
    }

    if ke_list.is_empty() {
        return ke_list;
    }

    //删除重复课
    let mut shang_shen_list = vec![];
    let mut ke_list_tmp = vec![];
    for v in ke_list {
        let shang_shen = get_shang_shen(sike, v);

        if !shang_shen_list.contains(shang_shen) {
            shang_shen_list.push(shang_shen.clone());
            ke_list_tmp.push(v);
        }
    }
    ke_list_tmp
}

fn get贼克(tian_pan: &TianPan, sike: &SiKe) -> Vec<DiZhi> {
    let 贼 = has贼(sike);
    let 克 = has克(sike);

    if 贼.len() == 1 {
        let chu = get_shang_shen(sike, 贼[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);
        return vec![chu.clone(), zhong, mo];
    }

    if 贼.len() > 1 {
        return get比用(tian_pan, sike, &贼);
    }

    if 克.len() == 1 {
        let chu = get_shang_shen(sike, 克[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);
        return vec![chu.clone(), zhong, mo];
    }
    if 克.len() > 1 {
        return get比用(tian_pan, sike, &克);
    }

    vec![]
}

fn get比用(tian_pan: &TianPan, sk: &SiKe, ke_list: &[u8]) -> Vec<DiZhi> {
    /* len(results) ==1, >1, ==0
       results中保存 取比用後的结果
    */
    let mut results = vec![];
    for it in ke_list {
        if get_shang_shen(sk, *it).masculine() == sk.gan.masculine() {
            results.push(*it);
        }
    }

    if results.len() == 1 {
        let chu = get_shang_shen(sk, results[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);
        return vec![chu.clone(), zhong, mo];
    }
    // 俱不比
    if results.is_empty() {
        return get涉害(tian_pan, sk, ke_list);
    }
    //多个俱比 len(result) >1
    return get涉害(tian_pan, sk, &results);
}

fn get涉害(tian_pan: &TianPan, sike: &SiKe, ke_list: &[u8]) -> Vec<DiZhi> {
    let is贼 = |it| {
        if it == 1 {
            return sike.gan.wu_xing().ke(&sike.gan_yang.wu_xing());
        }
        if it == 2 {
            return sike.gan_yang.ke(&sike.gan_ying);
        }
        if it == 3 {
            return sike.zhi.ke(&sike.zhi_yang);
        }
        sike.zhi_yang.ke(&sike.zhi_ying)
    };

    //使用涉归，不用孟仲季取法
    // __课的涉害深度中的key存储地支的num，value存储涉害深度
    let mut __课的涉害深度 = HashMap::new();
    for v in ke_list {
        let 课 = get_shang_shen(sike, *v);
        let 临地盘 = tian_pan.down(课);

        let mut count = 0u8;
        for i in 0..课.minus(&临地盘) {
            let __d = 临地盘.plus(i.into());
            let ji_gan = ji_gong_gan(&__d);
            if is贼(*v) {
                if __d.ke(课) {
                    count += 1;
                }
                for g in ji_gan {
                    if g.wu_xing().ke(&课.wu_xing()) {
                        count += 1;
                    }
                }
            } else {
                if 课.ke(&__d) {
                    count += 1;
                }
                for g in ji_gan {
                    if 课.wu_xing().ke(&g.wu_xing()) {
                        count += 2;
                    }
                }
            }
        }
        __课的涉害深度.insert(*v, count);
    }

    let 最大涉害深度 = __课的涉害深度
        .iter()
        .map(|(_, &v)| v)
        .max()
        .unwrap_or_default();

    let 有最大涉害深度的支组: Vec<_> = __课的涉害深度
        .iter()
        .filter(|(_, &v)| v == 最大涉害深度)
        .map(|(&k, _)| k)
        .collect();

    if 有最大涉害深度的支组.len() == 1 {
        let chu = get_shang_shen(sike, 有最大涉害深度的支组[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);

        return vec![chu.clone(), zhong, mo];
    }

    //涉害深度相同
    // 找出四孟
    let 四孟支组: Vec<_> = 有最大涉害深度的支组
        .iter()
        .filter(|&&v| {
            let 临地盘 = tian_pan.down(get_shang_shen(sike, v));
            临地盘.minus(&寅) % 3 == 0
        })
        .collect();

    if 四孟支组.len() == 1 {
        let chu = get_shang_shen(sike, *四孟支组[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);

        return vec![chu.clone(), zhong, mo];
    }

    if 四孟支组.len() > 0 {
        if sike.gan.masculine() {
            let chu = sike.gan_yang.clone();
            let zhong = tian_pan.up(&chu);
            let mo = tian_pan.up(&zhong);

            return vec![chu, zhong, mo];
        } else {
            let chu = sike.zhi_yang.clone();
            let zhong = tian_pan.up(&chu);
            let mo = tian_pan.up(&zhong);
            return vec![chu, zhong, mo];
        }
    }

    // 从仲发用
    let 四仲支组: Vec<_> = 有最大涉害深度的支组
        .iter()
        .filter(|&&v| {
            let 临地盘 = tian_pan.down(get_shang_shen(sike, v));
            临地盘.minus(&子) % 3 == 0
        })
        .collect();

    if 四仲支组.len() == 1 {
        let chu = get_shang_shen(sike, *四仲支组[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);

        return vec![chu.clone(), zhong, mo];
    }
    // 剩余情况有两种:
    // len(四仲支组) > 1
    // len(四仲支组) == 0，此种情况，支组临于四季，与前一种情况一并按“复等课”取三传
    if sike.gan.masculine() {
        let chu = sike.gan_yang.clone();
        let zhong = tian_pan.up(&chu);
        let mo = tian_pan.up(&zhong);

        return vec![chu, zhong, mo];
    } else {
        let chu = sike.zhi_yang.clone();
        let zhong = tian_pan.up(&chu);
        let mo = tian_pan.up(&zhong);
        return vec![chu, zhong, mo];
    }
}

fn get遥克(tian_pan: &TianPan, sike: &SiKe) -> Vec<DiZhi> {
    let mut ke = vec![];
    if sike.gan_ying.wu_xing().ke(&sike.gan.wu_xing()) {
        ke.push(2);
    }
    if sike.zhi_yang.wu_xing().ke(&sike.gan.wu_xing()) {
        ke.push(3);
    }
    if sike.zhi_ying.wu_xing().ke(&sike.gan.wu_xing()) {
        ke.push(4);
    }

    if ke.len() == 0 {
        if sike.gan.wu_xing().ke(&sike.gan_ying.wu_xing()) {
            ke.push(2);
        }
        if sike.gan.wu_xing().ke(&sike.zhi_yang.wu_xing()) {
            ke.push(3);
        }
        if sike.gan.wu_xing().ke(&sike.zhi_ying.wu_xing()) {
            ke.push(4);
        }
    }
    if ke.len() == 0 {
        return vec![];
    }
    // 去重
    let mut ke_list = vec![];
    for v in ke {
        let k0 = get_shang_shen(sike, v);
        let has = ke_list.iter().find(|&&v| {
            let k1 = get_shang_shen(sike, v);
            k0 == k1
        });
        if has.is_none() {
            ke_list.push(v);
        }
    }

    if ke_list.len() == 1 {
        let chu = get_shang_shen(sike, ke_list[0]);
        let zhong = tian_pan.up(chu);
        let mo = tian_pan.up(&zhong);
        return vec![chu.clone(), zhong, mo];
    }
    // 如果有两课遥克，取比用，遥克课取比用之後，不会出现涉害（需要验证）
    get比用(tian_pan, sike, &ke_list)
}

// 不判断是否为昴星课，只按昴星课取三传
fn get昴星(tian_pan: &TianPan, sike: &SiKe) -> Vec<DiZhi> {
    if sike.gan.masculine() {
        let chu = tian_pan.up(&酉);
        let zhong = sike.zhi_yang.clone();
        let mo = sike.gan_yang.clone();
        return vec![chu, zhong, mo];
    } else {
        let chu = tian_pan.down(&酉);
        let zhong = sike.gan_yang.clone();
        let mo = sike.zhi_yang.clone();
        return vec![chu, zhong, mo];
    }
}

// 不判断是否为别责课，只按别责取传
fn get别责(tianpan: &TianPan, sike: &SiKe) -> Vec<DiZhi> {
    let d = if sike.gan.masculine() {
        ji_gong(&sike.gan.plus(5))
    } else {
        //干为阴的情况
        sike.zhi.plus(4)
    };
    //中末皆取干上神
    // 阴日，取三合上神为初传，不取三合为初传，《六壬直指》原文的括号中的注解似乎有错
    let chu = tianpan.up(&d);
    let zhong = sike.gan_yang.clone();
    let mo = sike.gan_yang.clone();
    return vec![chu, zhong, mo];
}

// 调用此函数前，需要先确定是八专课，此函数只按八专取三传，不判断是否为八专课
fn get八专(sike: &SiKe) -> Vec<DiZhi> {
    if sike.gan.masculine() {
        let chu = sike.gan_yang.plus(2);
        vec![chu, sike.gan_yang.clone(), sike.gan_yang.clone()]
    } else {
        let chu = sike.zhi_ying.plus(-2);
        vec![chu, sike.gan_yang.clone(), sike.gan_yang.clone()]
    }
}

fn get伏呤(sike: &SiKe) -> Vec<DiZhi> {
    // 六乙、六癸日 无克，阳日取干上神发用
    let chu = if sike.gan == 乙 || sike.gan == 癸 || sike.gan.masculine() {
        sike.gan_yang.clone()
    } else {
        // 阴日，非六乙日、六癸
        sike.zhi_yang.clone()
    };

    //求得中传
    let mut zhong = Default::default();
    for i in 0..12 {
        let d = chu.plus(i);
        if chu.xing(&d) {
            zhong = d;
            break;
        }
    }
    // 初为自刑，阳日、六乙日、六癸日取支上神为中传
    if chu == zhong {
        if sike.gan == 乙 || sike.gan == 癸 || sike.gan.masculine() {
            //六乙、六癸、阳日，初传传自刑，取支上神为中传
            zhong = sike.zhi_yang.clone()
        } else {
            zhong = sike.gan_yang.clone()
        }
    }

    //求末传
    let mut mo = Default::default();
    for i in 0..12 {
        let d = chu.plus(i);
        if zhong.xing(&d) {
            mo = d;
            break;
        }
    }

    //中传自刑，取中所冲之神
    if zhong == mo {
        mo = zhong.plus(6);
    }

    // 初、中互刑，如：子、卯，末取中所冲之神
    if zhong.xing(&chu) {
        mo = zhong.plus(6)
    }
    vec![chu, zhong, mo]
}

fn get反呤(tian_pan: &TianPan, sike: &SiKe) -> Vec<DiZhi> {
    let sc = get贼克(tian_pan, sike);
    if !sc.is_empty() {
        return sc;
    }

    // 驿马计算
    let mut zhi = sike.zhi.clone();
    // 根据日支找到三合日支的四孟地支
    // 驿马 = 四孟地支 + 6
    let mut yi_ma = Default::default();
    for _ in 0..3 {
        if zhi.minus(&寅) % 3 == 0 {
            yi_ma = zhi.plus(6);
            break;
        }
        // 下一个三合支 = 地支 + 4
        zhi = zhi.plus(4);
    }
    let chu = yi_ma;
    let zhong = sike.zhi_yang.clone();
    let mo = sike.gan_yang.clone();
    vec![chu, zhong, mo]
}

// 获得某课的上神
// 第一课:1,第二课：2，第三课：3，第四课：4
fn get_shang_shen(sk: &SiKe, n: u8) -> &DiZhi {
    if n == 1 {
        return &sk.gan_yang;
    }
    if n == 2 {
        return &sk.gan_ying;
    }
    if n == 3 {
        return &sk.zhi_yang;
    }

    return &sk.zhi_ying;
}
