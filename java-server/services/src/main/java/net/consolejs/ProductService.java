package net.consolejs;

import javax.enterprise.context.ApplicationScoped;
import java.util.ArrayList;
import java.util.List;

@ApplicationScoped
public class ProductService {
    private static final List<ProductEntity> myItems = getFakeItems();

    public PaginationEntity<ProductEntity> getProductsPaginated(int page, int size) {
        int start = (size * page) - size;
        List<ProductEntity> items = ProductService.myItems.subList(start, start + size);

        return PaginationEntity.<ProductEntity>newBuilder()
                .withItems(items)
                .withPage(page)
                .withLastPage(ProductService.myItems.size() / size)
                .build();
    }

    private static List<ProductEntity> getFakeItems() {
        List<ProductEntity> items = new ArrayList<>();
        for (int i = 0; i < 100; i++) {
            items.add(ProductEntity.newBuilder()
                    .withId(i)
                    .withName("Product #" + i)
                    .withDescription("This is a random description for the product")
                    .build());
        }

        return items;
    }
}

