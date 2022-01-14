// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Table } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";
import { generatePath, Link } from "react-router-dom";

export default function Problems(): React.Node {
  const { coding_problems } = useLazyLoadQuery(
    graphql`
      query ProblemsQuery {
        coding_problems {
          edges {
            node {
              id
              name
            }
          }
        }
      }
    `,
    {}
  );
  return (
    <>
      <h3>Problems</h3>
      <Table hover>
        <thead>
          <tr>
            <th>Name</th>
          </tr>
        </thead>
        <tbody>
          {coding_problems.edges.map(({ node }) => (
            <tr>
              <td>
                <Link to={generatePath("/problem/:id", { id: node.id })}>
                  {node.name}
                </Link>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </>
  );
}
