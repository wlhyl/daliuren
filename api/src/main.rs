use std::{env, net::SocketAddrV4};

use actix_cors::Cors;
use actix_web::{middleware::Logger, web, App, HttpServer};
use clap::Parser;
use daliuren::{args, daliuren::daliuren_routes, routers::health_routes, state::AppState};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv::dotenv().ok();

    let log4rs_config = env::var("LOG4RS_CONFIG")
        .expect("没设置 LOG4RS_CONFIG 环境变量，可在.env文件中设置或export LOG4RS_CONFIG=...");

    log4rs::init_file(log4rs_config, Default::default()).unwrap();

    let ephe_path = env::var("EPHE_PATH")
        .expect("没设置 EPHE_PATH 环境变量，可在.env文件中设置或export EPHE_PATH=...");

    let shared_data = web::Data::new(AppState { ephe_path });

    let args = args::Args::parse();
    HttpServer::new(move || {
        let cors = Cors::default()
            .allow_any_origin()
            .allow_any_method()
            .allow_any_header()
            .max_age(3600);

        App::new()
            .app_data(shared_data.clone())
            .configure(health_routes)
            .service(web::scope("/api").configure(daliuren_routes))
            .wrap(cors)
            .wrap(Logger::default())
    })
    .workers(args.n)
    .bind(SocketAddrV4::new(args.ip, args.port))?
    .run()
    .await
}
