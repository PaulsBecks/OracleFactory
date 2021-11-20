import { Label } from "semantic-ui-react";

export default function OracleOnOffRibbon({ oracleStarted }) {
  return (
    <Label
      ribbon
      content={oracleStarted ? "Subscribed" : "Unsubscribed"}
      color={oracleStarted ? "green" : "red"}
    />
  );
}
