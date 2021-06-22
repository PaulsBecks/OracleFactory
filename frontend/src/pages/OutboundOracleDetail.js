import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Button, Table } from "semantic-ui-react";
import useOutboundOracle from "../hooks/useOutboundOracle";
import { OutboundOracleForm } from "../components";

export default function OutboundOracleDetail({}) {
  const { outboundOracleID } = useParams();
  const [outboundOracle, updateOutboundOracle, loading] =
    useOutboundOracle(outboundOracleID);

  const [localOutboundOracles, setLocalOutboundOracle] = useState({});

  useEffect(() => {
    setLocalOutboundOracle(outboundOracle);
  }, [outboundOracle]);
  if (!outboundOracle) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Outbound Oracle</h1>
      <div>
        <div style={{ maxWidth: "400px" }}>
          <OutboundOracleForm
            outboundOracle={localOutboundOracles}
            setOutboundOracle={setLocalOutboundOracle}
          />
          {JSON.stringify(localOutboundOracles) !==
            JSON.stringify(outboundOracle) && (
            <>
              <br />
              <Button
                loading={loading}
                content="Save"
                positive
                basic
                onClick={async () => updateOutboundOracle(localOutboundOracles)}
              />
            </>
          )}
        </div>
        <br />
      </div>
      <div>
        <h2>Events</h2>
        {outboundOracle.OutboundEvents.length > 0 ? (
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>At</Table.HeaderCell>
                <Table.HeaderCell>Content</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {outboundOracle.OutboundEvents.map((outboundEvent) => (
                <Table.Row>
                  <Table.Cell>{outboundEvent.ID}</Table.Cell>
                  <Table.Cell>
                    {new Date(outboundEvent.CreatedAt).toLocaleString()}
                  </Table.Cell>
                  <Table.Cell>
                    {outboundEvent.EventValues.map((value) => (
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
