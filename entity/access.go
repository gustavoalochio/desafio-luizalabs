package entity

type Access struct {
	Url   string
	Count int
}

func NewAccess(url string, count int) Access {
	return Access{
		Url:   url,
		Count: count,
	}
}
