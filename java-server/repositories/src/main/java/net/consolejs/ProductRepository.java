package net.consolejs;

import javax.enterprise.context.ApplicationScoped;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;

@ApplicationScoped
public class ProductRepository {

    public List<ProductEntity> getProducts(int start, int limit) {
        try {
            PreparedStatement preparedStatement = DBCSource.getConnection().prepareStatement("SELECT * FROM products LIMIT ? OFFSET ?");
            preparedStatement.setInt(1, limit);
            preparedStatement.setInt(2, start);

            ResultSet resultSet = preparedStatement.executeQuery();
            List<ProductEntity> response = new ArrayList<>();
            while (resultSet.next()) {
                response.add(ProductEntity.newBuilder()
                        .withId(resultSet.getInt(1))
                        .withName(resultSet.getString(2))
                        .withDescription(resultSet.getString(3))
                        .withPrice(resultSet.getInt(4))
                        .withCreatedAt(resultSet.getInt(5))
                        .withUpdatedAt(resultSet.getInt(6))
                        .build());
            }
            return response;
        } catch (SQLException sqlException) {
            return new ArrayList<>();
        }
    }

    public int getProductsCount() {
        try {
            PreparedStatement preparedStatement = DBCSource.getConnection().prepareStatement("SELECT COUNT(*) from products");
            ResultSet resultSet = preparedStatement.executeQuery();
            while (resultSet.next()) {
                return resultSet.getInt(1);
            }
        } catch (SQLException sqlException) {
            return 0;
        }
        return 0;
    }
}