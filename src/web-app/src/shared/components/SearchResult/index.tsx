import classNames from "classnames";
import { FC } from "react";
import styled from "styled-components";

import { Theme } from "@theme";

import { SearchResult as SR } from "../../api/types";

// ----------------------------------------------------------------------------

interface SearchResultProps {
  className?: string;
  document: SR;
}

const _SearchResult: FC<SearchResultProps> = (props) => {
  // -------------------------------------
  // Props destructuring
  // -------------------------------------

  const { className } = props;

  // -------------------------------------
  // Hooks (e.g. useState, useMemo ...)
  // -------------------------------------

  // -------------------------------------
  // Effects
  // -------------------------------------

  // -------------------------------------
  // Component functions
  // -------------------------------------

  // -------------------------------------
  // Component local variables
  // -------------------------------------

  return <div className={classNames([className])}>SearchResult works!</div>;
};

// ----------------------------------------------------------------------------

const SearchResult = styled(_SearchResult)<Theme>`
  & {
  }
`;

export default SearchResult;
