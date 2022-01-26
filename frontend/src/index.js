// @flow

import "bootstrap/dist/css/bootstrap.min.css";

import React, { Suspense } from "react";
import ReactDOM from "react-dom";
import { RelayEnvironmentProvider } from "react-relay";
import { BrowserRouter } from "react-router-dom";

import App from "./App";
import RelayEnvironment from "./RelayEnvironment";

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
