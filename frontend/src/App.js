// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay";

export default function App(): React.Node {
  const data = useLazyLoadQuery(
    graphql`
      query AppQuery($ids: [ID!]!) {
        user(id: "1") {
          name
          age
        }
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
  return (
    <div className="App">
      {data.nodes[0].id}
      User: {data.user?.age}
    </div>
  );
}
