package net.consolejs;

import net.consolejs.config.AutoScanFeature;
import org.glassfish.grizzly.http.server.HttpServer;
import org.glassfish.jersey.grizzly2.httpserver.GrizzlyHttpServerFactory;
import org.glassfish.jersey.server.ResourceConfig;

import java.net.URI;
import java.util.logging.Level;
import java.util.logging.Logger;

public class MainApp {

    private static final Logger LOGGER = Logger.getLogger(MainApp.class.getName());
    public static final String BASE_URI = "http://0.0.0.0:8080/";

    public static HttpServer startServer() {
        final ResourceConfig config = new ResourceConfig();
        config.packages(true, "net.consolejs.api");
        config.register(AutoScanFeature.class);

        LOGGER.info("Starting Server........");
        return GrizzlyHttpServerFactory.createHttpServer(URI.create(BASE_URI), config);

    }

    public static void main(String[] args) {

        try {
            final HttpServer httpServer = startServer();
            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                try {
                    System.out.println("Shutting down the application...");

                    httpServer.shutdownNow();

                    System.out.println("Done, exit.");
                } catch (Exception e) {
                    Logger.getLogger(MainApp.class.getName()).log(Level.SEVERE, null, e);
                }
            }));

            System.out.printf("Application started.%nStop the application using CTRL+C%n");
            Thread.currentThread().join();

        } catch (InterruptedException ex) {
            Logger.getLogger(MainApp.class.getName()).log(Level.SEVERE, null, ex);
        }

    }
}
