import React, { useState } from "react";
import { useParams } from "react-router";
import postData from "../services/postData";
import { OutboundOracleForm } from "../components";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function OutboundOracleCreate() {
  const history = useHistory();
  const { outboundOracleTemplateID } = useParams();
  const [outboundOracle, setOutboundOracle] = useState({
    Oracle: { Name: "" },
    URI: "",
  });
  const [loading, setLoading] = useState(false);

  return (
    <div>
      <h1>Create Outbound Oracle</h1>
      <OutboundOracleForm
        outboundOracle={outboundOracle}
        setOutboundOracle={setOutboundOracle}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={"/outboundOracleTemplates/" + outboundOracleTemplateID}
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(
            `/outboundOracleTemplates/${outboundOracleTemplateID}/outboundOracles`,
            {
              ...outboundOracle,
            }
          );
          setLoading(false);
          history.push("/outboundOracleTemplates/" + outboundOracleTemplateID);
        }}
      />
    </div>
  );
}
