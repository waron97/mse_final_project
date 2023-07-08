import * as bodyParser from 'body-parser';
import * as cors from 'cors';
import * as express from 'express';
import mongoose from 'mongoose';

import appRouter from './api';
import { createDefaultKeys } from './api/keys/service';
import { scheduleObsoleteLogsRemoval } from './api/logs/service';
import appEnv from './constants/env';
import { dbName } from './constants/mongo';

mongoose.connect(appEnv.mongoUri, { dbName }).then(() => {
  createDefaultKeys();
  scheduleObsoleteLogsRemoval();
});

const app = express();

app.use(express.static('static'));
app.use(express.static('frontend'));
app.use(cors());
app.use(bodyParser.json({ limit: '50mb' }));
app.use(appRouter);

setImmediate(() => {
  // eslint-disable-next-line
  console.log('Starting logs server with params', appEnv);
});

app.listen(appEnv.port, () => {
  // eslint-disable-next-line
});

export default app;
