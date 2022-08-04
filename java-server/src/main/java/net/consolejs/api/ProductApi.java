package net.consolejs.api;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.PathParam;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import net.consolejs.entities.PaginationEntity;
import net.consolejs.entities.ProductEntity;

import java.util.ArrayList;
import java.util.List;

@Path("/api/v1/products")
@Produces(MediaType.APPLICATION_JSON)
public class ProductApi {

    private static final List<ProductEntity> myItems = getFakeItems();
    private static final int PER_PAGE = 10;

    @Path("/page/{page}")
    @GET
    public Response getProducts(@PathParam("page") int page) {
        int start = (PER_PAGE * page) - PER_PAGE;
        List<ProductEntity> items = ProductApi.myItems.subList(start, start + PER_PAGE);

        return Response.status(Response.Status.OK).entity(PaginationEntity.<ProductEntity>newBuilder()
                .withItems(items)
                .withPage(page)
                .withLastPage(ProductApi.myItems.size() / PER_PAGE)
                .build()).build();
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
