export interface FrontierPage {
  url: string;
  priority: number;
}

export type FrontierDocument = FrontierPage & {
  _id: string;
};
