// @flow

import type {
  AppQuery$variables,
  AppQuery$data,
} from "./__generated__/AppQuery.graphql";

import * as React from "react";
import { graphql, useLazyLoadQuery } from "react-relay";

export default function App(): React.Node {
  const data = useLazyLoadQuery<AppQuery$variables, AppQuery$data>(
    graphql`
      query AppQuery($ids: [ID!]!) {
        nodes(ids: $ids) {
          id
          ... on User {
            age
          }
        }
      }
    `,
    { ids: ["1", "2"] }
  );
  return <div className="App">{data.nodes[0].age}</div>;
}
