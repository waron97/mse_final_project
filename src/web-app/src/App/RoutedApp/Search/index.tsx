import { Button, Card, Input, Pagination, Skeleton } from "antd";
import classNames from "classnames";
import { FC, useState } from "react";
import { AiOutlineSearch } from "react-icons/ai";
import { useQuery } from "react-query";
import { useLocation } from "react-router-dom";
import styled from "styled-components";

import { Theme } from "@theme";

import { getSearchResult } from "../../../shared/api";
import SearchResult from "../../../shared/components/SearchResult";

// ----------------------------------------------------------------------------

interface SearchProps {
  className?: string;
}

const _Search: FC<SearchProps> = (props) => {
  // -------------------------------------
  // Props destructuring
  // -------------------------------------

  const { className } = props;

  // -------------------------------------
  // Hooks (e.g. useState, useMemo ...)
  // -------------------------------------

  const queryPrams = useLocation().search;
  const initialQuery = new URLSearchParams(queryPrams).get("query") || "";

  const [input, setInput] = useState<string>(initialQuery);
  const [query, setQuery] = useState<string>(initialQuery);

  const [page, setPage] = useState<number>(1);

  const { data, isFetching } = useQuery({
    queryKey: ["search", query, page],
    queryFn: () => getSearchResult(query, page),
    keepPreviousData: true,
    enabled: !!query,
  });

  console.log(isFetching);

  // -------------------------------------
  // Effects
  // -------------------------------------

  // -------------------------------------
  // Component functions
  // -------------------------------------

  function handleSearch(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setQuery(input);
    setPage(1);
  }

  function renderSearchForm() {
    return (
      <form onSubmit={handleSearch} className="search-form">
        <Input
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Input your query"
        ></Input>
        <Button type="primary" shape="circle" htmlType="submit">
          <AiOutlineSearch />
        </Button>
      </form>
    );
  }

  function renderLoader() {
    return (
      <div className="results">
        {Array(10)
          .fill(true)
          .map((_, index) => {
            return <Skeleton key={index} active loading />;
          })}
      </div>
    );
  }

  function renderResults() {
    if (isFetching) {
      return renderLoader();
    }
    return data?.data.map((document) => {
      return <SearchResult document={document} />;
    });
  }

  // -------------------------------------
  // Component local variables
  // -------------------------------------

  return (
    <div className={classNames([className])}>
      <div className="main">
        <Card className="form-card">{renderSearchForm()}</Card>
        <Card className="results-card">
          <div className="results">{renderResults()}</div>
          <div className="pagination">
            <Pagination
              current={page}
              total={data?.meta?.total}
              onChange={(pag) => setPage(pag)}
              showSizeChanger={false}
            />
          </div>
        </Card>
      </div>
    </div>
  );
};

// ----------------------------------------------------------------------------

const Search = styled(_Search)<Theme>`
  & {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 50px 0px;
    min-height: 100vh;
    overflow: hidden;
    flex-direction: column;

    .main {
      max-width: 600px;
      width: 100%;
      min-height: 100%;
      flex: 1;
      display: flex;
      flex-direction: column;

      .search-form {
        display: flex;
        gap: 12px;
        align-items: center;

        svg {
          font-size: 20px;
        }
      }

      .results-card {
        margin-top: 24px;
        flex: 1;
        .results {
          display: flex;
          flex-direction: column;
          gap: 30px;
        }
        .pagination {
          display: flex;
          justify-content: center;
          margin-top: 24px;
          width: 100%;
        }
      }
    }
  }
`;

export default Search;
