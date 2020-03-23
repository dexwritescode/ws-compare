package webservice;

import akka.actor.AbstractActor;
import akka.actor.Props;
import akka.japi.pf.FI;

import webservice.HelloMessages.GetHelloMessage;

public class HelloActor extends AbstractActor {
    private HelloService helloService = new HelloService();

    static Props props() {
        return Props.create(HelloActor.class);
    }

    @Override
    public Receive createReceive() {
        return receiveBuilder()
                .match(HelloMessages.GetHelloMessage.class, handleGetHello())
                .build();
    }

    private FI.UnitApply<GetHelloMessage> handleGetHello() {
        return getHelloMessage -> {
            sender().tell(helloService.getHello(), getSelf());
        };
    }
}
