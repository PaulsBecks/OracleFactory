import React, { useState } from "react";
import { Link } from "react-router-dom";
import { Form, Message } from "semantic-ui-react";
import postData from "../services/postData";

const SIGNUP = "signup";
const LOGIN = "login";

export default function Login() {
  const [loginData, setLoginData] = useState({ password: "", email: "" });
  const [showErrorMessage, setShowErrorMessage] = useState(false);
  const [loginOrSignup, setLoginOrSignup] = useState(LOGIN);

  const changeForm = ({ target: { value, name } }) => {
    setLoginData({ ...loginData, [name]: value });
  };

  const login = async () => {
    try {
      const url = loginOrSignup === LOGIN ? "/users/login" : "/users/signup";
      const response = await postData(url, loginData);
      window.localStorage.setItem(
        "authToken",
        JSON.stringify({ token: response.token })
      );
      window.location.reload();
    } catch (err) {
      setShowErrorMessage(true);
      setTimeout(() => setShowErrorMessage(false), 5000);
    }
  };

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
        width: "100%",
      }}
    >
      <div
        style={{
          width: "40em",
          padding: "2em",
        }}
      >
        <Form onSubmit={login}>
          <h2>{loginOrSignup === LOGIN ? "Login" : "Signup"}</h2>
          {showErrorMessage && (
            <Message negative>Email or password not found.</Message>
          )}
          <Form.Input
            label="Email"
            type="text"
            placeholder="Your email address"
            value={loginData.email}
            name="email"
            onChange={changeForm}
          />
          <Form.Input
            label="Password"
            type="password"
            placeholder="Your Passwort"
            value={loginData.password}
            name="password"
            onChange={changeForm}
          />
          <Form.Button fluid type="submit" positive>
            {loginOrSignup === LOGIN ? "Login" : "Signup"}
          </Form.Button>
          <p>
            Or{" "}
            {loginOrSignup === LOGIN ? (
              <a
                onClick={() => setLoginOrSignup(SIGNUP)}
                style={{ cursor: "pointer" }}
              >
                create a new account
              </a>
            ) : (
              <a
                onClick={() => setLoginOrSignup(LOGIN)}
                style={{ cursor: "pointer" }}
              >
                log in
              </a>
            )}
            .
          </p>
        </Form>
      </div>
    </div>
  );
}
