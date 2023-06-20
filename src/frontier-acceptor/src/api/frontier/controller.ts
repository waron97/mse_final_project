import { RequestHandler } from "express";

import { FrontierCollection } from "../../db/frontier";

export const get: RequestHandler = async (req, res, next) => {
  try {
    const items = await FrontierCollection.find({})
      .sort("priority", 1)
      .toArray();

    const item = items[0];

    const lastItem = items[items.length - 1];

    if (!item) {
      throw new Error("No items in frontier");
    }

    items[0].priority = lastItem.priority + 1;
    items.sort((a, b) => a.priority - b.priority);

    for (let i = 0; i < items.length; i++) {
      const curr = items[i];

      await FrontierCollection.updateOne(
        { _id: curr._id },
        {
          $set: {
            priority: i + 1,
          },
        }
      );
    }

    res.send(item.url);
  } catch (err) {
    next(err);
  }
};
