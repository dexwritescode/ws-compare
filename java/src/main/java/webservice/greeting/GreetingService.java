package webservice.greeting;

public class GreetingService {

    public GreetingView greet(String name) {
        return new GreetingView(name);
    }
}
