package frantic

import (
	"fmt"
	"strings"
)

type Identity struct {
	name       string
	suffix     string
	isOfficial bool
}

func New(name string) (Identity, error) {
	// name must have two parts, separated by a - or _ character (e.g. "frantic-go")
	// the first part is the name, the second part is the suffix
	// if there is no - or _ character, return an error

	err := validateName(name)
	if err != nil {
		return Identity{}, err
	}

	name = sanitizeName(name)

	verified := false
	validErr := ValidateIdentityOrigin(name)
	if validErr != nil {
		verified = true
	}

	parts := strings.Split(name, "-")
	if len(parts) != 2 {
		return Identity{}, fmt.Errorf("name must have only two parts, separated by a - or _ character")
	}

	rtn := Identity{}
	rtn.name = strings.ToLower(name)
	rtn.suffix = strings.ToLower(parts[1])
	rtn.isOfficial = verified
	return rtn, nil

}

func (i *Identity) String() string {
	return i.name
}

func (i *Identity) Name() string {
	return i.String()
}

func (i *Identity) Prefix() string {
	return strings.Split(i.name, "-")[0]
}
func (i *Identity) Suffix() string {
	return i.suffix
}

func (i *Identity) IsOfficial() error {
	return ValidateIdentityOrigin(i.name)
}

func ValidateIdentityOrigin(name string) error {
	validOrigins := []string{
		"deomon-pear",
		"cosmic-orange",
		"trnsl8r_connect",
		"trnsl8r_service",
		"frantic-core",
		"frantic-pear",
		"frantic-orange",
		"frantic-aegis",
		"frantic-aliquid",
		"frantic-agentis",
		"frantic-webcore",
		"frantic-template",
	}

	err := validateName(name)
	if err != nil {
		return err
	}

	name = sanitizeName(name)

	found := false
	for _, origin := range validOrigins {
		if strings.EqualFold(name, origin) {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("invalid this is not a valid origin")
	}
	return nil
}

func sanitizeName(name string) string {
	name = strings.ReplaceAll(name, "_", "-")
	return name
}

func validateName(name string) error {
	if len(name) < 5 {
		return fmt.Errorf("name must contain a - or _ character")
	}
	if !strings.Contains(name, "-") || !strings.Contains(name, "_") {
		return fmt.Errorf("name must contain a - or _ character")
	}
	return nil
}
