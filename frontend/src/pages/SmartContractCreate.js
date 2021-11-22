import { Tab } from "semantic-ui-react";
import SmartContractCreateDetail from "../components/SmartContractCreateDetail";
import ProviderCreate from "../components/ProviderCreate";

export default function SmartContractCreate() {
  return (
    <div>
      <h1>Create Provider/Consumer/Events</h1>
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
            menuItem: "Consumer",
            render: () => (
              <Tab.Pane attached={false}>
                <SmartContractCreateDetail pubSub key="Inbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Blockchain Event",
            render: () => (
              <Tab.Pane attached={false}>
                <SmartContractCreateDetail outbound key="Outbound" />
              </Tab.Pane>
            ),
          },
          {
            menuItem: "Consumer/Event From ABI",
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
