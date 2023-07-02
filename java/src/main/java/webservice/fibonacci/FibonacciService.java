package webservice.fibonacci;

public class FibonacciService {

    public FibonacciView fibonacci(long number) {
        return new FibonacciView(number, fib(number));
    }

    private long fib(long number) {
        if (number <= 1)
            return number;
        return fib(number - 1) + fib(number - 2);
    }
}

