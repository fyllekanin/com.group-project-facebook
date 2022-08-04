package net.consolejs.entities;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class ProductEntityTest {

    @Test
    public void shouldBuildCorrectly() {
        // Given & When
        ProductEntity productEntity = ProductEntity.newBuilder()
                .withId(1)
                .withName("name")
                .withDescription("description")
                .withPrice(50)
                .build();

        // Then
        assertEquals(productEntity.getId(), 1);
        assertEquals(productEntity.getName(), "name");
        assertEquals(productEntity.getDescription(), "description");
        assertEquals(productEntity.getPrice(), 50);
    }

}
