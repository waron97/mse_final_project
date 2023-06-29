import { Response } from 'express';

export const success =
  (res: Response, status = 200) =>
  (resource: any) => {
    res.status(status).json(resource);
  };

export const paginated =
  (res: Response, { page, pageSize, size }) =>
  (records: any) => {
    res.status(200).json({
      data: records,
      pagination: {
        page,
        pageSize,
        maxPages: Math.ceil(size / pageSize) || 1,
        resultCount: size,
      },
    });
  };
