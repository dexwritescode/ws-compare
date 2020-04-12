use serde::Serialize;

#[derive(Serialize)]
pub struct FibonacciResponse {
    pub input: i64,
    pub output: i64
}

impl FibonacciResponse {
    pub fn new(input: i64, output: i64) -> Self {
        FibonacciResponse {
            input,
            output
        }
    }
}

fn fibonacci(number: i64) -> i64 {
    if number <= 1 {
        return number;
    }
    return fibonacci(number-1) + fibonacci(number-2)
}

pub fn get(input: i64) -> FibonacciResponse {
    let output = fibonacci(input);
    FibonacciResponse::new(input, output)
}

