package org.example;

import com.zeroc.Ice.Communicator;
import com.zeroc.Ice.ObjectAdapter;
import com.zeroc.Ice.Util;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Server {
    private static final String ADAPTER_NAME = "adapter";
    private static final String SERVANT_LOCATOR_PREFIX = "";
    private final Communicator communicator;
    private final ObjectAdapter adapter;
    private final Locator locator;

    public Server(int serverId, String[] args) {
        this.communicator = Util.initialize(args);
        String adapterEndpoints = getAdapterEndpoints(serverId);
        System.out.println(adapterEndpoints);
        this.adapter = communicator.createObjectAdapterWithEndpoints(ADAPTER_NAME, adapterEndpoints);
        this.locator = new Locator(String.valueOf(serverId));
        adapter.addServantLocator(locator, SERVANT_LOCATOR_PREFIX);
    }

    private String getAdapterEndpoints(int serverId) {
        int offset = 10;
        int offsetServerId = offset + serverId;
        return "tcp -h 127.0.0.!! -p 100!! : udp -h 127.0.0.!! -p 100!!"
                .replace("!!", String.valueOf(offsetServerId));
    }

    public void run() {
        adapter.activate();
        System.out.println("in the loop...");

        int status = 0;
        boolean quit = false;
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        while (!quit) {
            String line;
            try {
                line = in.readLine();
            } catch (IOException e) {
                break;
            }

            switch (line) {
                case "devices":
                    locator.servantToString();
                    break;
                case "quit":
                    quit = true;
                    try {
                        adapter.deactivate();
                        communicator.shutdown();
                        communicator.destroy();
                    } catch (Exception e) {
                        System.out.println(e.getMessage());
                        status = 1;
                    }
                    break;
                default:
                    System.out.println("Invalid command");
            }
        }
        System.exit(status);
    }
}
