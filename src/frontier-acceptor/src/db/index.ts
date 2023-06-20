import { MongoClient } from "mongodb";

import { mongoConnectionString } from "../constants";

export const client = new MongoClient(mongoConnectionString, {
  monitorCommands: true,
});

await client.connect();
export const db = client.db("mse");
