import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Button, Segment, Table } from "semantic-ui-react";
import useOutboundOracle from "../hooks/useOutboundOracle";
import {
  ExampleRequest,
  OracleOnOffRibbon,
  OutboundOracleForm,
  StartStopButton,
} from "../components";
import FilterForm from "../components/FilterForm";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import SmartContractListenerCard from "../components/SmartContractListenerCard";

export default function OutboundOracleDetail({}) {
  const { outboundOracleID } = useParams();
  const [
    outboundOracle,
    updateOutboundOracle,
    loading,
    startOutboundOracle,
    stopOutboundOracle,
  ] = useOutboundOracle(outboundOracleID);

  const [localOutboundOracles, setLocalOutboundOracle] = useState({});

  useEffect(() => {
    setLocalOutboundOracle(outboundOracle);
  }, [outboundOracle]);
  if (!outboundOracle) {
    return "Loading...";
  }

  let oracleStarted = outboundOracle.Oracle.Status === ORACLE_STATUS_STARTED;
  return (
    <div>
      <h1>Outbound Oracle</h1>
      <div style={{ display: "flex", flexWrap: "wrap" }}>
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <Segment>
            <OracleOnOffRibbon oracleStarted={oracleStarted} />
            <OutboundOracleForm
              outboundOracle={localOutboundOracles}
              setOutboundOracle={setLocalOutboundOracle}
            />
            <br />
            <label>
              <b>Forward to:</b>
            </label>
            <p>{outboundOracle.WebServicePublisher.Url}</p>
            {JSON.stringify(localOutboundOracles) !==
              JSON.stringify(outboundOracle) && (
              <>
                <br />
                <Button
                  loading={loading}
                  content="Save"
                  positive
                  basic
                  onClick={async () =>
                    updateOutboundOracle(localOutboundOracles)
                  }
                />
                <br />
              </>
            )}
            <br />
            <StartStopButton
              loading={loading}
              oracleStarted={oracleStarted}
              stopOracle={stopOutboundOracle}
              startOracle={startOutboundOracle}
            />
          </Segment>
        </div>
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <ExampleRequest
            eventParameters={
              outboundOracle.SmartContractListener.ListenerPublisher
                .EventParameters
            }
          />
        </div>
      </div>
      <br />
      <FilterForm
        oracleID={outboundOracle.OracleID}
        listenerPublisherID={
          outboundOracle.SmartContractListener.ListenerPublisher.ID
        }
      />
      <br />
      <div>
        <h2>Events</h2>
        {outboundOracle.Oracle.Events.length > 0 ? (
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>At</Table.HeaderCell>
                <Table.HeaderCell>Content</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {outboundOracle.Oracle.Events.map((outboundEvent) => (
                <Table.Row>
                  <Table.Cell>{outboundEvent.ID}</Table.Cell>
                  <Table.Cell>
                    {new Date(outboundEvent.CreatedAt).toLocaleString()}
                  </Table.Cell>
                  <Table.Cell>
                    {outboundEvent.EventValues.map((value) => {
                      console.log(value);
                      return (
                        <>
                          <b>{value.EventParameter.Name}:</b>
                          {value.Value}
                          <br />
                        </>
                      );
                    })}
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
