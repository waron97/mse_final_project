import { RequestHandler } from "express";

import { crawlCollection } from "../../db/crawl";

export const create: RequestHandler = async (req, res, next) => {
  try {
    const { body } = req;
    const { url } = body;
    const date = new Date();
    const found = await crawlCollection.findOne({ url });
    if (found) {
      await crawlCollection.updateOne(
        { _id: found._id },
        { $set: { ...body, date } }
      );
    } else {
      await crawlCollection.insertOne({ ...body, date });
    }
    res.status(201).send();
  } catch (error) {
    next(error);
  }
};

export const get: RequestHandler = async (req, res, next) => {
  let { url } = req.params;
  url = decodeURIComponent(url);
  const item = await crawlCollection.findOne({ url });
  if (item) {
    res.status(200).json(item);
  } else {
    res.status(404).send();
  }
};
