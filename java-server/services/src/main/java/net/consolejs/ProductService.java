package net.consolejs;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.List;

@ApplicationScoped
public class ProductService {
    @Inject()
    ProductRepository productRepository;

    public PaginationEntity<ProductEntity> getProductsPaginated(int page, int size) {
        int start = (size * page) - size;
        List<ProductEntity> items = productRepository.getProducts(start, size);

        return PaginationEntity.<ProductEntity>newBuilder()
                .withItems(items)
                .withPage(page)
                .withLastPage((productRepository.getProductsCount() / size) + 1)
                .build();
    }
}

