package org.solution;

import lombok.Getter;

@Getter
public class Message {
    private final String type;
    private final String name;

    public Message(String type, String name) {
       this.type = type;
       this.name = name;
    }

    public String createStringMsg() {
        return type + " " + name;
    }
}
