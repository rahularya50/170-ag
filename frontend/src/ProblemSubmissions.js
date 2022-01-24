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
    <>
      {my_submissions.edges.map(({ node: submission }) => (
        <p key={submission.id}>Submission (status={submission.status})</p>
      ))}
    </>
  );
}
