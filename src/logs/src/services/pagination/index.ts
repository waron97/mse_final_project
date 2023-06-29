import { RequestHandler } from 'express';

export const pagination: () => RequestHandler = () => (req, res, next) => {
  const { page: page_str = '1', pageSize: pageSize_str = '20' } = req.query;
  const page = parseInt(page_str as string);
  const pageSize = parseInt(pageSize_str as string);

  const skip = (page - 1) * pageSize;
  const limit = pageSize;

  req.pagination = {
    skip,
    limit,
    page,
    pageSize,
  };

  next();
};
