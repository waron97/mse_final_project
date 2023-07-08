import { MongoClient } from "mongodb";

import appEnv from "../util/constants";

export const client = new MongoClient(appEnv.mongoUri, {
  monitorCommands: true,
});

await client.connect();
export const db = client.db("mse");
