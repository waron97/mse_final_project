import appEnv from "./constants";

export enum LogLevel {
  DEBUG = "debug",
  INFO = "info",
  WARNING = "warning",
  ERROR = "error",
  CRITICAL = "critical",
}

type AnyObject = { [key: string]: any };

interface LogParams {
  location: string;
  message: string;
  detail?: AnyObject;
}

function createLog(params: LogParams & { level: LogLevel }) {
  const body = {
    appId: appEnv.logsAppName,
    level: params.level,
    location: params.location,
    message: params.message,
    detail: params.detail,
  };
  return fetch(appEnv.logsUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `apiKey ${appEnv.logsApiKey}`,
    },
    body: JSON.stringify(body),
  })
    .then((res) => {
      if (res.ok) {
        return true;
      } else {
        return false;
      }
    })
    .catch(() => {
      return false;
    });
}

const Log = {
  debug: (location: string, message: string, data?: AnyObject) =>
    createLog({ level: LogLevel.DEBUG, location, message, detail: data }),
  info: (location: string, message: string, data?: AnyObject) =>
    createLog({ level: LogLevel.INFO, location, message, detail: data }),
  warning: (location: string, message: string, data?: AnyObject) =>
    createLog({ level: LogLevel.WARNING, location, message, detail: data }),
  error: (location: string, message: string, data?: AnyObject) =>
    createLog({ level: LogLevel.ERROR, location, message, detail: data }),
  critical: (location: string, message: string, data?: AnyObject) =>
    createLog({ level: LogLevel.CRITICAL, location, message, detail: data }),
};

export default Log;
