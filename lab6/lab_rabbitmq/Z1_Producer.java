import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;

import java.util.Scanner;

public class Z1_Producer {

    public static void main(String[] argv) throws Exception {
        // info
        System.out.println("Z1 PRODUCER");

        // connection & channel
        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost("localhost");
        Connection connection = factory.newConnection();
        Channel channel = connection.createChannel();

        // queue
        String QUEUE_NAME = "queue1";
        channel.queueDeclare(QUEUE_NAME, false, false, false, null);

        // producer (publish msg)
        Scanner scanner = new Scanner(System.in);

        System.out.println("> Twoja wiadomość: ");
        String message = scanner.nextLine();

        System.out.println("> To zajmie nam: ");
        String time = scanner.nextLine();

        int timeToSleep = Integer.parseInt(time);
        Thread.sleep(timeToSleep * 1000);

        channel.basicPublish("", QUEUE_NAME, null, message.getBytes());
        System.out.println("Sent: " + message);

        // close
        channel.close();
        connection.close();
    }
}
