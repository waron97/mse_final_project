import { Router } from "express";

import { get } from "./controller";

const frontierRouter = Router();

frontierRouter.get("/", get);

export default frontierRouter;
