// @flow
import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import { Form } from "react-bootstrap";
import { useLazyLoadQuery, useMutation } from "react-relay/hooks";
import { useNavigate } from "react-router-dom";

import type { ProblemsCreateMutation } from "./__generated__/ProblemsCreateMutation.graphql";
import LoadingButton from "./LoadingButton";
import ProblemTable from "./ProblemTable";

export default function Problems(): React.Node {
  const [includeUnreleased, setIncludeUnreleased] = useState(false);

  const query = useLazyLoadQuery(
    graphql`
      query ProblemsQuery($include_unreleased: Boolean!) {
        viewer {
          is_staff
        }
        ...ProblemTable @arguments(includeUnreleased: $include_unreleased)
      }
    `,
    { include_unreleased: includeUnreleased }
  );

  const [createProblem, isCreatingProblem] =
    useMutation<ProblemsCreateMutation>(
      graphql`
        mutation ProblemsCreateMutation($input: CodingProblemInput!) {
          new_problem(input: $input) {
            id
          }
        }
      `
    );

  const navigate = useNavigate();

  const handleCreateProblem = () => {
    createProblem({
      variables: {
        input: {
          name: "New Problem",
          statement: "Problem Statement",
          released: false,
        },
      },
      onCompleted: (resp) => {
        navigate(`/problem/${resp.new_problem.id}`);
      },
    });
  };

  return (
    <>
      <h3>Problems</h3>
      {query.viewer?.is_staff && (
        <Form.Group className="mb-3" controlId="unreleasedCheckbox">
          <Form.Check
            type="checkbox"
            label="Include unreleased"
            checked={includeUnreleased}
            onChange={(e) => setIncludeUnreleased(e.target.checked)}
          />
        </Form.Group>
      )}
      <React.Suspense>
        <ProblemTable query={query} />
      </React.Suspense>
      {query.viewer?.is_staff && (
        <LoadingButton
          onClick={handleCreateProblem}
          isUpdating={isCreatingProblem}
        >
          New Problem
        </LoadingButton>
      )}
    </>
  );
}
