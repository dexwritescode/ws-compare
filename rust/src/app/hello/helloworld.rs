use serde::Serialize;

#[derive(Serialize)]
pub struct HelloResponse<'a> {
    pub message: &'a str,
}

impl<'a> HelloResponse<'a> {
    fn get() -> Self {
        HelloResponse {
            message: "Hello, World!"
        }
    }
}

pub fn get<'a>() -> HelloResponse<'a> {
    HelloResponse::get()
}