package utils

func Iterate(from int64, to int64) []int64 {
	var items []int64
	for i := from; i < to; i++ {
		items = append(items, i)
	}
	return items
}
