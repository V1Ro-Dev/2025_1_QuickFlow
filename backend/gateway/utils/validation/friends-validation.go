package validation

const (
	TypeOutComing   = "outcoming"
	TypeInComing    = "incoming"
	TypeAll         = "all"
	TypeNewIncoming = "new_incoming"
)

var acceptedReqTypesSet = map[string]struct{}{
	TypeOutComing:   {},
	TypeInComing:    {},
	TypeAll:         {},
	TypeNewIncoming: {},
}

func ValidateFriendReqType(reqType string) bool {
	if _, ok := acceptedReqTypesSet[reqType]; !ok {
		return false
	}

	return true
}
