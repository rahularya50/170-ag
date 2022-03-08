// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useMemo } from "react";
import { Table } from "react-bootstrap";
import { useFragment } from "react-relay/hooks";
import { generatePath, Link } from "react-router-dom";

import type { ProblemValidation_problem$key } from "./__generated__/ProblemValidation_problem.graphql";
import ProblemValidationRefetchQuery from "./__generated__/ProblemValidationRefetchQuery.graphql";
import { useRefreshingQuery } from "./useRefreshingQuery";

type Props = {
  problem: ProblemValidation_problem$key,
};

export default function ProblemValidation(props: Props): React.Node {
  const { id, validation } = useFragment(
    graphql`
      fragment ProblemValidation_problem on CodingProblem
      @refetchable(queryName: "ProblemValidationRefetchQuery") {
        id
        validation: my_submissions(is_validation: true, last: 1) {
          edges {
            node {
              status
              create_time
              staff_data {
                output
                stderr
                exit_error
              }
              results {
                case_results {
                  case_name
                  result
                }
              }
            }
          }
        }
      }
    `,
    props.problem
  );

  useRefreshingQuery(
    ProblemValidationRefetchQuery,
    useMemo(() => ({ id }), [id])
  );

  const edge = validation.edges[0];
  if (edge == null) {
    return null;
  }
  const {
    status,
    create_time,
    staff_data: { output, stderr, exit_error },
    results: { case_results },
  } = edge.node;

  return (
    <div>
      <h3>
        Validation run at{" "}
        {Intl.DateTimeFormat("en-US", {
          hour: "numeric",
          minute: "numeric",
          day: "numeric",
          month: "numeric",
          year: "numeric",
        }).format(new Date(create_time))}
      </h3>
      <p className="lead">Status: {status}</p>
      {output != null && (
        <div>
          Standard Output:
          <pre>{output || "(None)"}</pre>
        </div>
      )}
      {stderr != null && (
        <div>
          Standard Error:
          <pre>{stderr || "(None)"}</pre>
        </div>
      )}
      {exit_error != null && (
        <div>
          Exit Status:
          <pre>{exit_error || "(None)"}</pre>
        </div>
      )}
      {case_results.length > 0 && (
        <Table hover>
          <thead>
            <tr>
              <th>Case</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {case_results.map(({ case_name, result }) => (
              <tr key={case_name}>
                <td>{case_name} (Public)</td>
                <td>{result}</td>
              </tr>
            ))}
          </tbody>
        </Table>
      )}
    </div>
  );
}
