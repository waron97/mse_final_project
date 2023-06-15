import bodyParser from "body-parser";
import cors from "cors";
import express from "express";

import frontierRouter from "./api/frontier";
import initFrontier from "./api/frontier/initFrontier";

async function initDb() {
  initFrontier();
}

export default async function startServer() {
  await initDb();
  const app = express();
  app.use(bodyParser.json());
  app.use(cors({ origin: "*" }));

  app.get("/", (_, res) => res.send("ok"));

  app.use("/frontier", frontierRouter);

  app.listen(3000, () => {
    // eslint-disable-next-line
    console.log("Frontier/Acceptor started on container port 3000");
  });
}
