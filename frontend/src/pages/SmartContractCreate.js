import { Tab } from "semantic-ui-react";
import SmartContractCreateDetail from "../components/SmartContractCreateDetail";
import WebServiceProviderCreate from "../components/WebServiceProviderCreate";
import WebServiceConsumerCreate from "../components/WebServiceConsumerCreate";

export default function PublishSubscriberCreate() {
  return (
    <div>
      <h1>Create Providers and Consumers</h1>
      <Tab
        menu={{ secondary: true, pointing: true }}
        panes={[
          {
            menuItem: "Web Service Provider",
            render: () => (
              <Tab.Pane attached={false}>
                <WebServiceProviderCreate provider />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Web Service Consumer",
            render: () => (
              <Tab.Pane attached={false}>
                <WebServiceConsumerCreate consumer />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Smart Contract Provider",
            render: () => (
              <Tab.Pane attached={false}>
                <SmartContractCreateDetail outbound key="Outbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Smart Contract Consumer",
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
