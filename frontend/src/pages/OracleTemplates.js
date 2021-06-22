import useOutboundOracleTemplates from "../hooks/useOutboundOracleTemplates";
import useInboundOraclesTemplates from "../hooks/useInboundOraclesTemplates";
import { useHistory, Link } from "react-router-dom";
import { Button, Table } from "semantic-ui-react";

function Home() {
  const history = useHistory();
  const [outboundOracleTemplates] = useOutboundOracleTemplates();
  const [inboundOracleTemplates] = useInboundOraclesTemplates();
  return (
    <div>
      <h1>Home</h1>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Template"
          as={Link}
          to="/oracleTemplates/create"
        />
      </div>
      <br />
      <div>
        <h2>Outbound Oracle Templates</h2>
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
          }}
        >
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>Event Name</Table.HeaderCell>
                <Table.HeaderCell>Contract Address</Table.HeaderCell>
                <Table.HeaderCell></Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {outboundOracleTemplates.map((outboundOracleTemplate) => {
                console.log(outboundOracleTemplate);
                return (
                  <Table.Row>
                    <Table.Cell>
                      {outboundOracleTemplate.EventName || ""}
                    </Table.Cell>
                    <Table.Cell>{outboundOracleTemplate.Address}</Table.Cell>
                    <Table.Cell>
                      <Button
                        as={Link}
                        to={
                          "/outboundOracleTemplates/" +
                          outboundOracleTemplate.ID
                        }
                        content="Select"
                      />
                    </Table.Cell>
                  </Table.Row>
                );
              })}
            </Table.Body>
          </Table>
        </div>
      </div>
      <br />
      <div>
        <h2>Inbound Oracle Templates</h2>
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
          }}
        >
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>Contract Name</Table.HeaderCell>
                <Table.HeaderCell>Contract Address</Table.HeaderCell>
                <Table.HeaderCell></Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {inboundOracleTemplates.map((inboundOracleTemplate) => (
                <Table.Row>
                  <Table.Cell>
                    {inboundOracleTemplate.ContractName || ""}
                  </Table.Cell>
                  <Table.Cell>
                    {inboundOracleTemplate.ContractAddress}
                  </Table.Cell>
                  <Table.Cell>
                    <Button
                      as={Link}
                      to={"/inboundOracleTemplates/" + inboundOracleTemplate.ID}
                      content="Select"
                    />
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        </div>
      </div>
    </div>
  );
}

export default Home;
