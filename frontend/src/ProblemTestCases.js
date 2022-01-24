// @flow

import type { ProblemTestCases_problem$key } from "./__generated__/ProblemTestCases_problem.graphql";
import type { ProblemTestCases_viewer$key } from "./__generated__/ProblemTestCases_viewer.graphql";
import type { ProblemTestCasesAddCaseMutation } from "./__generated__/ProblemTestCasesAddCaseMutation.graphql";

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useFragment, useMutation } from "react-relay/hooks";
import LoadingButton from "./LoadingButton";
import ProblemTestCase from "./ProblemTestCase";

type Props = {
  viewer: ProblemTestCases_viewer$key,
  problem: ProblemTestCases_problem$key,
};

export default function ProblemTestCases(props: Props): React.Node {
  const viewer = useFragment(
    graphql`
      fragment ProblemTestCases_viewer on User {
        is_staff
        ...ProblemTestCase_viewer
      }
    `,
    props.viewer
  );

  const problem = useFragment(
    graphql`
      fragment ProblemTestCases_problem on CodingProblem {
        id
        test_cases {
          id
          ...ProblemTestCase_testCase
        }
      }
    `,
    props.problem
  );

  const [addTestCase, isAddingTestCase] =
    useMutation<ProblemTestCasesAddCaseMutation>(
      graphql`
        mutation ProblemTestCasesAddCaseMutation($input: CreateTestCaseInput!) {
          add_test_case(input: $input) {
            coding_problem {
              ...ProblemTestCases_problem
            }
          }
        }
      `
    );

  const handleAddNewCase = () => {
    addTestCase({
      variables: { input: { problem_id: problem.id } },
      onError: (e) => alert(e),
    });
  };

  return (
    <>
      <h3>
        Test Cases{" "}
        {viewer.is_staff && (
          <LoadingButton
            onClick={handleAddNewCase}
            isUpdating={isAddingTestCase}
            size="sm"
          >
            New Case
          </LoadingButton>
        )}
      </h3>
      {problem.test_cases.map((testCase) => (
        <ProblemTestCase
          key={testCase.id}
          viewer={viewer}
          testCase={testCase}
        />
      ))}
    </>
  );
}
