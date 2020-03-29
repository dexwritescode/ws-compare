package webservice.greeting;

public class GreetingView {
    public final String greeting;

    GreetingView(String name) {
        this.greeting = String.format("Hello, %s!", name);
    }
}
