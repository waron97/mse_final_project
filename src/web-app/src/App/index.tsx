import classNames from "classnames";
import { FC } from "react";
import styled from "styled-components";

import RoutedApp from "./RoutedApp";

// ----------------------------------------------------------------------------

interface AppProps {
  className?: string;
}

const _App: FC<AppProps> = (props) => {
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
      <RoutedApp />
    </div>
  );
};

// ----------------------------------------------------------------------------

const App = styled(_App)`
  & {
  }
`;

export default App;
