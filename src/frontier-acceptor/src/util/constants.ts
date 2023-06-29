const requireProcessEnv = (key: string) => {
  const value = process.env[key];
  if (!value) {
    throw new Error(`Process env not found: "${key}"`);
  }
  return value;
};

const appEnv = {
  logsUrl: "http://logs:8080/logs",
  appEnv: requireProcessEnv("APP_ENV"),
  logsAppName: requireProcessEnv("LOGS_APP_NAME"),
  logsApiKey: requireProcessEnv("LOGS_KEY"),
};

export default appEnv;
