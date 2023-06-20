import { MongoClient } from "mongodb";

import { mongoConnectionString } from "../constants";

export const client = new MongoClient(mongoConnectionString);
await client.connect();
export const db = client.db("mse");
