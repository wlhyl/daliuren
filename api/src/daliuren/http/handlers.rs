use crate::{daliuren::shi_pan::DaliurenShiPan, state::AppState};
use actix_web::{post, web, HttpResponse, Responder};
use ganzhiwuxing::DiZhi::*;

use super::request::DaLiuRenReust;

/// 大六壬排盘
///
/// method: post
///
/// api: /daliuren
#[post("daliuren")]
pub async fn daliuren(
    app_state: web::Data<AppState>,
    r: actix_web_validator::Json<DaLiuRenReust>,
) -> impl Responder {
    let r = r.into_inner();
    let yue_jiang = if let Some(s) = r.yue_jiang {
        match s.as_str() {
            "子" => Some(子),
            "丑" => Some(丑),
            "寅" => Some(寅),
            "卯" => Some(卯),
            "辰" => Some(辰),
            "巳" => Some(巳),
            "午" => Some(午),
            "未" => Some(未),
            "申" => Some(申),
            "酉" => Some(酉),
            "戌" => Some(戌),
            "亥" => Some(亥),
            _ => return HttpResponse::BadRequest().json(format!("没有此地支：{}", s)),
        }
    } else {
        None
    };

    let divination_time = match r.divination_time.as_str() {
        "子" => 子,
        "丑" => 丑,
        "寅" => 寅,
        "卯" => 卯,
        "辰" => 辰,
        "巳" => 巳,
        "午" => 午,
        "未" => 未,
        "申" => 申,
        "酉" => 酉,
        "戌" => 戌,
        "亥" => 亥,
        _ => return HttpResponse::BadRequest().json(format!("没有此地支：{}", r.divination_time)),
    };

    let pan = DaliurenShiPan::new(
        r.year,
        r.month,
        r.day,
        r.hour,
        r.minute,
        r.second,
        yue_jiang,
        divination_time,
        r.diurnal,
        r.year_of_birth,
        r.masculine,
        &app_state.ephe_path,
    );
    match pan {
        Ok(pan) => HttpResponse::Ok().json(pan),
        Err(e) => HttpResponse::BadRequest().json(e),
    }
}
