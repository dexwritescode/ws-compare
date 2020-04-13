//use actix_web::{web, App, HttpRequest, HttpServer, Responder, HttpResponse};
use actix;
/*async fn greet(req: HttpRequest) -> impl Responder {
    let name = req.match_info().get("name").unwrap_or("World");
    format!("Hello {}!", &name)
}

async fn hello_world(req: HttpRequest) -> impl Responder {
    HttpResponse::Ok().body("Hello, World!")
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(hello_world))
            .route("/{name}", web::get().to(greet))
    })
        .bind("127.0.0.1:8000")?
        .run()
        .await
}*/

mod app;

fn main() {
    let sys = actix::System::new("webservice");

    app::start();

    let _ = sys.run();
}
