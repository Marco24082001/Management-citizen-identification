const express = require('express');
const router = express.Router();
const controller= require('../controllers/identication.controller');

router.get('', controller.getAll);
router.get("/:id", controller.get);
router.post('', controller.create);
router.put('/:id', controller.edit);
router.delete("/:id", controller.delete);
router.get('/history/:id', controller.getHistory);
router.post('/upload', controller.upload);

module.exports = router;