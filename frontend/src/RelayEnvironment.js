// @flow

import { Environment, Network, RecordSource, Store } from "relay-runtime";

function fetchQuery(operation, variables) {
  const controller = new AbortController();
  setTimeout(() => controller.abort(), 5000);
  return fetch("/query", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      query: operation.text,
      variables,
    }),
    signal: controller.signal,
  }).then((response) => {
    return response.json();
  });
}

// Export a singleton instance of Relay Environment configured with our network function:
const environment: Environment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource()),
});

export default environment;
