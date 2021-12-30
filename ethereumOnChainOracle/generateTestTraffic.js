const Oracle = artifacts.require("Oracle");

async function generateTestTraffic() {
  try {
    let instance = await Oracle.deployed();
    //let address = "0x40536521353F9f4120A589C9ddDEB6188EF61922";
    //let balance = await web3.eth.getBalance(address);
    let contracts = [
      "0x68697Ed883c1b51d14370991dA756577DDCCBc7A", //SimpleStorage
      "0xe3Fb42873f615fcF8b0Af6e1580A7E35ec04798b", //SimpleStorage1
      "0x6e10CD1cC7c760903afa08FD504c5302a148F490", //SimpleStorage2
    ];
    let topic = "/this/is/the/topic";
    let value = 1;
    for (let contractPos in contracts) {
      await instance.subscribeInteger(topic, contracts[contractPos]);
      let total = parseInt(contractPos) + 1;
      let tx = await instance.publishInteger(topic, value);
      let gasPerEvent = tx.receipt.gasUsed / total;
      console.log(
        "Notify costs for " +
          total +
          " subscriptions per event: " +
          gasPerEvent +
          " total: " +
          tx.receipt.gasUsed
      );
      value++;
    }
    for (let contractPos in contracts) {
      let total = 3 - parseInt(contractPos);
      let tx = await instance.publishInteger(topic, value);
      let gasPerEvent = tx.receipt.gasUsed / total;
      console.log(
        "Notify costs for " +
          total +
          " subscriptions per event: " +
          gasPerEvent +
          " total: " +
          tx.receipt.gasUsed
      );
      await instance.unsubscribeInteger(topic, contracts[contractPos]);
      value++;
    }
    /*balance = await web3.eth.getBalance(address);
    console.log(balance);
    
    console.log(publishIntegerResult.receipt.rawLogs);
    balance = await web3.eth.getBalance(address);
    console.log(balance);
    let unsubscribeResult = await instance.unsubscribeInteger(
      "/this/is/the/topic",
      "0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00"
    );
    console.log(unsubscribeResult.logs);
    balance = await web3.eth.getBalance(address);
    console.log(balance);*/
  } catch (err) {
    console.log(err);
  }
}

module.exports = generateTestTraffic;
