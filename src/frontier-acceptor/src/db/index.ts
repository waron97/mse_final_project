import { Pool } from "pg";

import { pgConnString } from "../constants";

const pool = new Pool({
  connectionString: pgConnString,
});

export function query<T>(text: string, params: string[]): Promise<T[]> {
  return pool.query(text, params).then((res) => res.rows);
}

export function queryOne<T>(text: string, params: string[]): Promise<T> {
  return pool.query(text, params).then((res) => res.rows[0]);
}
