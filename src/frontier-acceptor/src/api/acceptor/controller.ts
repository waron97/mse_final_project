import { RequestHandler } from "express";
import { ObjectId } from "mongodb";

import { crawlCollection } from "../../db/crawl";
import Log from "../../util/logs";
import splitPassages from "../../util/passages";

export const create: RequestHandler = async (req, res, next) => {
  try {
    const { body } = req;
    const passages = splitPassages(body.mainText || body.bodyText || "");
    body.passages = passages.map((p) => {
      return {
        text: p,
        _id: new ObjectId(),
      };
    });
    const { url } = body;
    const crawlDate = new Date();
    const found = await crawlCollection.findOne({ url });
    if (found) {
      Log.debug("acceptor.create", "Updating crawled page", { url });
      await crawlCollection.updateOne(
        { _id: found._id },
        { $set: { ...body, crawlDate } }
      );
    } else {
      Log.debug("acceptor.create", "Creating crawled page", { url });
      await crawlCollection.insertOne({ ...body, crawlDate });
    }
    res.status(201).send();
  } catch (error) {
    Log.error("acceptor.create", "Error creating crawled page", error);
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
