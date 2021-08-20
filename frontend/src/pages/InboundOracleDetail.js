import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useInboundOracle from "../hooks/useInboundOracle";
import { Button, Form, Table } from "semantic-ui-react";
import { InboundOracleForm } from "../components";
import FilterForm from "../components/FilterForm";

export default function InboundOracleDetail({}) {
  const { inboundOracleID } = useParams();
  const [inboundOracle, updateInboundOracle, loading] =
    useInboundOracle(inboundOracleID);

  const [localInboundOracle, setLocalInboundOracle] = useState();

  useEffect(() => {
    setLocalInboundOracle(inboundOracle);
  }, [inboundOracle]);

  const createFilter = () => {
    console.log(createFilter);
  };

  if (!inboundOracle) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Inbound Oracle</h1>
      <div>
        <InboundOracleForm
          inboundOracle={localInboundOracle}
          setInboundOracle={setLocalInboundOracle}
        />
        {JSON.stringify(localInboundOracle) !==
          JSON.stringify(inboundOracle) && (
          <>
            <br />
            <Button
              positive
              basic
              loading={loading}
              content="Save"
              onClick={() => updateInboundOracle(localInboundOracle)}
            />
            <br />
          </>
        )}
        <br />
        <p>
          <b>Webhook:</b> http://localhost:8080/inboundOracle/
          {inboundOracle.ID}/events
        </p>
      </div>
      <br />
      <FilterForm
        oracleTemplateID={inboundOracle.InboundOracleTemplate.OracleTemplate.ID}
        oracleID={inboundOracle.Oracle.ID}
      />
      <br />
      <div>
        <h2>Events</h2>
        {inboundOracle.Oracle.Events.length > 0 ? (
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>At</Table.HeaderCell>
                <Table.HeaderCell>Content</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {inboundOracle.Oracle.Events.map((inboundEvent) => (
                <Table.Row>
                  <Table.Cell>{inboundEvent.ID}</Table.Cell>
                  <Table.Cell>
                    {new Date(inboundEvent.CreatedAt).toLocaleString()}
                  </Table.Cell>
                  <Table.Cell>
                    {inboundEvent.EventValues.map((value) => (
                      <>
                        <b>{value.EventParameter.Name}:</b>
                        {value.Value}
                        <br />
                      </>
                    ))}
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        ) : (
          <div>No events registered so far</div>
        )}
      </div>
    </div>
  );
}
