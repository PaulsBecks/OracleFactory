import { Container, Icon } from "semantic-ui-react";

export default function Footer() {
  return (
    <div
      style={{
        borderTop: "4px solid var(--primary-color)",
        marginTop: "5em",
        padding: "2em 0",
      }}
    >
      <Container>
        <p style={{ textAlign: "center" }}>
          <a href="https://github.com">
            View the source code <Icon size="large" name="github" />
          </a>
        </p>
      </Container>
    </div>
  );
}
