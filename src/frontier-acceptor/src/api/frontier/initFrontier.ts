import { Client } from "pg";

import { pgConnString } from "../../constants";

// this is wrong because it
const createCommand = `
    CREATE TABLE IF NOT EXISTS frontier (
        id SERIAL PRIMARY KEY,
        url VARCHAR(255) NOT NULL,
        priority INTEGER NOT NULL
    );
`;

const initialPages = [
  {
    url: "https://www.google.com",
    priority: 1,
  },
  {
    url: "https://www.facebook.com",
    priority: 1,
  },
  {
    url: "https://www.twitter.com",
    priority: 1,
  },
];

export default async function initFrontier() {
  const client = new Client({ connectionString: pgConnString });
  await client.connect();
  await client.query(createCommand);
  const frontierPages = await client.query("SELECT * FROM frontier");
  if (frontierPages.rowCount !== 0) {
    // frontier already initialized
    return;
  }

  for (const page of initialPages) {
    await client.query("INSERT INTO frontier (url, priority) VALUES ($1, $2)", [
      page.url,
      page.priority,
    ]);
  }
}
