package net.consolejs.service;

import org.jvnet.hk2.annotations.Service;

@Service
public class MessageServiceImpl implements MessageService {

    @Override
    public String getHello() {
        return "Hello World Jersey from HK2";
    }

}
