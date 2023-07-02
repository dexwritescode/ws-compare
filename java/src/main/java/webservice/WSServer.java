package webservice;

import akka.http.javadsl.marshallers.jackson.Jackson;
import akka.http.javadsl.model.StatusCodes;
import akka.http.javadsl.server.HttpApp;
import akka.http.javadsl.server.Route;
import webservice.fibonacci.FibonacciService;
import webservice.greeting.GreetingService;
import webservice.hello.HelloService;

import static akka.http.javadsl.server.PathMatchers.*;

public class WSServer extends HttpApp {
    private final GreetingService greetingService;
    private final FibonacciService fibonacciService;
    private final HelloService helloService;

    WSServer() {
        this.greetingService = new GreetingService();
        this.fibonacciService = new FibonacciService();
        this.helloService = new HelloService();
    }

    @Override
    public Route routes() {
        return getAllRoutes();
    }

    private Route getAllRoutes() {

        return get(() -> concat(
                path("hello", this::helloWorld),
                path(segment("greeting").slash(remaining()), this::greeting),
                path(segment("fibonacci").slash(longSegment()), this::fibonacci)
        ));
    }

    private Route helloWorld() {

        return complete(StatusCodes.OK, helloService.hello(), Jackson.marshaller());
    }

    private Route greeting(String name) {
        return complete(StatusCodes.OK, greetingService.greet(name), Jackson.marshaller());
    }

    private Route fibonacci(long number) {
        return complete(StatusCodes.OK, fibonacciService.fibonacci(number), Jackson.marshaller());
    }
}
