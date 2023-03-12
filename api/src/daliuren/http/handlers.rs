use crate::{daliuren::shi_pan::DaliurenShiPan, state::AppState};
use actix_web::{post, web, HttpResponse, Responder};

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
    let pan = DaliurenShiPan::new(
        r.year,
        r.month,
        r.day,
        r.hour,
        r.minute,
        r.second,
        r.yue_jiang,
        &r.divination_time,
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
