package org.solution;

import java.io.IOException;
import java.util.*;

import org.apache.zookeeper.*;

public class DataMonitor implements Watcher {
    ZooKeeper zk;
    String znode;
    String[] exec;
    Process client;

    public DataMonitor(String hostPort, String znode, String[] exec) throws IOException {
        this.zk = new ZooKeeper(hostPort, 3000, this);
        this.znode = znode;
        this.exec = exec;
        zk.exists(znode, this, null, true);
    }

    @Override
    public void process(WatchedEvent watchedEvent) {
        switch (watchedEvent.getType()) {
            case Event.EventType.NodeCreated -> {
                System.out.println(watchedEvent.getPath());
                try {
                    handleCreation(watchedEvent.getPath());
                } catch (InterruptedException | KeeperException e) {
                    throw new RuntimeException(e);
                }
            }
            case Event.EventType.NodeDeleted -> {
                handleDeletion(watchedEvent.getPath());
            }
        }
    }

    public void runClient() throws InterruptedException, KeeperException {
        this.zk.addWatch(znode, AddWatchMode.PERSISTENT_RECURSIVE);
        loop();
    }

    private void handleDeletion(String watchedEventPath) {
        if (znode.equals(watchedEventPath)) {
            System.out.println(client);
            closeClient();
        }
    }

    private void handleCreation(String watchedEventPath) throws InterruptedException, KeeperException {
        if (watchedEventPath.startsWith(znode)) {
            System.out.println(znode.length() + " " + watchedEventPath.length());
            if (watchedEventPath.length() == znode.length()) {
                System.out.println(" [INFO] Created watched znode...");
                try {
                    System.out.println(" [INFO] Exec is running...");
                    client = Runtime.getRuntime().exec(exec);
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            } else {
                System.out.println(" [INFO] Created child " + zk.getChildren(znode, true) + " for " + znode);
                List<String> children = createTree();
                System.out.printf(" [INFO] There is %d children\n", children.toArray().length - 1);
            }
        }
    }

    private void closeClient() {
        if (client == null || !client.isAlive()) {
            System.out.println(" [INFO] Process already dead...");
        } else {
            client.destroy();
            System.out.println(" [INFO] Killing client process...");
        }
    }

    private void loop() throws InterruptedException, KeeperException {
        Scanner scanner = new Scanner(System.in);
        while (true) {
            String line = scanner.nextLine();
            System.out.println(line);
            if (line.equals("quit")) {
                closeClient();
                zk.close();
                break;
            } else if (line.equals("show tree")) {
                System.out.println(" [INFO] Printing znode tree...");
                printTree();
            }
        }
    }

    private List<String> createTree() throws InterruptedException, KeeperException {
        Deque<String> queue = new ArrayDeque<>();
        queue.add(znode);
        String father;
        List<String> children;
        List<String> result = new ArrayList<>();
        while (!queue.isEmpty()) {
            father = queue.remove();
            children = zk.getChildren(father, false);
            for (String child : children) {
                queue.add(String.format("%s/%s", father, child));
            }
            result.add(father);
        }
        return result;
    }

    private void printTree() throws InterruptedException, KeeperException {
       List<String> tree = createTree();

        for (String treePart : tree) {
            StringBuilder sb =  new StringBuilder();
            StringTokenizer tokenizer = new StringTokenizer(treePart, "/");
            String token = null;
            while (tokenizer.hasMoreTokens()) {
                if (token != null) {
                    sb.append("    ");
                }
                token = tokenizer.nextToken();
            }
            sb.append("|-- ");
            sb.append(token);
            System.out.println(sb);
        }
    }
}