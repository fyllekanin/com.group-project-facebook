package net.consolejs.entities;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class PaginationEntityTest {

    @Test
    public void shouldBuildCorrectly() {
        // Given & When
        PaginationEntity<ArrayList<?>> paginationEntity = PaginationEntity.<ArrayList<?>>newBuilder()
                .withPage(5)
                .withLastPage(10)
                .withItems(List.of(new ArrayList<>()))
                .build();

        // Then
        assertEquals(paginationEntity.getPage(), 5);
        assertEquals(paginationEntity.getLastPage(), 10);
        assertEquals(paginationEntity.getItems(), List.of(new ArrayList<>()));
    }

}
