import React, { useState } from "react";
import postData from "../services/postData";
import { InboundOracleTemplateForm } from "../components";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function InboundOracleTemplateCreate() {
  const history = useHistory();
  const [inboundOracleTemplate, setInboundOracleTemplate] = useState({
    BlockchainAddress: "",
    BlockchainName: "Ethereum",
    ContractAddress: "",
    inboundOracleTemplates: [],
  });

  console.log(inboundOracleTemplate);

  const [loading, setLoading] = useState(false);

  return (
    <div>
      <h1>Create Inbound Oracle</h1>
      <InboundOracleTemplateForm
        inboundOracleTemplate={inboundOracleTemplate}
        setInboundOracleTemplate={setInboundOracleTemplate}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={"/"}
      />
      <Button
        loading={loading}
        basic
        positive
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
              });
            } else {
              result = await postData(`/outboundOracleTemplates`, {
                BlockchainAddress: inboundOracleTemplate.BlockchainAddress,
                BlockchainName: inboundOracleTemplate.BlockchainName,
                ContractAddress: inboundOracleTemplate.ContractAddress,
                EventName: element.ContractName,
              });
            }
            const oracleTemplate =
              element.Type === "function"
                ? result.inboundOracleTemplate
                : result.outboundOracleTemplate;
            console.log(oracleTemplate);
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
    </div>
  );
}
