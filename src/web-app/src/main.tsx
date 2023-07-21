import "./index.css";

import React from "react";
import ReactDOM from "react-dom/client";
import { QueryClient, QueryClientProvider } from "react-query";

import App from "./App";
import AppTheme from "./shared/theme";

const client = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={client}>
      <AppTheme>
        <App />
      </AppTheme>
    </QueryClientProvider>
  </React.StrictMode>
);
