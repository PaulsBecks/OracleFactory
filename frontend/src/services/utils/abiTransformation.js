export function formToAbi(subscriptions) {
  return JSON.stringify(
    subscriptions.map((subscription) => {
      delete subscription.ContractAddress;
      return subscription;
    })
  )
    .replace(/"ContractName":/g, '"name":')
    .replace(/"Name":/g, '"name":')
    .replace(/"Type":/g, '"type":');
}

export function parseAbi(abi) {
  let subscriptions = [];
  try {
    JSON.parse(abi);
    const parsedAbi = JSON.parse(
      abi.replace(/"name":/g, '"Name":').replace(/"type":/g, '"Type":')
    );
    subscriptions = parsedAbi.map((subscription) => {
      const _subscription = {
        ...subscription,
        ContractName: subscription.Name,
      };
      delete _subscription.Name;
      return _subscription;
    });
    return subscriptions;
  } catch (err) {
    console.log(err, "No valid abi");
  }
  return [];
}
