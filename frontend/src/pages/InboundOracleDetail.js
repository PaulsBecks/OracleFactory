import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useInboundOracle from "../hooks/useInboundOracle";
import { Button, Icon, Segment, Table, TableCell } from "semantic-ui-react";
import {
  ExampleRequest,
  InboundOracleForm,
  OracleOnOffRibbon,
  StartStopButton,
} from "../components";
import FilterForm from "../components/FilterForm";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import config from "../config";

export default function InboundOracleDetail({}) {
  const { inboundOracleID } = useParams();
  const [
    inboundOracle,
    updateInboundOracle,
    loading,
    startInboundOracle,
    stopInboundOracle,
  ] = useInboundOracle(inboundOracleID);

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
      <div
        style={{
          display: "flex",
          flexWrap: "wrap",
        }}
      >
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <Segment>
            <OracleOnOffRibbon
              oracleStarted={
                inboundOracle.Oracle.Status === ORACLE_STATUS_STARTED
              }
            />
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
              <b>Webhook:</b> {config.BASE_URL}/inboundOracles/
              {inboundOracle.ID}/events
            </p>
            <StartStopButton
              loading={loading}
              oracleStarted={
                inboundOracle.Oracle.Status === ORACLE_STATUS_STARTED
              }
              stopOracle={stopInboundOracle}
              startOracle={startInboundOracle}
            />
          </Segment>
        </div>
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <ExampleRequest
            eventParameters={
              inboundOracle.SmartContractPublisher.ListenerPublisher
                .EventParameters
            }
          />
        </div>
      </div>
      <br />
      <FilterForm
        listenerPublisherID={
          inboundOracle.SmartContractPublisher.ListenerPublisher.ID
        }
        oracleID={inboundOracle.Oracle.ID}
      />
      <br />
      <div>
        <h2>Events</h2>
        {inboundOracle.Oracle.Events.length > 0 ? (
          <Table unstackable>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>At</Table.HeaderCell>
                <Table.HeaderCell>Content</Table.HeaderCell>
                <Table.HeaderCell>Success</Table.HeaderCell>
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
                  <TableCell>
                    {inboundEvent.Success ? (
                      <Icon name="check" />
                    ) : (
                      <Icon name="times" />
                    )}
                  </TableCell>
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
