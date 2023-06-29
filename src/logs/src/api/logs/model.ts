import { Document, Schema, model } from 'mongoose';

export const LogLevels = ['debug', 'info', 'warning', 'error', 'critical'];

export enum LogLevel {
  DEBUG = 'debug',
  INFO = 'info',
  WARNING = 'warning',
  ERROR = 'error',
  CRITICAL = 'critical',
}

type AnyObject = { [key: string]: any };

export interface ILog {
  appId: string;
  level: LogLevel;
  location: string;
  message: string;
  detail: AnyObject;
  date: Date;
}

export type LogDocument = ILog & Document;

const logSchema = new Schema<ILog>({
  appId: {
    type: String,
    required: true,
    index: true,
  },
  level: {
    type: String,
    required: true,
    default: LogLevel.DEBUG,
    enum: LogLevels,
  },
  location: {
    type: String,
  },
  message: {
    type: String,
  },
  detail: {
    type: {},
  },
  date: {
    type: Date,
    default: Date.now,
  },
});

const Log = model<ILog>('Log', logSchema);

export default Log;
