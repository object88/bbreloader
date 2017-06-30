import express from 'express';
import proxy from 'express-http-proxy';

const app = express();

const APP_PORT = 8181;

app.use('/', proxy('www.google.com'));

app.listen(APP_PORT, () => {
  console.log(`App is now running on http://localhost:${APP_PORT}`);
});