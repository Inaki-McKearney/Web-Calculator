'use strict';

const express = require('express');

const PORT = 80;
const HOST = '0.0.0.0';

var sub = require('./subtract');

const app = express();
app.get('/subtract', (req,res) => {

    var output = {
        'error': false,
        'string': '',
        'answer': 0
    };

    res.setHeader('Content-Type', 'application/json');
    // res.setHeader('Access-Control-Allow-Origin', '*')
    var x = req.query.x;
    var y = req.query.y;
    if (isNaN(x) || isNaN(y)){
        output.error = true;
        output.answer = 0;
    }
    else {
        output.answer = sub.subtract(x,y);
        output.string = x + '-' + y + '=' + output.answer;
    }

    if (isNaN(x))
        output.string = 'x parameter invalid/missing'
    else if (isNaN(y))
        output.string = 'y parameter invalid/missing'

    res.end(JSON.stringify(output));
});

app.listen(PORT, HOST);

module.exports = app; // for testing

