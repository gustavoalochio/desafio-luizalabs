package entity

/*
Example:
10,google.com
20,yahoo.com

List
[
	Access[
		Url: google.com,
		Count: 10
	],
	Access[
		Url: yahoo.com,
		Count: 20
	],
]
*/

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
