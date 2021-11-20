import { Tab } from "semantic-ui-react";
import SmartContractCreateDetail from "../components/SmartContractCreateDetail";
import ProviderCreate from "../components/ProviderCreate";

export default function PublishSubscriberCreate() {
  return (
    <div>
      <h1>Create Listeners and Publishers</h1>
      <Tab
        menu={{ secondary: true, pointing: true }}
        panes={[
          {
            menuItem: "Provider",
            render: () => (
              <Tab.Pane attached={false}>
                <ProviderCreate listener />
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
                <SmartContractCreateDetail pubSub key="Inbound" />
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
