import { RequestHandler } from "express";

import { query } from "../../db";
import { FrontierPage } from "./types";

export const get: RequestHandler = async () => {
  const items = await query<FrontierPage>(
    "SELECT * FROM frontier ORDER_BY priority DESC LIMIT 1",
    []
  );
  return items[0];
};
