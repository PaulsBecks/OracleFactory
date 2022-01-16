const Token = artifacts.require("Token");

async function generateTestTraffic() {
  try {
    let instance = await Token.deployed();
    instance.mint("0x40536521353F9f4120A589C9ddDEB6188EF61922", 100000);
    let i = 0;
    while (i < 300) {
      await instance.transfer("0xE2bC23ad9d3D3a4Ab205AA46cbbD7648A2eEaaD4", 1);
      i++;
      console.log(i);
    }
  } catch (err) {
    console.log(err);
  }
}

module.exports = generateTestTraffic;
