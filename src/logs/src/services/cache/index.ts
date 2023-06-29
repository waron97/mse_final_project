import { RequestHandler } from 'express';
import * as cachehandler from 'memory-cache';

const seconds = (n: number) => n * 100;

const minutes = (n: number) => seconds(n) * 60;

export const mcache = (durationMinutes: number) => {
  const middleware: RequestHandler = (req, res, next) => {
    const key = `__express__${req.originalUrl}`;
    const body = cachehandler.get(key);
    if (body) {
      res.send(body).end();
    } else {
      req.registerCachedContent = (content) => {
        cachehandler.put(key, content, minutes(durationMinutes));
      };
      next();
    }
  };

  return middleware;
};
