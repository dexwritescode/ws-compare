use serde::Serialize;

#[derive(Serialize)]
pub struct HelloResponse {
    pub message: String,
}

impl HelloResponse {
    fn get() -> Self {
        HelloResponse {
            message: "Hello, World!".to_owned()
        }
    }
}

pub fn get() -> HelloResponse {
    HelloResponse::get()
}