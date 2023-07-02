use actix_web::{Responder, HttpResponse, web};
use serde::Deserialize;

mod fibonacci;

#[derive(Deserialize)]
pub struct Info {
    input: i64,
}

pub async fn get(info: web::Path<Info>) -> impl Responder {
    HttpResponse::Ok().json(fibonacci::get(info.input))
}