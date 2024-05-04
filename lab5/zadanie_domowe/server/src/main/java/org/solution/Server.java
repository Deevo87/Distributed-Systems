package org.solution;

import io.grpc.Grpc;
import io.grpc.InsecureServerCredentials;

import java.io.IOException;

public class Server {
    private int port = 5000;
    private io.grpc.Server server;

    public Server() {}

    public void run() {
        try {
            this.server = Grpc.newServerBuilderForPort(port, InsecureServerCredentials.create())
                    .addService(new ExecutionServiceImpl())
                    .build()
                    .start();
            System.out.println("Listening on port " + port);
            server.awaitTermination();
        } catch (IOException | InterruptedException e) {
            System.err.println("Failed to start server");
            throw new RuntimeException(e);
        }
    }
}
