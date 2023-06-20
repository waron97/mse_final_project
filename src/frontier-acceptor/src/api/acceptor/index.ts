import { Router } from "express";

import { create } from "./controller";

const acceptorRouter = Router();
acceptorRouter.post("/", create);

export default acceptorRouter;
