// @flow

import * as React from "react";
import { useCallback, useState } from "react";
import graphql from "babel-plugin-relay/macro";
import {
  fetchQuery,
  useFragment,
  useRelayEnvironment,
} from "react-relay/hooks";
import useInterval from "use-interval";
import type { ProblemSubmissions_problem$key } from "./__generated__/ProblemSubmissions_problem.graphql";
import ProblemSubmissionsRefetchQuery from "./__generated__/ProblemSubmissionsRefetchQuery.graphql";
import { Table } from "react-bootstrap";

type Props = {
  problem: ProblemSubmissions_problem$key,
};

export default function ProblemSubmissions(props: Props): React.Node {
  const { id, my_submissions } = useFragment(
    graphql`
      fragment ProblemSubmissions_problem on CodingProblem
      @refetchable(queryName: "ProblemSubmissionsRefetchQuery") {
        id
        my_submissions {
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

  const environment = useRelayEnvironment();
  const [isRefreshing, setIsRefreshing] = useState(false);
  const refresh = useCallback(() => {
    if (isRefreshing) {
      return;
    }
    setIsRefreshing(true);
    fetchQuery(environment, ProblemSubmissionsRefetchQuery, { id }).subscribe({
      complete: () => {
        setIsRefreshing(false);
      },
      error: () => setIsRefreshing(false),
    });
  }, [environment, isRefreshing, id]);

  useInterval(refresh, 5000);

  return (
    <Table hover>
      <thead>
        <tr>
          <th>Submission Time</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {my_submissions.edges.map(({ node: submission }) => (
          <tr key={submission.id}>
            <td>
              {Intl.DateTimeFormat("en-US", {
                hour: "numeric",
                minute: "numeric",
                day: "numeric",
                month: "numeric",
                year: "numeric",
              }).format(new Date(submission.create_time))}
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
