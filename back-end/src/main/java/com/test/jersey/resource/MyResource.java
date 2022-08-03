package com.test.jersey.resource;

import com.test.jersey.service.MessageService;
import jakarta.inject.Inject;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.PathParam;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;

class Testing {

    private final String value;

    Testing(String value) {
        this.value = value;
    }

    public String getValue() {
        return this.value;
    }
}

@Path("/api/hello")
public class MyResource {

    // DI via HK2
    @Inject
    private MessageService messageService;

    // output text
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public Response hello() {
        return Response.status(Response.Status.OK).entity(new Testing("rawr")).build();
    }

    // output text with argument
    @Path("/{name}")
    @GET
    @Produces(MediaType.TEXT_PLAIN)
    public String hello(@PathParam("name") String name) {
        return "Jersey: hello " + name;
    }

    // for dependency injection
    @Path("/hk2")
    @GET
    @Produces(MediaType.TEXT_PLAIN)
    public String helloHK2() {
        return messageService.getHello();
    }

}
