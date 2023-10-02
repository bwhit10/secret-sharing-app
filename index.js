const path = require('node:path');
const express = require('express');

const app = express();
app.use(express.urlencoded({ extended: false }));

app.engine('.html', require('ejs').__express);
app.set('views', path.join(__dirname, 'templates'));

app.use('/assets', express.static('./assets'));

app.get('/', (_, res) => {
  res.render('index.html');
});

app.post('/save', (req, res) => {
    // TODO: save req.body.text to display later
    res.send('TODO');
});

app.listen(3000, () => {
  console.log('started on 3000');
});
