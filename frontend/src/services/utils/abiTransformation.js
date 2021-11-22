export function formToAbi(oracles) {
  return JSON.stringify(
    oracles.map((oracle) => {
      const _oracle = { ...oracle };
      oracle["readableName"] = oracle.Name;
      delete _oracle.Name;
      return oracle;
    })
  )
    .replace(/"EventName":/g, '"name":')
    .replace(/"Name":/g, '"name":')
    .replace(/"Type":/g, '"type":');
}

export function parseAbi(abi) {
  let oracles = [];
  try {
    let parsedAbi = JSON.parse(
      abi.replace(/"name":/g, '"Name":').replace(/"type":/g, '"Type":')
    );

    parsedAbi = parsedAbi.filter(
      (method) =>
        (method.Type === "function" || method.Type === "event") &&
        method.stateMutability !== "view" &&
        method.stateMutability !== "pure"
    );
    oracles = parsedAbi
      .map((oracle) => ({
        ...oracle,
        EventName: oracle.Name,
      }))
      .map((oracle) => ({ ...oracle, Name: oracle.readableName }));
    return oracles;
  } catch (err) {}
  return [];
}
