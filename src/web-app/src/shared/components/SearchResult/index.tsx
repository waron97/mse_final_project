import { Typography } from "antd";
import classNames from "classnames";
import { FC } from "react";
import styled from "styled-components";

import { Theme } from "@theme";

import { SearchResult as SR } from "../../api/types";

const { Link, Paragraph } = Typography;

// ----------------------------------------------------------------------------

interface SearchResultProps {
  className?: string;
  document: SR;
}

const _SearchResult: FC<SearchResultProps> = (props) => {
  // -------------------------------------
  // Props destructuring
  // -------------------------------------

  const { className, document } = props;

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

  return (
    <div className={classNames([className])}>
      <Link
        className="doc-title"
        href={document.documentUrl}
        target="_blank"
        rel="noreferrer"
      >
        {document.documentTitle}
      </Link>
      <Link
        href={document.documentUrl}
        target="_blank"
        rel="noreferrer"
        className="doc-url"
      >
        {document.documentUrl}
      </Link>
      <Paragraph className="doc-text">
        {document.documentDescription || document.bestPassageText}
      </Paragraph>
    </div>
  );
};

// ----------------------------------------------------------------------------

const SearchResult = styled(_SearchResult)<Theme>`
  & {
    .doc-title {
      font-size: 16px;
      margin-bottom: 0px;
      padding: 0;
      display: block;
      color: ${({ theme }) => theme.colors.textPrimary};
      font-weight: 600;
    }
    .doc-url {
      font-size: 11px;
      color: ${({ theme }) => theme.colors.textMuted};
      display: block;
    }
    .doc-text {
      margin-top: 6px;
    }
  }
`;

export default SearchResult;
