import { Router } from 'express';

import { apiKey } from '../../services/auth/apiKey';
import { mcache } from '../../services/cache';
import { pagination } from '../../services/pagination';
import { create, getAppIds, index } from './controller';
import { validate } from './validate';

const router = Router();

router.get('/', apiKey({ types: ['readonly', 'admin'] }), pagination(), index);

router.get('/app-ids', apiKey({ types: ['admin'] }), mcache(5), getAppIds);

router.post('/', apiKey({ types: ['admin', 'writeonly'] }), validate, create);

export default router;
