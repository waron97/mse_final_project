import { RequestHandler } from "express";

import { crawlCollection } from "../../db/crawl";

export const create: RequestHandler = async (req, res, next) => {
  try {
    const { body } = req;
    const { url } = body;
    const crawlDate = new Date();
    const found = await crawlCollection.findOne({ url });
    if (found) {
      await crawlCollection.updateOne(
        { _id: found._id },
        { $set: { ...body, crawlDate } }
      );
    } else {
      await crawlCollection.insertOne({ ...body, crawlDate });
    }
    res.status(201).send();
  } catch (error) {
    next(error);
  }
};

export const get: RequestHandler = async (req, res, next) => {
  try {
    let { url } = req.params;
    url = decodeURIComponent(url);
    const item = await crawlCollection.findOne({ url });
    if (item) {
      res.status(200).json(item);
    } else {
      res.status(404).send();
    }
  } catch (e) {
    if (e instanceof URIError) {
      res.status(404).send();
      // eslint-disable-next-line
      console.log("maflormed url requested", req.params.url);
    }
    next(e);
  }
};
