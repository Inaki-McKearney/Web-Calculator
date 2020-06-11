'use strict';

const express = require('express');
const fs = require('fs');

const handlebars = require('handlebars');

const PORT = process.env.PORT || 80;
const HOST = '0.0.0.0';

const app = express();

const proxyURL = process.env.PROXYURL;


app.get("/", (req, res) => {
    fs.readFile('./index.html', function ( err, template) {
        if (err) {
            //
        }

        var temp = handlebars.compile(template.toString())
        res.end(temp({
            proxy: proxyURL
        }))
    })
});


app.listen(PORT, () => {
    console.log(`Listening to requests on http://localhost:${PORT}`);
});
