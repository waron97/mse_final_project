import { RequestHandler } from "express";

import { FrontierCollection } from "../../db/frontier";

export const get: RequestHandler = async (req, res, next) => {
  try {
    const items = await FrontierCollection.find({})
      .sort({ priority: -1 })
      .toArray();
    const item = items[0];

    const lastItem = items[items.length - 1];

    if (!item) {
      throw new Error("No items in frontier");
    }

    await FrontierCollection.updateOne(item._id, {
      $set: { priority: lastItem.priority + 1 },
    });

    res.send(item.url);
  } catch (err) {
    next(err);
  }
};
