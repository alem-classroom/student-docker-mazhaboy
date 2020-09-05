const express = require('express');
const app = express();
const PORT = process.env.PORT || '3000';

app.get('/', (req, res) => {
    res.send({ "Hello": process.env.APP_TYPE });
});

app.get('/items/:item_id', (req, res) => {
    const itemID = req.params.item_id
    res.send({
        page: "items",
        item_id: itemID
    });
});

app.listen(parseInt(PORT), () => {
    console.log(`listening on port :${PORT}`);
});
