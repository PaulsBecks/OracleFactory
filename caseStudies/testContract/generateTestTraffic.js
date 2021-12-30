const SimpleStorage = artifacts.require("SimpleStorage");
const SimpleStorage1 = artifacts.require("SimpleStorage1");

async function generateTestTraffic() {
  try {
    let instance = await SimpleStorage.deployed();
    let instance1 = await SimpleStorage1.deployed();
    let address = "0x40536521353F9f4120A589C9ddDEB6188EF61922";

    let topic = "topic";
    let value = 1;

    let tx = await instance.subscribeOnChain(topic);
    console.log("Subscribe on chain costs: " + tx.receipt.gasUsed);

    tx = await instance1.subscribeOnChain(topic);
    console.log("Subscribe on chain costs: " + tx.receipt.gasUsed);

    tx = await instance.unsubscribeOnChain(topic);
    console.log("Unsubscribe on chain costs: " + tx.receipt.gasUsed);

    tx = await instance.subscribeOffChain(topic);
    console.log("Subscribe off chain costs: " + tx.receipt.gasUsed);

    tx = await instance.unsubscribeOffChain(topic);
    console.log("Unsubscribe off chain costs: " + tx.receipt.gasUsed);

    tx = await instance.integerCallback(topic, value);
    console.log("Notification costs: " + tx.receipt.gasUsed);
  } catch (err) {
    console.log(err);
  }
}

module.exports = generateTestTraffic;
