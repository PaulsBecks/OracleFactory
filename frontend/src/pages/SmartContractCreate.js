import { Tab } from "semantic-ui-react";
import SmartContractCreateDetail from "../components/SmartContractCreateDetail";
import WebServiceListenerCreate from "../components/WebServiceListenerCreate";
import WebServicePublisherCreate from "../components/WebServicePublisherCreate";

export default function PublishSubscriberCreate() {
  return (
    <div>
      <h1>Create Listeners and Publishers</h1>
      <Tab
        menu={{ secondary: true, pointing: true }}
        panes={[
          {
            menuItem: "Web Service Listener",
            render: () => (
              <Tab.Pane attached={false}>
                <WebServiceListenerCreate listener />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Web Service Publisher",
            render: () => (
              <Tab.Pane attached={false}>
                <WebServicePublisherCreate publisher />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Smart Contract Listener",
            render: () => (
              <Tab.Pane attached={false}>
                <SmartContractCreateDetail outbound key="Outbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Smart Contract Publisher",
            render: () => (
              <Tab.Pane attached={false}>
                <SmartContractCreateDetail inbound key="Inbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "From ABI",
            render: () => (
              <Tab.Pane attached={false}>
                <SmartContractCreateDetail fromABI key="ABI" />
              </Tab.Pane>
            ),
          },
        ]}
      />
    </div>
  );
}
