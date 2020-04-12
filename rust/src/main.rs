use actix;

mod app;

fn main() {
    let sys = actix::System::new("webservice");
    app::start();
    let _ = sys.run();
}
