import { RequestHandler } from 'express';

import Key, { KeyType } from '../../api/keys/model';

interface Params {
  types: KeyType[];
}

export const apiKey: (params?: Params) => RequestHandler =
  (params) => async (req, res, next) => {
    const { authorization } = req.headers;

    const types = params?.types ?? ['admin', 'readonly', 'writeonly'];

    const authFail = () =>
      res.status(401).json({
        error: 'INVALID_API_KEY',
        message: 'No API key or invalid API key provided.',
      });

    if (!authorization) {
      authFail();
      return;
    }

    const apiKey = authorization?.split?.('apiKey ')?.[1];

    if (apiKey) {
      const key = await Key.findOne({ key: apiKey });
      if (key && types.includes(key.type)) {
        next();
      } else {
        authFail();
      }
    } else {
      authFail();
    }
  };
