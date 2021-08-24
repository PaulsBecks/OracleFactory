import { Button } from "semantic-ui-react";

export default function StartStopButton({
  loading,
  oracleStarted,
  stopOracle,
  startOracle,
}) {
  return (
    <Button
      basic
      fluid
      loading={loading}
      content={oracleStarted ? "Stop Oracle" : "Start Oracle"}
      color={oracleStarted ? "negative" : "positive"}
      onClick={oracleStarted ? stopOracle : startOracle}
    />
  );
}
