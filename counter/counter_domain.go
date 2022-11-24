package counter

import (
	"fmt"
	"github.com/gustavoalochio/desafio-luizalabs/entity"
	"sort"
	"strings"
)

type CounterDomain interface {
	CountDomains(data []entity.Access) map[string]*entity.Domain
	PrintDomainsSorted(domains map[string]*entity.Domain)
}

type counterDomain struct{}

func NewCounterDomain() CounterDomain {
	return &counterDomain{}
}

// receive a access list and split the access struct in domains struct
/*
Example:

loja.google.com
yahoo.com

Each level in this tree has a hashmap
				  com
				/  	  \
               /       \
            google      yahoo
			/
           /
         loja
*/
func (d *counterDomain) CountDomains(data []entity.Access) map[string]*entity.Domain {

	domains := map[string]*entity.Domain{}

	for _, access := range data {

		tokens := strings.Split(access.Url, ".")

		subDomains := domains

		// loja.google.com -> tokens = [loja, google, com]
		for len(tokens) > 0 {
			lastToken := tokens[len(tokens)-1]

			var domain *entity.Domain

			_, ok := subDomains[lastToken]
			if !ok {
				domain = entity.NewDomain(lastToken, access.Count)
				subDomains[lastToken] = domain
			} else {
				domain = subDomains[lastToken]
				domain.Count += access.Count
			}

			subDomains = domain.SubDomains
			tokens = tokens[:len(tokens)-1]
		}
	}

	return domains
}
func (d *counterDomain) PrintDomainsSorted(domains map[string]*entity.Domain) {
	d.printDomainsSorted(domains, nil)
}

func (d *counterDomain) printDomainsSorted(domains map[string]*entity.Domain, tokens []string) {

	if domains == nil {
		return
	}

	keys := make([]string, 0, len(domains))

	for k := range domains {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		tokensUse := append([]string{domains[k].Token}, tokens...)
		fmt.Println(domains[k].Count, strings.Join(tokensUse, "."))
		d.printDomainsSorted(domains[k].SubDomains, tokensUse)
	}
}
