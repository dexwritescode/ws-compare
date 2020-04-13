use actix_web::{HttpServer, App, web, HttpRequest, Responder, HttpResponse};
use serde::Serialize;

pub fn start() {
    let bind_address = "127.0.0.1:8080";
    HttpServer::new(move|| {
        App::new().configure(routes)
        })
        .bind(&bind_address)
        .unwrap_or_else(|_| panic!("Could not bind server to address {}", &bind_address))
        .run();

    println!("You can access the server at {}", &bind_address);
}

fn routes(app: &mut web::ServiceConfig) {
    app
        .route("/", web::get().to(get_hello))
        .route("/greeting/{name}", web::get().to(greet));
}

async fn greet(req: HttpRequest) -> impl Responder {
    let name = req.match_info().get("name").unwrap_or("World");
    HttpResponse::Ok().body(format!("Hello, {}!", &name))
}

async fn get_hello(_req: HttpRequest) -> impl Responder {
    HttpResponse::Ok().json(HelloResponse::get())
}

#[derive(Serialize)]
pub struct HelloResponse {
    pub message: &'static str,
}

impl HelloResponse {
    fn get() -> Self {
        HelloResponse {
            message: "Hello, World!"
        }
    }
}
