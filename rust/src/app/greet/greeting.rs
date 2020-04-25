use serde::Serialize;

#[derive(Serialize)]
pub struct GreetingResponse {
    pub greeting: String,
}

impl GreetingResponse {
    pub fn new(greeting: String) -> Self {
        GreetingResponse {
            greeting
        }
    }
}

pub fn get(name: &String) -> GreetingResponse {
    let mut greeting: String = String::from("Hello, ");
    greeting.push_str(&name);
    greeting.push('!');
    GreetingResponse::new(greeting)
}