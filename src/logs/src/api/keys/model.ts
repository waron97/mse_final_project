import { Document, Schema, model } from 'mongoose';

export type KeyType = 'readonly' | 'writeonly' | 'admin';

export interface IKey {
  key: string;
  type: KeyType;
}

export type KeyDocument = IKey & Document;

const keySchema = new Schema<IKey>({
  key: {
    type: String,
    required: true,
    unique: true,
  },
  type: {
    type: String,
    required: true,
    enum: ['readonly', 'writeonly', 'admin'],
  },
});

const Key = model<IKey>('Key', keySchema);

export default Key;
