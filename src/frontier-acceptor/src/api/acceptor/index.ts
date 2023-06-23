import { Router } from "express";

import { create, get } from "./controller";

const acceptorRouter = Router();
acceptorRouter.post("/", create);
acceptorRouter.get("/:url", get);

export default acceptorRouter;
