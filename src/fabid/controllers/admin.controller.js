// const { gateway } = require('../services/gateway')
// const { sign } = require('jsonwebtoken');
// require('dotenv').config()


module.exports.login = async function(req, res) {
  try {
    // const network = await gateway.getNetwork('fabidchannel');
    // const contract = network.getContract('fabid');
    const { email, password} = req.body;
    
    // const admin = JSON.parse(await contract.evaluateTransaction('readAdminByEmail', email));
    // if (admin[0].password === password) {
    //     return res.json({ response: "Success" });
    // } else {
    //     return res.json({ error: "Password wrong" });
    // }
    return res.json({ response: "Success"});
  } catch (err){
    console.log(err);
    return res.json({ error: "System catch error"});
  }
}
