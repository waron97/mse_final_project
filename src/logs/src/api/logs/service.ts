import * as dayjs from 'dayjs';
import * as cron from 'node-cron';

import appEnv from '../../constants/env';
import Log, { LogLevel } from './model';

export const processLogFilters = (params: any) => {
  const { since, appId, levels, text } = params;
  const filters: { [key: string]: any } = {};

  if (since) {
    filters.date = {
      $gt: since,
    };
  }

  if (appId) {
    filters.appId = appId;
  }

  if (levels && levels.length) {
    filters.level = { $in: levels };
  }

  if (text) {
    filters.$or = [
      { location: { $regex: text, $options: 'im' } },
      { message: { $regex: text, $options: 'im' } },
    ];
  }

  return filters;
};

const getLogDeleteTime = (date: Date, level: LogLevel) => {
  switch (level) {
    case LogLevel.DEBUG: {
      return dayjs(date).add(appEnv.lifetimeDaysDebug, 'days');
    }
    case LogLevel.INFO: {
      return dayjs(date).add(appEnv.lifetimeDaysInfo, 'days');
    }
    case LogLevel.WARNING: {
      return dayjs(date).add(appEnv.lifetimeDaysWarning, 'days');
    }
    case LogLevel.ERROR: {
      return dayjs(date).add(appEnv.lifetimeDaysError, 'days');
    }
    case LogLevel.CRITICAL: {
      return dayjs(date).add(appEnv.lifetimeDaysCritical, 'days');
    }
    default: {
      return dayjs(date).add(appEnv.lifetimeDaysCritical, 'days');
    }
  }
};

export const scheduleObsoleteLogsRemoval = () => {
  function removeObsoleteLogs() {
    Log.find({})
      .cursor()
      .eachAsync(async (log) => {
        const { date, level } = log;
        const deleteTime = getLogDeleteTime(date, level);
        const now = dayjs();
        if (now.isAfter(deleteTime)) {
          await log.remove();
        }
      });
  }

  cron.schedule('0 * * * *', removeObsoleteLogs);
};
