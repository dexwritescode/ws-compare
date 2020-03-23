package webservice;

import akka.actor.ActorRef;
import akka.http.javadsl.marshallers.jackson.Jackson;
import akka.http.javadsl.model.StatusCodes;
import akka.http.javadsl.server.HttpApp;
import akka.http.javadsl.server.Route;
import akka.pattern.Patterns;
import webservice.HelloMessages.GetHelloMessage;

import java.util.concurrent.CompletionStage;

public class WSServer extends HttpApp {
    private final ActorRef helloActor;

    java.time.Duration timeout = java.time.Duration.ofSeconds(5);

    WSServer(ActorRef helloActor) {
        this.helloActor = helloActor;
    }

    @Override
    public Route routes() {
        return path("hello", this::getHello);
    }

    private Route getHello() {
        return get(() -> {
            CompletionStage<String> helloCompletionStage = Patterns.ask(helloActor, new GetHelloMessage(), timeout)
                    .thenApply(obj -> (String) obj);

            return onSuccess(() -> helloCompletionStage, performed -> {
                return complete(StatusCodes.OK, performed, Jackson.marshaller());
            });
        });
    }
}
