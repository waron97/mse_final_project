import { Router } from 'express';

import { apiKey } from '../services/auth/apiKey';
import logsRouter from './logs';

const appRouter = Router();

appRouter.get('/', (req, res) => {
  res.send('Server alive');
});

appRouter.get('/validate-key', apiKey(), (req, res) => res.send('ok'));

appRouter.use('/logs', logsRouter);

export default appRouter;
