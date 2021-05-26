package models

import (
	"fmt"
	"gorm.io/gorm"
)

type OutboundOracle struct {
	gorm.Model
	URI                      string `json:uri`
	OutboundOracleTemplateID uint   `json:outbound_oracle_template_id`
	OutboundOracleTemplate   OutboundOracleTemplate
	OutboundEvents           []OutboundEvent
}

func (o *OutboundOracle) CreateManifest() Manifest {
	return Manifest(`SET BLOCKCHAIN \"` + o.OutboundOracleTemplate.Blockchain + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION \"` + o.OutboundOracleTemplate.GetConnectionString() + `\";


BLOCKS (0) (CONTINUOUS) {
	LOG ENTRIES (` + o.OutboundOracleTemplate.Address + `) (` + o.OutboundOracleTemplate.EventName + `(` + o.OutboundOracleTemplate.GetEventParametersString() + `)) {
		EMIT HTTP REQUEST (\"` + o.oracleFactoryOutboundEventLink() + `\") (` + o.OutboundOracleTemplate.GetEventParameterNamesString() + `);
	}
}`)
}

func (o *OutboundOracle) oracleFactoryOutboundEventLink() string {
	return "http://oracle-factory:8080/outboundOracles/" + fmt.Sprint(o.ID) + "/events"
}
