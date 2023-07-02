use actix_web::{HttpRequest, Responder, HttpResponse};

mod helloworld;

pub async fn get(_req: HttpRequest) -> impl Responder {
    HttpResponse::Ok().json(helloworld::get())
}