package webservice;

public class FibonacciService {

    long getFibonacci(long number) {
        if (number <= 1)
            return number;
        return getFibonacci(number - 1) + getFibonacci(number - 2);
    }
}

