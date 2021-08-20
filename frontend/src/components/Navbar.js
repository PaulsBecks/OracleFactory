import React from "react";
import { Link } from "react-router-dom";
import { Button, Container } from "semantic-ui-react";
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
      <Container
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <Link to="/" style={{ color: "white", fontSize: "25px" }}>
          Oracle Factory
        </Link>
        <div>
          <Button
            basic
            icon="settings"
            content="Settings"
            as={Link}
            to="/settings"
          />
          <Button
            content="Logout"
            primary
            style={{ margin: "2em" }}
            onClick={() => {
              localStorage.removeItem("authToken");
              window.location.reload();
            }}
          />
        </div>
      </Container>
    </div>
  );
}
