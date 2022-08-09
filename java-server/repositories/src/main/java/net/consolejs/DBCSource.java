package net.consolejs;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class DBCSource {
    private static Connection CONNECTION = null;

    public static Connection getConnection() throws SQLException {
        if (CONNECTION == null) {
            CONNECTION = DriverManager.getConnection("jdbc:postgresql://localhost:5432/go-server", "username", "password");
        }
        return CONNECTION;
    }
}
