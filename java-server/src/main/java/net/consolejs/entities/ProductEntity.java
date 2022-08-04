package net.consolejs.entities;

public class ProductEntity {
    private final int myId;
    private final String myName;
    private final String myDescription;
    private final int myPrice;

    private ProductEntity(ProductEntityBuilder builder) {
        myId = builder.myId;
        myName = builder.myName;
        myDescription = builder.myDescription;
        myPrice = builder.myPrice;
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

    public static ProductEntityBuilder newBuilder() {
        return new ProductEntityBuilder();
    }

    public static class ProductEntityBuilder {
        private int myId;
        private String myName;
        private String myDescription;
        private int myPrice;

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

        public ProductEntity build() {
            return new ProductEntity(this);
        }
    }
}
