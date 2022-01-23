package forms

import (
	"fmt"
	"regexp"
)

type InboundSubscriptionBody struct {
	Subscription            SubscriptionBody
	SmartContractConsumerID uint
	WebServiceProviderID    uint
}

// TODO: create real validation
func (o *InboundSubscriptionBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Subscription.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
