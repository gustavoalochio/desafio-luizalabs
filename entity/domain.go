package entity

/*
Example: google.com

Domain[

	Token: com,
	SubDomains: [
		Domain[
			Token: google
			SubDomains: nil
		]
	]

]
*/
type Domain struct {
	Token      string
	SubDomains map[string]*Domain
	Count      int
}

func NewDomain(token string, count int) *Domain {
	return &Domain{
		Token:      token,
		Count:      count,
		SubDomains: map[string]*Domain{},
	}
}
