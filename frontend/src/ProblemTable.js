// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useFragment } from "react-relay/hooks";
import { Table } from "react-bootstrap";
import { generatePath, Link } from "react-router-dom";
import type { ProblemTable$key } from "./__generated__/ProblemTable.graphql";

type Props = {
  query: ProblemTable$key,
};

export default function ProblemTable(props: Props): React.Node {
  const { coding_problems } = useFragment(
    graphql`
      fragment ProblemTable on Query {
        coding_problems(include_unreleased: $include_unreleased) {
          edges {
            node {
              id
              name
            }
          }
        }
      }
    `,
    props.query
  );

  return (
    <Table hover>
      <thead>
        <tr>
          <th>Name</th>
        </tr>
      </thead>
      <tbody>
        {coding_problems.edges.map(({ node }) => (
          <tr key={node.id}>
            <td>
              <Link to={generatePath("/problem/:id", { id: node.id })}>
                {node.name}
              </Link>
            </td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
}
