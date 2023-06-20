import { WithId } from "mongodb";

import { db } from "..";

export type Crawl = {
  title: string;
  url: string;
  description: string;
  keywords: string[];
  rawHtml: string;
  bodyTextContent: string;
  mainTextContent: string;
  date: Date;
  links: { text: string; href: string }[];
};

export type CrawlDocument = WithId<Crawl>;

export const crawlCollection = db.collection<Crawl>("crawl");
