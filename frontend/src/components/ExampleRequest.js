const typeToExampleParameter = {
  uint8: "1",
  uint16: "1",
  uint32: "1",
  uint64: "1",
  uint128: "1",
  uint256: "1",
  uint: "1",
  int: "1",
  integer: "1",
  bytes32: '"Hello there!"',
  address: '"0x07A93d6C2D964b702662971Efaca43499fEB198c"',
  bytes: '"Hello there!"',
  bool: '"true"',
  string: '"Hello there!"',
};

export default function ExampleRequest({ eventParameters }) {
  return (
    <p>
      <b>Example Request:</b>
      <pre
        style={{
          background: "#232324",
          color: "white",
          padding: "1em",
          borderRadius: "10px",
        }}
      >
        {JSON.stringify(
          JSON.parse(
            "{" +
              eventParameters
                .map(
                  (parameter) =>
                    '"' +
                    parameter.Name +
                    '": ' +
                    typeToExampleParameter[parameter.Type]
                )
                .join(",") +
              "}"
          ),
          null,
          2
        )}
      </pre>
    </p>
  );
}
