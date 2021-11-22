import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import usePubSubOracle from "../hooks/usePubSubOracle";
import {
  Button,
  Card,
  Form,
  Icon,
  Segment,
  Table,
  TableCell,
} from "semantic-ui-react";
import {
  PubSubOracleForm,
  OracleOnOffRibbon,
  StartStopButton,
} from "../components";
import FilterForm from "../components/FilterForm";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import ConsumerCard from "../components/ConsumerCard";
import ProviderCard from "../components/ProviderCard";
import OutboundOracleCard from "../components/OutboundOracleCard";

export default function PubSubOracleDetail({}) {
  const { pubSubOracleID } = useParams();
  const [
    pubSubOracle,
    updatePubSubOracle,
    loading,
    startPubSubOracle,
    stopPubSubOracle,
  ] = usePubSubOracle(pubSubOracleID);

  const [localPubSubOracle, setLocalPubSubOracle] = useState();
  useEffect(() => {
    setLocalPubSubOracle(pubSubOracle);
  }, [pubSubOracle]);

  if (!pubSubOracle) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Pub-Sub Oracle</h1>
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
                pubSubOracle.Oracle.Status === ORACLE_STATUS_STARTED
              }
            />
            <PubSubOracleForm
              pubSubOracle={localPubSubOracle}
              setPubSubOracle={setLocalPubSubOracle}
            />
            {JSON.stringify(localPubSubOracle) !==
              JSON.stringify(pubSubOracle) && (
              <>
                <br />
                <Button
                  positive
                  basic
                  fluid
                  loading={loading}
                  content="Save"
                  onClick={() => updatePubSubOracle(localPubSubOracle)}
                />
                <br />
              </>
            )}
            <br />
            <StartStopButton
              loading={loading}
              oracleStarted={
                pubSubOracle.Oracle.Status === ORACLE_STATUS_STARTED
              }
              stopOracle={stopPubSubOracle}
              startOracle={startPubSubOracle}
            />
          </Segment>
        </div>
      </div>
      <br />
      <Card.Group>
        <ProviderCard provider={pubSubOracle.Provider} />
        <ConsumerCard consumer={pubSubOracle.Consumer} />
        <OutboundOracleCard outboundOracle={pubSubOracle.SubOracle} />
        <OutboundOracleCard outboundOracle={pubSubOracle.UnsubOracle} />
      </Card.Group>
      <br />
      <FilterForm
        listenerPublisherID={pubSubOracle.Consumer.ListenerPublisher.ID}
        oracleID={pubSubOracle.Oracle.ID}
      />
      <br />
      <div>
        <h2>Events</h2>
        {pubSubOracle.Oracle.Events.length > 0 ? (
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
              {pubSubOracle.Oracle.Events.map((pubSubEvent) => (
                <Table.Row>
                  <Table.Cell>{pubSubEvent.ID}</Table.Cell>
                  <Table.Cell>
                    {new Date(pubSubEvent.CreatedAt).toLocaleString()}
                  </Table.Cell>
                  <Table.Cell>
                    {pubSubEvent.EventValues.map((value) => (
                      <>
                        <b>{value.EventParameter.Name}:</b>
                        {value.Value}
                        <br />
                      </>
                    ))}
                  </Table.Cell>
                  <TableCell>
                    {pubSubEvent.Success ? (
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
