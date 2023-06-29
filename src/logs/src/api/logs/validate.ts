import { RequestHandler } from 'express';
import * as Yup from 'yup';

import { LogLevels } from './model';

export const validate: RequestHandler = (req, res, next) => {
  const schema = Yup.object().shape({
    level: Yup.string().required().oneOf(LogLevels),
    appId: Yup.string().required(),
    location: Yup.string().required(),
    message: Yup.string(),
    detail: Yup.mixed(),
  });

  schema
    .validate(req.body, { stripUnknown: true, abortEarly: false })
    .then((validated) => {
      req.body = validated;
      next();
    })
    .catch((reason) => {
      const errors = reason.inner.map(({ path, type, errors }) => ({
        path,
        type,
        error: errors[0],
      }));
      res.status(422).json({ errors }).end();
    });
};
