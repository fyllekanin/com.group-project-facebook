package net.consolejs;

import lombok.Data;
import org.eclipse.microprofile.openapi.annotations.media.Schema;

@Data
public class ProductEntity {

    @Schema(hidden = true)
    private final int myId;
    @Schema(hidden = true)
    private final String myName;
    @Schema(hidden = true)
    private final String myDescription;
    @Schema(hidden = true)
    private final int myPrice;
    @Schema(hidden = true)
    private final int myCreatedAt;
    @Schema(hidden = true)
    private final int myUpdatedat;

    private ProductEntity(ProductEntityBuilder builder) {
        myId = builder.myId;
        myName = builder.myName;
        myDescription = builder.myDescription;
        myPrice = builder.myPrice;
        myCreatedAt = builder.myCreatedAt;
        myUpdatedat = builder.myUpdatedAt;
    }

    public int getId() {
        return myId;
    }

    public String getName() {
        return myName;
    }

    public String getDescription() {
        return myDescription;
    }

    public int getPrice() {
        return myPrice;
    }

    public int getCreatedAt() {
        return myCreatedAt;
    }

    public int getUpdatedAt() {
        return myUpdatedat;
    }

    public static ProductEntityBuilder newBuilder() {
        return new ProductEntityBuilder();
    }

    public static class ProductEntityBuilder {
        private int myId;
        private String myName;
        private String myDescription;
        private int myPrice;
        private int myCreatedAt;
        private int myUpdatedAt;

        private ProductEntityBuilder() {
        }

        public ProductEntityBuilder withId(int id) {
            myId = id;
            return this;
        }

        public ProductEntityBuilder withName(String name) {
            myName = name;
            return this;
        }

        public ProductEntityBuilder withDescription(String description) {
            myDescription = description;
            return this;
        }

        public ProductEntityBuilder withPrice(int price) {
            myPrice = price;
            return this;
        }

        public ProductEntityBuilder withCreatedAt(int createdAt) {
            myCreatedAt = createdAt;
            return this;
        }

        public ProductEntityBuilder withUpdatedAt(int updatedAt) {
            myUpdatedAt = updatedAt;
            return this;
        }

        public ProductEntity build() {
            return new ProductEntity(this);
        }
    }
}