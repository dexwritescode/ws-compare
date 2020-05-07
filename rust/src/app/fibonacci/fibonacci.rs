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

pub fn fibonacci(n: i64) -> i64 {
    fn fibonacci_lr(n: i64, a: i64, b: i64) -> i64 {
        match n {
            0 => a,
            _ => fibonacci_lr(n - 1, a + b, a),
        }
    }
    fibonacci_lr(n, 1, 0)
}

pub fn get(input: i64) -> FibonacciResponse {
    let output = fibonacci(input);
    FibonacciResponse::new(input, output)
}

