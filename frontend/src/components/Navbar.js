import React from "react";
import { Link } from "react-router-dom";
import { Button, Container } from "semantic-ui-react";
export default function Navbar() {
  return (
    <div
      style={{
        width: "100vw",
        height: "70px",
        marginBottom: "2em",
        borderBottom: "3px solid var(--primary-color)",
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
        <Link
          to="/"
          style={{ color: "var(--primary-color)", fontSize: "25px" }}
        >
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
            negative
            basic
            style={{ margin: "2em" }}
            onClick={() => {
              localStorage.removeItem("authToken");
              document.location.href = "/";
            }}
          />
        </div>
      </Container>
    </div>
  );
}
