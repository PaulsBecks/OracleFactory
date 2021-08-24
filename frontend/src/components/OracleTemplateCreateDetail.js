import React, { useState } from "react";
import postData from "../services/postData";
import { InboundOracleTemplateForm } from ".";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function OracleTemplateCreate({ fromABI, inbound, outbound }) {
  const history = useHistory();
  const [inboundOracleTemplate, setInboundOracleTemplate] = useState({
    Private: true,
    BlockchainName: "Ethereum",
    ContractAddress: "",
    inboundOracleTemplates: fromABI
      ? []
      : [{ ContractName: "", Type: "function", inputs: [] }],
  });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <InboundOracleTemplateForm
        inboundOracleTemplate={inboundOracleTemplate}
        setInboundOracleTemplate={setInboundOracleTemplate}
        fromABI={fromABI}
        inbound={inbound}
        outbound={outbound}
      />
      <br />
      <Button
        loading={loading}
        basic
        positive
        fluid
        floated="right"
        content="Create"
        onClick={async () => {
          setLoading(true);
          for (const element of inboundOracleTemplate.inboundOracleTemplates) {
            let result;
            if (element.Type === "function") {
              result = await postData(`/inboundOracleTemplates`, {
                BlockchainAddress: inboundOracleTemplate.BlockchainAddress,
                BlockchainName: inboundOracleTemplate.BlockchainName,
                ContractAddress: inboundOracleTemplate.ContractAddress,
                ContractName: element.ContractName,
                Private: inboundOracleTemplate.Private,
              });
            } else {
              result = await postData(`/outboundOracleTemplates`, {
                BlockchainAddress: inboundOracleTemplate.BlockchainAddress,
                BlockchainName: inboundOracleTemplate.BlockchainName,
                ContractAddress: inboundOracleTemplate.ContractAddress,
                EventName: element.ContractName,
                Private: inboundOracleTemplate.Private,
              });
            }
            const oracleTemplate =
              element.Type === "function"
                ? result.inboundOracleTemplate
                : result.outboundOracleTemplate;
            for (const input of element.inputs) {
              if (element.Type === "function") {
                await postData(
                  `/inboundOracleTemplates/${oracleTemplate.ID}/eventParameters`,
                  input
                );
              } else {
                await postData(
                  `/outboundOracleTemplates/${oracleTemplate.ID}/eventParameters`,
                  input
                );
              }
            }
          }
          setLoading(false);
          history.push("/oracleTemplates");
        }}
      />
      {/*<Button
        loading={loading}
        basic
        negative
        floated="right"
        content="Cancel"
        as={Link}
        to={"/"}
      />*/}
      <br />
      <br />
    </div>
  );
}
