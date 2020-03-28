package webservice;

public class HelloService {
    public String getHello(String name) {
        return String.format("Hello, %s!", name);
    }
}
