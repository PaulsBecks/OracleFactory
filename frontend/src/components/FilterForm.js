import { Button, Form, Segment, Table, TableHeader } from "semantic-ui-react";
import useFilters from "../hooks/useFilters";
import useParameterFilters from "../hooks/useParameterFilters";
import useParameters from "../hooks/useEventParameters";
import { useState } from "react";

const createParameterFilterSceleton = (subscriptionID) => ({
  EventParameterID: 0,
  FilterID: 0,
  SubscriptionID: subscriptionID,
  Scheme: "",
});

export default function FilterForm({ providerConsumerID, subscriptionID }) {
  const [parameterFilters, createParameterFilter, deleteParameterFilter] =
    useParameterFilters(subscriptionID);
  const [parameters] = useParameters(providerConsumerID);
  const [filters] = useFilters();
  const [newParameterFilter, setNewParameterFilter] = useState(
    createParameterFilterSceleton(subscriptionID)
  );

  console.log(parameters);

  if (!parameters) {
    return "";
  }

  return (
    <div>
      <Segment>
        <h2>Filters</h2>
        {parameterFilters.length > 0 ? (
          <Table unstackable>
            <Table.Header>
              <Table.HeaderCell>Event Parameter</Table.HeaderCell>
              <Table.HeaderCell>Filter</Table.HeaderCell>
              <Table.HeaderCell>Scheme</Table.HeaderCell>
              <Table.HeaderCell>Action</Table.HeaderCell>
            </Table.Header>
            <Table.Body>
              {parameterFilters.map((parameterFilter) => (
                <Table.Row>
                  <Table.Cell>{parameterFilter.EventParameter.Name}</Table.Cell>
                  <Table.Cell>{parameterFilter.Filter.Type}</Table.Cell>
                  <Table.Cell> {parameterFilter.Scheme}</Table.Cell>
                  <Table.Cell>
                    <Button
                      icon="close"
                      basic
                      negative
                      onClick={() => deleteParameterFilter(parameterFilter.ID)}
                    />
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        ) : (
          <p>No filters set for this subscription yet.</p>
        )}
        <Form>
          <Form.Group widths="3">
            <Form.Dropdown
              options={parameters.map((eventParameter) => ({
                key: eventParameter.ID,
                text: eventParameter.Name,
                value: eventParameter.ID,
              }))}
              label="Parameter"
              placeholder="Select event parameter"
              value={newParameterFilter.EventParameterID}
              onChange={(_, { value }) =>
                setNewParameterFilter({
                  ...newParameterFilter,
                  EventParameterID: value,
                })
              }
              selection
            />
            <Form.Dropdown
              options={filters.map((filter) => ({
                key: filter.ID,
                text: filter.Type,
                value: filter.ID,
              }))}
              selection
              placeholder="Select filter"
              onChange={(_, { value }) => {
                setNewParameterFilter({
                  ...newParameterFilter,
                  FilterID: value,
                });
              }}
              label="Filter"
              value={newParameterFilter.FilterID}
            />
            <Form.Input
              label="Scheme"
              value={newParameterFilter.Scheme}
              onChange={(_, { value }) =>
                setNewParameterFilter({ ...newParameterFilter, Scheme: value })
              }
            />
          </Form.Group>
        </Form>
        <Button
          content="Add new Filter"
          basic
          fluid
          primary
          icon="plus"
          onClick={async () => {
            await createParameterFilter(newParameterFilter);
            setNewParameterFilter(
              createParameterFilterSceleton(subscriptionID)
            );
          }}
        />
      </Segment>
    </div>
  );
}
