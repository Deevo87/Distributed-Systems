package org.example;

import org.apache.logging.log4j.core.config.Configurator;
import org.apache.logging.log4j.core.config.DefaultConfiguration;
import java.util.Arrays;


public class Main {
    public static void main(String[] args) {
        Configurator.initialize(new DefaultConfiguration());
//        Configurator.setRootLevel(Level.INFO);

        if (args.length < 1) {
            System.out.println("You need to type in serverId");
        } else {
            try {
                int serverId = Integer.parseInt(args[0]);
                if (serverId < 0 || serverId > 89) {
                    System.out.println("serverId out of supported range [0, 89]");
                } else {
                    String[] iceArgs = Arrays.copyOfRange(args, 1, args.length);

                    Server server = new Server(serverId, iceArgs);
                    server.run();
                    System.out.println("Server started successfully");
                }
            } catch (NumberFormatException e) {
                System.out.println("Invalid serverId format");
                System.exit(1);
            }
        }
    }
}