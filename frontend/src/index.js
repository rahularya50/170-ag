// @flow

import React, { Suspense } from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { RelayEnvironmentProvider } from "react-relay";
import RelayEnvironment from "./RelayEnvironment";
import { BrowserRouter } from "react-router-dom";

import "bootstrap/dist/css/bootstrap.min.css";

ReactDOM.render(
  <React.StrictMode>
    <RelayEnvironmentProvider environment={RelayEnvironment}>
      <BrowserRouter>
        <Suspense fallback="Loading...">
          <App />
        </Suspense>
      </BrowserRouter>{" "}
    </RelayEnvironmentProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
