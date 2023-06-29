import bodyParser from "body-parser";
import cors from "cors";
import express from "express";

import acceptorRouter from "./api/acceptor";
import frontierRouter from "./api/frontier";
import initFrontier from "./api/frontier/initFrontier";
import Log from "./util/logs";

async function initDb() {
  initFrontier();
}

export default async function startServer() {
  await initDb();
  const app = express();
  app.use(bodyParser.json({ limit: "50mb" }));
  app.use(cors({ origin: "*" }));

  app.get("/", (_, res) => res.send("ok"));

  app.use("/frontier", frontierRouter);
  app.use("/acceptor", acceptorRouter);

  app.listen(3000, () => {
    // eslint-disable-next-line
    console.log("Frontier/Acceptor started on container port 3000");

    Log.info("Bootstrap", "Frontier/Acceptor started on container port 3000");
  });
}
