import { RequestHandler } from "express";

import { query } from "../../db";
import { FrontierPage } from "./types";

export const get: RequestHandler = async (req, res, next) => {
  try {
    const items = await query<FrontierPage>(
      "SELECT * FROM frontier ORDER BY priority ASC LIMIT 1",
      []
    );
    const item = items[0];

    type MinPrio = { max: number };
    const minPrio =
      (await query<MinPrio>("SELECT MAX(priority) FROM frontier", []))?.[0]
        ?.max ?? 0;

    await query("UPDATE frontier SET priority=$2 WHERE id=$1", [
      item.id,
      minPrio + 1,
    ]);

    // rescale priorities to start from 1
    await query<void>(
      `
        UPDATE frontier 
        SET priority=q.row_number
        FROM
        (SELECT row_number() over(order by priority asc), * from frontier) as q
        WHERE frontier.id = q.id;
    `,
      []
    );

    res.send(item.url);
  } catch (err) {
    next(err);
  }
};
