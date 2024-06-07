package org.solution;

import java.io.IOException;

import org.apache.zookeeper.KeeperException;

public class Main {

    public static void main(String[] args) throws IOException, KeeperException, InterruptedException {
        if (args.length < 3) {
            System.err
                    .println("USAGE: Executor hostPort znode program [args ...]");
            System.exit(2);
        }
        String hostPort = args[0];
        String znode = args[1];
        String[] exec = new String[args.length - 2];
        System.arraycopy(args, 2, exec, 0, exec.length);
        DataMonitor dataMonitor = new DataMonitor(hostPort, znode, exec);
        dataMonitor.runClient();
    }
}
