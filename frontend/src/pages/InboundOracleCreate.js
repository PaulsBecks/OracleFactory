import React, { useState } from "react";
import { useParams } from "react-router";
import postData from "../services/postData";
import { InboundOracleForm } from "../components";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function InboundOracleCreate() {
  const history = useHistory();
  const { inboundOracleTemplateID } = useParams();
  const [inboundOracle, setInboundOracle] = useState({ Name: "", URI: "" });
  const [loading, setLoading] = useState(false);

  return (
    <div>
      <h1>Create Inbound Oracle</h1>
      <InboundOracleForm
        inboundOracle={inboundOracle}
        setInboundOracle={setInboundOracle}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={"/inboundOracleTemplates/" + inboundOracleTemplateID}
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(
            `/inboundOracleTemplates/${inboundOracleTemplateID}/inboundOracles`,
            {
              ...inboundOracle,
            }
          );
          setLoading(false);
          history.push("/inboundOracleTemplates/" + inboundOracleTemplateID);
        }}
      />
    </div>
  );
}
