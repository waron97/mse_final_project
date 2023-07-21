import { Paginated, SearchResult } from "./types";

export function getBaseUrl() {
  return "http://localhost:3005";
}

export function getSearchResult(
  query: string,
  page?: number
): Promise<Paginated<SearchResult>> {
  page = page || 1;
  const limit = 20;
  const relativeUrl = `/rank?query=${encodeURIComponent(
    query
  )}&page=${page}&limit=${limit}`;
  const url = getBaseUrl() + relativeUrl;
  return fetch(url).then((res) => res.json());
}
