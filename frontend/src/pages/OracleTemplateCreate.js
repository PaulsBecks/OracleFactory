import { Link } from "react-router-dom";
import { Button, Tab } from "semantic-ui-react";
import OracleTemplateCreateDetail from "../components/OracleTemplateCreateDetail";

export default function OracleTemplateCreate() {
  return (
    <div>
      <h1>Create Oracle Template</h1>
      <Tab
        menu={{ secondary: true, pointing: true }}
        panes={[
          {
            menuItem: "Outbound Oracle",
            render: () => (
              <Tab.Pane attached={false}>
                <OracleTemplateCreateDetail outbound key="Outbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Inbound Oracle",
            render: () => (
              <Tab.Pane attached={false}>
                <OracleTemplateCreateDetail inbound key="Inbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "From ABI",
            render: () => (
              <Tab.Pane attached={false}>
                <OracleTemplateCreateDetail fromABI key="ABI" />
              </Tab.Pane>
            ),
          },
        ]}
      />
    </div>
  );
}
