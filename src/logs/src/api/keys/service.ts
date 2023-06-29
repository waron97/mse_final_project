import appEnv from '../../constants/env';
import Key from './model';

export const createDefaultKeys = async () => {
  await Key.create({ type: 'admin', key: appEnv.defaultAdminKey }).catch(() => {
    //
  });
  await Key.create({ type: 'readonly', key: appEnv.defaultReadonlyKey }).catch(
    () => {
      //
    }
  );
  await Key.create({
    type: 'writeonly',
    key: appEnv.defaultWriteonlyKey,
  }).catch(() => {
    //
  });
};
