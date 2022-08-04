package net.consolejs;

import org.eclipse.microprofile.openapi.annotations.media.Schema;

import java.util.ArrayList;
import java.util.List;

public class PaginationEntity<T> {
    @Schema(hidden = true)
    private final List<T> myItems;
    @Schema(hidden = true)
    private final int myPage;
    @Schema(hidden = true)
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
