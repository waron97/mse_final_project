import classNames from "classnames";
import { FC } from "react";
import { HashRouter, Route, Routes } from "react-router-dom";
import styled from "styled-components";

import { Theme } from "@theme";

import Landing from "./Home";
import Search from "./Search";

// ----------------------------------------------------------------------------

interface RoutedAppProps {
  className?: string;
}

const _RoutedApp: FC<RoutedAppProps> = (props) => {
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

  return (
    <div className={classNames([className])}>
      <HashRouter>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/search" element={<Search />} />
        </Routes>
      </HashRouter>
    </div>
  );
};

// ----------------------------------------------------------------------------

const RoutedApp = styled(_RoutedApp)<Theme>`
  & {
    width: 100vw;
    min-height: 100vh;
  }
`;

export default RoutedApp;
