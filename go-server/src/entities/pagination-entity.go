package entities

type PaginationEntity[T any] struct {
	Items    []T `json:"items"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
}
