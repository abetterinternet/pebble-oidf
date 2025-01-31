package openidfederation

import (
	"fmt"
	"net/url"
)

// EntityIdentifier identifies an entity in an OpenID Federation.
type EntityIdentifier struct {
	url url.URL
}

// NewEntityIdentifier returns an EntityIdentifier if it the provided identifier is a valid OpenID
// Federation entity identifier.
func NewEntityIdentifier(identifier string) (EntityIdentifier, error) {
	entityURL, err := url.Parse(identifier)
	if err != nil {
		return EntityIdentifier{}, fmt.Errorf("identifier '%s' is not a valid OIDF entity identifier: %w", identifier, err)
	}

	if entityURL.Scheme != "https" {
		return EntityIdentifier{}, fmt.Errorf("identifier '%s' is not a valid OIDF entity identifier: scheme must be https", identifier)
	}

	return EntityIdentifier{url: *entityURL}, nil
}
