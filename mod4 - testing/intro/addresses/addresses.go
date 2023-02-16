package addresses

import "strings"

// AddressType check if address has a valid type
func AddressType(address string) string {
	validTypes := []string{"Street", "Avenue", "Road"}

	firstWord := strings.ToLower(strings.Split(address, " ")[0])

	for _, t := range validTypes {
		if strings.ToLower(t) == firstWord {
			return t
		}
	}

	return "Invalid"
}
