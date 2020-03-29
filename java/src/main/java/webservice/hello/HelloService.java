package webservice.hello;

public class HelloService {

    public HelloView hello() {
        return new HelloView("Hello, World");
    }
}
