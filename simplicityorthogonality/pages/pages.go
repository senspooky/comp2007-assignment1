package pages

type PagedData[T any] interface {
	Page(int) []T
}

func New[T any](pageSize int, data []T) PagedData[T] {
	return pagedData[T]{
		pageSize: pageSize,
		data:     data,
	}
}

type pagedData[T any] struct {
	pageSize int
	data     []T
}

func (p pagedData[T]) Page(num int) []T {
	i := num * p.pageSize
	j := i + p.pageSize
	if i > len(p.data) || j > len(p.data) {
		return nil
	}
	return p.data[i:j]
}
