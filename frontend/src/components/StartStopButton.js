import { Button } from "semantic-ui-react";

export default function StartStopButton({
  loading,
  oracleStarted,
  stopOracle,
  startOracle,
}) {
  return (
    <Button
      loading={loading}
      content={oracleStarted ? "Stop" : "Start"}
      color={oracleStarted ? "negative" : "positive"}
      onClick={oracleStarted ? stopOracle : startOracle}
    />
  );
}
