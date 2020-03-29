package webservice;

import akka.actor.ActorSystem;

public class WebServiceMain {

    public static void main(String[] args) throws Exception {
        ActorSystem system = ActorSystem.create("webserviceServer");
        WSServer server = new WSServer();
        server.startServer("localhost", 8080, system);
    }
}
