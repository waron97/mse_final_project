import { RequestHandler } from "express";

import { crawlCollection } from "../../db/crawl";

export const create: RequestHandler = async (req, res, next) => {
  try {
    const { body } = req;
    const { url } = body;
    const date = new Date();
    const found = await crawlCollection.findOne({ url });
    if (found) {
      await crawlCollection.updateOne({ url }, { $set: { ...body, date } });
    } else {
      await crawlCollection.insertOne({ ...body, date });
    }
    res.status(201).send();
  } catch (error) {
    next(error);
  }
};
