package forms

import (
	"fmt"
	"regexp"
)

type OutboundSubscriptionPostBody struct {
	Subscription            SubscriptionBody
	WebServiceConsumerID    uint
	SmartContractProviderID uint
}

// TODO: create real validation
func (o *OutboundSubscriptionPostBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Subscription.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
