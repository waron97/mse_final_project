const requireProcessEnv = (key: string) => {
  const value = process.env[key];
  if (!value) {
    throw new Error(`Process env not found: "${key}"`);
  }
  return value;
};

const appEnv = {
  mongoUri: requireProcessEnv('MONGO_URI'),
  port: requireProcessEnv('APP_PORT'),
  defaultAdminKey: requireProcessEnv('DEFAULT_ADMIN_KEY'),
  defaultReadonlyKey: requireProcessEnv('DEFAULT_READONLY_KEY'),
  defaultWriteonlyKey: requireProcessEnv('DEFAULT_WRITEONLY_KEY'),
  lifetimeDaysDebug: parseInt(requireProcessEnv('LIFETIME_DAYS_DEBUG')),
  lifetimeDaysInfo: parseInt(requireProcessEnv('LIFETIME_DAYS_INFO')),
  lifetimeDaysWarning: parseInt(requireProcessEnv('LIFETIME_DAYS_WARNING')),
  lifetimeDaysError: parseInt(requireProcessEnv('LIFETIME_DAYS_ERROR')),
  lifetimeDaysCritical: parseInt(requireProcessEnv('LIFETIME_DAYS_CRITICAL')),
  appEnv: requireProcessEnv('APP_ENV'),
};

export default appEnv;
