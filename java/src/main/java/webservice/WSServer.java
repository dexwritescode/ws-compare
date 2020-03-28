package webservice;

import akka.http.javadsl.model.StatusCodes;
import akka.http.javadsl.server.HttpApp;
import akka.http.javadsl.server.Route;

import static akka.http.javadsl.server.PathMatchers.*;

public class WSServer extends HttpApp {
    private final HelloService helloService;
    private final FibonacciService fibonacciService;

    WSServer() {
        this.helloService = new HelloService();
        this.fibonacciService = new FibonacciService();
    }

    @Override
    public Route routes() {
        return getAllRoutes();
    }

    private Route getAllRoutes() {

        return get(() -> concat(
                path(segment("hello").slash(remaining()), this::hello),
                path(segment("fibonacci").slash(longSegment()), this::fibonacci)
        ));
    }

    private Route hello(String name) {
        return complete(StatusCodes.OK, helloService.getHello(name));
    }

    private Route fibonacci(long number) {
        return complete(StatusCodes.OK, String.valueOf(fibonacciService.getFibonacci(number)));
    }
}
