package com.example.stomp;

import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.stereotype.Controller;

@Controller
public class GreetingController {

    @MessageMapping("/greeting")
    @SendTo("/topic/greeting.back")
    public String handle(String greeting) {
        return "server: " + greeting;
    }

}
