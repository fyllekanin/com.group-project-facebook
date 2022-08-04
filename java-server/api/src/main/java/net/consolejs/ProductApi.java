package net.consolejs;

import org.eclipse.microprofile.openapi.annotations.Operation;
import org.eclipse.microprofile.openapi.annotations.media.Content;
import org.eclipse.microprofile.openapi.annotations.media.Schema;
import org.eclipse.microprofile.openapi.annotations.responses.APIResponse;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.core.Response;

import static javax.ws.rs.core.MediaType.APPLICATION_JSON;
import static org.eclipse.microprofile.openapi.annotations.enums.SchemaType.OBJECT;

@Path("/api/v1/products")
public class ProductApi {

    @Inject()
    ProductService myProductService;

    @GET
    @Path("/page/{page}")
    @Operation(operationId = "GetProduct",
            summary = "Paginated object of Products"
    )
    @APIResponse(responseCode = "200",
            content = @Content(mediaType = APPLICATION_JSON,
                    schema = @Schema(type = OBJECT,
                            implementation = ProductEntity.class)))
    public Response getProducts(@PathParam("page") int page) {
        return Response.status(Response.Status.OK)
                .entity(myProductService.getProductsPaginated(page, 10))
                .build();
    }
}
