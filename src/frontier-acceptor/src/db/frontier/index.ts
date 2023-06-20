import { ObjectId } from "mongodb";

import { FrontierPage } from "./types";
import { db } from "..";

export const FrontierCollection = db.collection<FrontierPage>("frontier");

export const Frontier = {
  get: async (id: string) => {
    return await FrontierCollection.findOne({ _id: new ObjectId(id) });
  },
  getList: async () => {
    return await FrontierCollection.find({}).toArray();
  },
  create: async (data: FrontierPage) => {
    return await FrontierCollection.insertOne(data);
  },
  delete: async (id: string) => {
    return await FrontierCollection.deleteOne({ _id: new ObjectId(id) });
  },
  update: async (id: string, data: Partial<FrontierPage>) => {
    return await FrontierCollection.updateOne(
      { _id: new ObjectId(id) },
      { $set: data }
    );
  },
};
