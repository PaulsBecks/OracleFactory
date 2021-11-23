package forms

import (
	"fmt"
	"regexp"

	"github.com/PaulsBecks/OracleFactory/src/models"
)

type SubscriptionBody struct {
	Name                 string
	topic                string
	filter               string
	callback             string
	smartContractAddress string
}

// TODO: create real validation
func (s *SubscriptionBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, s.topic)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}

func (s *SubscriptionBody) CreateSubscription(outboundOracle *models.OutboundOracle) *models.Subscription {
	return outboundOracle.CreateSubscription(s.topic, s.filter, s.callback, s.smartContractAddress)
}
