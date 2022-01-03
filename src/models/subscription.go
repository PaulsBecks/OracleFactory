package models

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cloudflare/cfssl/log"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Subscription struct {
	gorm.Model
	OutboundOracleID     uint
	OutboundOracle       OutboundOracle
	Topic                string
	Filter               string
	Callback             string
	SmartContractAddress string
}

func GetSubsriptionsMatchingTopic(topic string) []Subscription {
	db := utils.DBConnection()
	var subscriptions []Subscription
	db.Preload(clause.Associations).Find(&subscriptions, fmt.Sprintf("'%s' LIKE topic || '%%' ", topic))
	return subscriptions
}

func (s *Subscription) GetOutboundOracle() *OutboundOracle {
	db := utils.DBConnection()
	var outboundOracle OutboundOracle
	db.Find(&outboundOracle, s.OutboundOracleID)
	return &outboundOracle
}

func (s *Subscription) Publish(eventData map[string]interface{}) {
	if !s.FilterRulesApply(eventData) {
		s.GetOutboundOracle().GetBlockchainConnector().CreateTransaction(s.SmartContractAddress, s.Callback, eventData)
	}
}

func (s *Subscription) FilterRulesApply(event map[string]interface{}) bool {
	if s.Filter == "" {
<<<<<<< HEAD
		return true
=======
		return false
>>>>>>> 5543f6e937756cab6804aa5e54bc2e5c593f34be
	}
	for _, filter := range strings.Split(s.Filter, ";") {
		nameOperatorValue := strings.Split(filter, " ")
		if len(nameOperatorValue) < 3 {
			log.Info(fmt.Sprintf("Subsriber %d has bad filter %s set", s.ID, s.Filter))
			return true
		}
		name := nameOperatorValue[0]
		operator := nameOperatorValue[1]
		value := nameOperatorValue[2]
		eventValue := event[name]
		switch operator {
		case "=":
			if value != eventValue {
				return true
			}
		case "<":
			fValue, err := strconv.ParseFloat(value, 64)
			fEventValue, err2 := getFloat(eventValue)
			if err != nil || err2 != nil || fValue >= fEventValue {
				return true
			}
		}
	}
	return false
}

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		return math.NaN(), fmt.Errorf("Can't convert %v to float64", unk)
	}
}
