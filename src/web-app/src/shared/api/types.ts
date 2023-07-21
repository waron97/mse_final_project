export type SearchResult = {
  documentId: string;
  documentScore: number;
  bestPassageId: string;
  documentTitle: string;
  bestPassageText: string;
  documentUrl: string;
  documentDescription: string;
};

export interface Paginated<T> {
  data: T[];
  meta: {
    page: number;
    limit: number;
    total: number;
  };
}
