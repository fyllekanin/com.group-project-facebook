package net.consolejs.api;

import jakarta.ws.rs.core.Response;
import net.consolejs.entities.PaginationEntity;
import net.consolejs.entities.ProductEntity;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class ProductApiTest {
    private ProductApi PRODUCT_API;

    @BeforeEach
    void configure() {
        PRODUCT_API = new ProductApi();
    }

    @Test
    public void shouldReturnTenItems() {
        // Given
        int page = 1;

        // When
        Response response = PRODUCT_API.getProducts(page);

        // Then
        PaginationEntity<ProductEntity> entity = (PaginationEntity<ProductEntity>) response.getEntity();
        assertEquals(entity.getItems().size(), 10);
    }

    @Test
    public void shouldReturnCorrectPage() {
        // Given
        int page = 5;

        // When
        Response response = PRODUCT_API.getProducts(page);

        // Then
        PaginationEntity<ProductEntity> entity = (PaginationEntity<ProductEntity>) response.getEntity();
        assertEquals(entity.getPage(), 5);
    }

    @Test
    public void shouldReturnCorrectLastPage() {
        // Given
        int page = 1;

        // When
        Response response = PRODUCT_API.getProducts(page);

        // Then
        PaginationEntity<ProductEntity> entity = (PaginationEntity<ProductEntity>) response.getEntity();
        assertEquals(entity.getLastPage(), 10);
    }
}
