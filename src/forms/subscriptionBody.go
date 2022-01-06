package forms

import (
	"fmt"
	"regexp"

	"github.com/PaulsBecks/OracleFactory/src/models"
)

type SubscriptionBody struct {
	Token                string
	DeferredChoiceID     string
	Topic                string
	Filter               string
	Callback             string
	SmartContractAddress string
}

// TODO: create real validation
func (s *SubscriptionBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, s.Topic)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}

func (s *SubscriptionBody) CreateSubscription(outboundOracle *models.OutboundOracle) *models.Subscription {
	return outboundOracle.CreateSubscription(s.Topic, s.Filter, s.Callback, s.SmartContractAddress)
}

type UnsubscriptionBody struct {
	Token string
	Topic string
}

// TODO: create real validation
func (s *UnsubscriptionBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, s.Topic)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}

func (u *UnsubscriptionBody) DeleteSubscription(outboundOracle *models.OutboundOracle) {
	outboundOracle.DeleteSubscription(u.Topic)
}
