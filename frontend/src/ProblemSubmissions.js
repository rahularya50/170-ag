// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useMemo } from "react";
import { Table } from "react-bootstrap";
import { useFragment } from "react-relay/hooks";
import { generatePath, Link } from "react-router-dom";

import type { ProblemSubmissions_problem$key } from "./__generated__/ProblemSubmissions_problem.graphql";
import ProblemSubmissionsRefetchQuery from "./__generated__/ProblemSubmissionsRefetchQuery.graphql";
import { useRefreshingQuery } from "./useRefreshingQuery";

type Props = {
  problem: ProblemSubmissions_problem$key,
};

export default function ProblemSubmissions(props: Props): React.Node {
  const { id, my_submissions } = useFragment(
    graphql`
      fragment ProblemSubmissions_problem on CodingProblem
      @refetchable(queryName: "ProblemSubmissionsRefetchQuery") {
        id
        my_submissions(is_validation: false) {
          edges {
            node {
              id
              status
              create_time
              points
            }
          }
        }
      }
    `,
    props.problem
  );

  useRefreshingQuery(
    ProblemSubmissionsRefetchQuery,
    useMemo(() => ({ id }), [id])
  );

  return (
    <Table hover>
      <thead>
        <tr>
          <th>Submission</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {my_submissions.edges.map(({ node: submission }) => (
          <tr key={submission.id}>
            <td>
              <Link to={generatePath("/submission/:id", { id: submission.id })}>
                {Intl.DateTimeFormat("en-US", {
                  hour: "numeric",
                  minute: "numeric",
                  day: "numeric",
                  month: "numeric",
                  year: "numeric",
                }).format(new Date(submission.create_time))}
              </Link>
            </td>
            <td>
              {submission.points == null ? (
                submission.status
              ) : (
                <>{submission.points} Points</>
              )}
            </td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
}
