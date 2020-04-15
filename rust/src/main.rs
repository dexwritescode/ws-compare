//use actix_web::{web, App, HttpRequest, HttpServer, Responder, HttpResponse};
use actix;
/*async fn app.greet(req: HttpRequest) -> impl Responder {
    let name = req.match_info().get("name").unwrap_or("World");
    format!("Hello {}!", &name)
}

async fn hello_world(req: HttpRequest) -> impl Responder {
    HttpResponse::Ok().body("Hello, World!")
}
*/

mod app;
//mod hello;

fn main() {
    let sys = actix::System::new("webservice");

    app::start();

    let _ = sys.run();
}
