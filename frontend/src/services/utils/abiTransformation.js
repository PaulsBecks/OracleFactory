export function formToAbi(oracles) {
  return JSON.stringify(
    oracles.map((oracle) => {
      delete oracle.ContractAddress;
      return oracle;
    })
  )
    .replace(/"ContractName":/g, '"name":')
    .replace(/"Name":/g, '"name":')
    .replace(/"Type":/g, '"type":');
}

export function parseAbi(abi) {
  let oracles = [];
  try {
    JSON.parse(abi);
    const parsedAbi = JSON.parse(
      abi.replace(/"name":/g, '"Name":').replace(/"type":/g, '"Type":')
    );
    oracles = parsedAbi.map((oracle) => {
      const _oracle = { ...oracle, ContractName: oracle.Name };
      delete _oracle.Name;
      return _oracle;
    });
    return oracles;
  } catch (err) {
    console.log(err, "No valid abi");
  }
  return [];
}
