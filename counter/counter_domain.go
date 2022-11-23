package counter

import (
	"github.com/gustavoalochio/desafio-luizalabs/entity"
	"strings"
)

type CounterDomain interface {
	CountDomains(data []entity.Access) map[string]int
}

type counterDomain struct{}

func NewCounterDomain() CounterDomain {
	return &counterDomain{}
}

func (d *counterDomain) CountDomains(data []entity.Access) map[string]int {

	m := map[string]int{}

	for _, access := range data {

		domains := split(access.Url)

		for _, domain := range domains {

			_, ok := m[domain]
			if !ok {
				m[domain] = access.Count
				continue
			}

			m[domain] += access.Count

		}
	}

	return m

}

// loja.google.com -> [loja.google.com, google.com, com ]
func split(url string) []string {

	urlTokens := strings.Split(url, ".") // [loja, google, com]

	var domains []string

	for i := 0; i < len(urlTokens); i++ { // 2

		var tokens []string

		for j := i; j < len(urlTokens); j++ { // 2
			tokens = append(tokens, urlTokens[j]) // [com]
		}

		domains = append(domains, strings.Join(tokens, ".")) // [loja.google.com, google.com, com]
	}

	return domains

}
