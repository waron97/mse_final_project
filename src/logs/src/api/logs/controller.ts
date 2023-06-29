import { RequestHandler } from 'express';

import { paginated, success } from '../../services/responses';
import Log from './model';
import { processLogFilters } from './service';

export const index: RequestHandler = async (req, res, next) => {
  const filters = processLogFilters(req.query);
  const {
    pagination: { page, pageSize, limit, skip },
  } = req;
  Log.count(filters)
    .then((size) => {
      return Log.find(filters)
        .skip(skip)
        .limit(limit)
        .sort({ date: 'desc' })
        .then(paginated(res, { page, pageSize, size }))
        .catch(next);
    })
    .catch(next);
};

export const create: RequestHandler = (req, res, next) => {
  Log.create(req.body).then(success(res, 201)).catch(next);
};

export const getAppIds: RequestHandler = async (req, res, next) => {
  try {
    const docs = await Log.distinct('appId');
    if (docs && docs.length) {
      req?.registerCachedContent?.(docs);
      return res.send(docs);
    }
    res.send([]);
  } catch (e) {
    next(e);
  }
};
