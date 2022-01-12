// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Navigate, useParams } from "react-router-dom";
import { Link } from "react-router-dom";

export default function Problems(): React.Node {
  const { id } = useParams();

  const { coding_problem } = useLazyLoadQuery(
    graphql`
      query ProblemQuery($id: ID!) {
        coding_problem(id: $id) {
          name
          statement
        }
      }
    `,
    { id }
  );

  if (!coding_problem) {
    return <Navigate to="404" />;
  }

  return (
    <>
      <Link to="/problems/">(Back to all problems)</Link>
      <h3>{coding_problem.name}</h3>
      {coding_problem.statement}
    </>
  );
}
