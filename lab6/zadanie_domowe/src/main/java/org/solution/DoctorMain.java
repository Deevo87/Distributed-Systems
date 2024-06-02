package org.solution;

import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;

import java.util.Arrays;
import java.util.Scanner;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class DoctorMain {

    private static final String EXCHANGE_NAME = "technician_queue";

    public static void main(String[] argv) throws Exception {
        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost("localhost");
        try (Connection connection = factory.newConnection();
             Channel channel = connection.createChannel()) {
            Doctor doctor = new Doctor(EXCHANGE_NAME, channel);
            Scanner scanner = new Scanner(System.in);
            ExecutorService executorService = Executors.newCachedThreadPool();

            while (true) {
                System.out.print(" [-] Type new client: ");
                String line = scanner.nextLine();
                String[] preFormat = line.split(" [-] Type new client: ");
                String[] splitLine = preFormat[0].split(" ");
                if (splitLine.length == 3) {
                    String command = splitLine[0];
                    if (command.equals("quit")) {
                        executorService.shutdown();
                        return;
                    } else {
                        String name = splitLine[1] + " " + splitLine[2];
                        Message msg = new Message(command, name);
                        CompletableFuture<String> responseFuture = doctor.call(msg);
                        responseFuture.thenAccept(response -> {
                            System.out.println("\n [x] Received: " + response);
                        });
                    }
                } else {
                    System.out.println(" [*] Wrong command! Usage: <type firstName secondName>");
                }
            }
        }
    }
}
