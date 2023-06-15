import { Pool } from "pg";

import { pgConnString } from "../constants";

const pool = new Pool({
  connectionString: pgConnString,
});

type Params = (string | number)[];

export function query<T>(text: string, params: Params): Promise<T[]> {
  return pool.query(text, params).then((res) => res.rows);
}

export function queryOne<T>(text: string, params: Params): Promise<T> {
  return pool.query(text, params).then((res) => res.rows[0]);
}
