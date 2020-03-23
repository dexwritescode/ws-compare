package webservice;

import akka.actor.ActorRef;
import akka.actor.ActorSystem;

public class WebService {

    public static void main(String[] args) throws Exception {
        ActorSystem system = ActorSystem.create("webserviceServer");
        ActorRef helloActor = system.actorOf(HelloActor.props(), "helloActor");
        WSServer server = new WSServer(helloActor);
        server.startServer("localhost", 8080, system);
    }

}
