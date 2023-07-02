use actix_web::{Responder, HttpResponse, web};
use serde::Deserialize;

mod greeting;

#[derive(Deserialize)]
pub struct Info {
    name: String,
}

pub async fn get(info: web::Path<Info>) -> impl Responder {
    HttpResponse::Ok().json(greeting::get(&info.name))
}