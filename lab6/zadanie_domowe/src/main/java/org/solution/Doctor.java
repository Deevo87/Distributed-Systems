package org.solution;

import com.rabbitmq.client.AMQP;
import com.rabbitmq.client.Channel;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.CompletableFuture;

public class Doctor {
    private final Channel channel;
    private final String exchangeName;
    private final String callbackQueueName;
    private Map<String, CompletableFuture<String>> pendingRequest = new HashMap<>();

    public Doctor(String exchangeName, Channel channel) throws IOException {
        this.exchangeName = exchangeName;
        this.channel = channel;
        this.callbackQueueName = channel.queueDeclare().getQueue();

        channel.basicConsume(callbackQueueName, true, (consumerTag, delivery) -> {
            String corrId = delivery.getProperties().getCorrelationId();
            if (corrId != null && pendingRequest.containsKey(corrId)) {
                String response = new String(delivery.getBody());
                pendingRequest.get(corrId).complete(response);
                pendingRequest.remove(corrId);
            }
        }, consumerTag -> {});
    }

    public CompletableFuture<String> call(Message msg) throws IOException {
        final String corrId = UUID.randomUUID().toString();

        AMQP.BasicProperties props = new AMQP.BasicProperties
                .Builder()
                .correlationId(corrId)
                .replyTo(callbackQueueName)
                .build();

        channel.exchangeDeclare(exchangeName, "direct");
        CompletableFuture<String> response = new CompletableFuture<>();
        pendingRequest.put(corrId, response);

        channel.basicPublish(exchangeName, msg.getType(), props, msg.createStringMsg().getBytes());

        return response;
    }
}
