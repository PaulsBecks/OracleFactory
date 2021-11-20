import { Card, Placeholder } from "semantic-ui-react";

export function PlaceholderCard() {
  return (
    <Card>
      <Card.Content>
        <Placeholder>
          <Placeholder.Header image>
            <Placeholder.Line />
            <Placeholder.Line />
          </Placeholder.Header>
          <Placeholder.Paragraph>
            <Placeholder.Line />
            <Placeholder.Line />
            <Placeholder.Line />
            <Placeholder.Line />
          </Placeholder.Paragraph>
        </Placeholder>
      </Card.Content>
    </Card>
  );
}
