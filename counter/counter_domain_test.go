package counter

import (
	"github.com/gustavoalochio/desafio-luizalabs/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (

	// 100 loja.google.com
	// 50 yahoo.com

	domains = map[string]*entity.Domain{
		"com": &entity.Domain{
			Token: "com",
			SubDomains: map[string]*entity.Domain{
				"google": &entity.Domain{
					Token: "google",
					SubDomains: map[string]*entity.Domain{
						"loja": &entity.Domain{
							Token:      "loja",
							SubDomains: map[string]*entity.Domain{},
							Count:      100,
						},
					},
					Count: 100,
				},
				"yahoo": &entity.Domain{
					Token:      "yahoo",
					SubDomains: map[string]*entity.Domain{},
					Count:      50,
				},
			},
			Count: 150,
		},
	}

	accessList = []entity.Access{
		entity.Access{
			Url:   "loja.google.com",
			Count: 100,
		},
		entity.Access{
			Url:   "yahoo.com",
			Count: 50,
		},
	}
)

func TestCounterDomain_CountDomains(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		//// arrange
		counterDomain := NewCounterDomain()

		//// action
		domainList := counterDomain.CountDomains(accessList)
		counterDomain.PrintDomainsSorted(domainList)

		//// assert
		assert.NotNil(t, domainList)
		assert.Equal(t, domains, domainList)
	})

	t.Run("error", func(t *testing.T) {

	})
}
