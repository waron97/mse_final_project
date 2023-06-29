import * as bodyParser from 'body-parser';
import * as cors from 'cors';
import * as express from 'express';
import mongoose from 'mongoose';

import appRouter from './api';
import { createDefaultKeys } from './api/keys/service';
import { scheduleObsoleteLogsRemoval } from './api/logs/service';
import appEnv from './constants/env';
import { mongoUri } from './constants/mongo';

mongoose.connect(mongoUri).then(() => {
  createDefaultKeys();
  scheduleObsoleteLogsRemoval();
});

const app = express();

app.use(express.static('static'));
app.use(express.static('frontend'));
app.use(cors());
app.use(bodyParser.json());
app.use(appRouter);

setImmediate(() => {
  // eslint-disable-next-line
  console.log('Starting logs server with params', appEnv);
});

app.listen(appEnv.port, () => {
  // eslint-disable-next-line
  console.log(`Logs service listening on port ${appEnv.port}`);
});

export default app;
