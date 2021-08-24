import { Label } from "semantic-ui-react";

export default function OracleOnOffRibbon({ oracleStarted }) {
  return (
    <Label
      ribbon
      content={oracleStarted ? "ON" : "OFF"}
      color={oracleStarted ? "green" : "red"}
    />
  );
}
