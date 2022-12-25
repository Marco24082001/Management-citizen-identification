const { gateway } = require('../services/gateway')
module.exports.create = async function(req, res) {
  try {
    const { id, name, dateOfbirth, sex, nationality, placeOforigin, placeOfresidence, dateOfExpiry, personalIdentification, personSignature, avatarUrl, signatureUrl, createdDate, isDelete} = req.body;
    const network = await gateway.getNetwork('fabidchannel');
    const contract = network.getContract('fabid');
    await contract.submitTransaction('CreateIdentication', id, name, dateOfbirth, sex, nationality, placeOforigin, placeOfresidence, dateOfExpiry, personalIdentification, personSignature, avatarUrl, signatureUrl, createdDate, isDelete);
    console.log('Transaction has been submitted');
    return res.send('Transaction has been submitted');
  } catch (error){
    return res.json({ error : error });
  }
}

module.exports.get = async function(req, res) {
  try {
    const network = await gateway.getNetwork('fabidchannel');
    const contract = network.getContract('fabid');
    const id = req.params.id;
    result = JSON.parse(await contract.evaluateTransaction('QueryIdentication', id));
    return res.status(200).json({ response: result })
  } catch (error) {
    return res.json({ error : error });
  }
}

module.exports.getAll = async function(req, res) {
  try {
    const network = await gateway.getNetwork('fabidchannel');
    const contract = network.getContract('fabid');
    const results = JSON.parse(await contract.evaluateTransaction('QueryAllIdentications'));
    return res.status(200).json({ response: results });
  } catch (error) {
    return res.json({ error : error });
  }
}

module.exports.edit = async function(req, res) {
  try {
    const network = await gateway.getNetwork("fabidchannel");
    const contract = network.getContract("fabid");
    
    const {id, name, dateOfbirth, sex, nationality, placeOforigin, placeOfresidence, dateOfExpiry, personalIdentification, personSignature, avatarUrl, signatureUrl, createdDate, isDelete} = req.body;
    console.log(id, name, dateOfbirth, sex, nationality, placeOforigin, placeOfresidence, dateOfExpiry, personalIdentification, personSignature, avatarUrl, signatureUrl, createdDate, isDelete)
    await contract.submitTransaction("UpdateIdentication", id, name, dateOfbirth, sex, nationality, placeOforigin, placeOfresidence, dateOfExpiry, personalIdentification, personSignature, avatarUrl, signatureUrl, createdDate, isDelete);
    return res.status(200).json({response: "success"})
  } catch (error) {
    console.log(error);
    return res.json({ error : error });
  }
}

module.exports.delete = async function(req, res) {
  try {
    const network = await gateway.getNetwork("fabidchannel");
    const contract = network.getContract("fabid");
    const id = req.params.id;
    await contract.submitTransaction("deleteIdentication", id);
    return res.status(200).json({ response: "success" });
  } catch (error) {
    return res.json({ error: error });
  }
}

module.exports.getHistory = async function(req, res) {
  try {
    const network = await gateway.getNetwork("fabidchannel");
    const contract = network.getContract("fabid");
    const id = req.params.id;
    console.log(id);
    const results = JSON.parse(await contract.evaluateTransaction("GetAssetHistory", "IDENTICATION" + id));
    return res.status(200).json({ response: results})
  } catch (error) {
    console.log(error);
    return res.json({ error: error });
  }
}

module.exports.upload = async function(req, res) {
  console.log(req.body);
  return res.status(200).json({ response: "success" });
}
