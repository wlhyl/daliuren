use actix_web::web;

use super::handlers::daliuren;

pub fn daliuren_routes(cfg: &mut web::ServiceConfig) {
    cfg.service(daliuren);
}