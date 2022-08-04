package net.consolejs.entities;

import java.util.ArrayList;
import java.util.List;

public class PaginationEntity<T> {
    private final List<T> myItems;
    private final int myPage;
    private final int myLastPage;

    private PaginationEntity(PaginationEntityBuilder builder) {
        myItems = builder.myItems;
        myPage = builder.myPage;
        myLastPage = builder.myLastPage;
    }

    public List<T> getItems() {
        return myItems;
    }

    public int getPage() {
        return myPage;
    }

    public int getLastPage() {
        return myLastPage;
    }

    public static <I> PaginationEntityBuilder<I> newBuilder() {
        return new PaginationEntityBuilder<>();
    }

    public static class PaginationEntityBuilder<I> {
        private List<I> myItems;
        private int myPage;
        private int myLastPage;

        private PaginationEntityBuilder() {

        }

        public PaginationEntityBuilder<I> withItems(List<I> items) {
            myItems = new ArrayList<>(items);
            return this;
        }

        public PaginationEntityBuilder<I> withPage(int page) {
            myPage = page;
            return this;
        }

        public PaginationEntityBuilder<I> withLastPage(int lastPage) {
            myLastPage = lastPage;
            return this;
        }

        public PaginationEntity<I> build() {
            return new PaginationEntity<I>(this);
        }
    }
}
