const path = require('path');
const express = require('express');
const log4js = require('log4js');

const PORT = process.env.PORT || 3000;
const LOG_LEVEL = process.env.LOG_LEVEL || 'info';

const logger = log4js.getLogger('Server');
const app = express();


log4js.configure({
    appenders: {
        'out': { type: 'stdout' },
        server: {
            type: 'multiFile',
            base: 'logs/',
            property: 'categoryName',
            extension: '.log',
            maxLogSize: 10485760,
            backups: 3,
            compress: true
        }
    },
    categories: { default: { appenders: ['out', 'server'], level: LOG_LEVEL } }
});

app.use((req, res, next) => {
    logger.info(req.method, req.headers['x-forwarded-for'] || req.connection.remoteAddress, req.path, req.params, req.query, req.body);
    next();
});

app.get('/', (req, res) => {
    res.status(200).json({
        message: 'Echo Server is running.'
    });
});

app.post('/', (req, res) => {
    req.pipe(res);
});

app.listen(PORT, (err) => {
    if (!err) {
        logger.info('Server is listening on port', PORT);
    } else {
        logger.error(err);
    }
});