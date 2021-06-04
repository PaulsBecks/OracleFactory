import React from "react";
import { Link } from "react-router-dom";
import { Container } from "semantic-ui-react";
export default function Navbar() {
  return (
    <div
      style={{
        width: "100vw",
        height: "70px",
        backgroundColor: "var(--primary-color)",
        marginBottom: "2em",
        display: "flex",
        justifyContent: "flex-start",
        alignItems: "center",
      }}
    >
      <Container>
        <Link to="/" style={{ color: "white", fontSize: "25px" }}>
          Oracle Factory
        </Link>
      </Container>
    </div>
  );
}
