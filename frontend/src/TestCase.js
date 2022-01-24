// @flow

import type { TestCase_viewer$key } from "./__generated__/TestCase_viewer.graphql";
import type { TestCase_testCase$key } from "./__generated__/TestCase_testCase.graphql";
import type { TestCaseUpdateMutation } from "./__generated__/TestCaseUpdateMutation.graphql";

import * as React from "react";
import { useState } from "react";
import graphql from "babel-plugin-relay/macro";
import { useFragment, useMutation } from "react-relay/hooks";
import { Button, Form, Spinner } from "react-bootstrap";

type Props = {
  viewer: TestCase_viewer$key,
  testCase: TestCase_testCase$key,
};

export default function TestCase(props: Props): React.Node {
  const { is_staff: isStaff } = useFragment(
    graphql`
      fragment TestCase_viewer on User {
        is_staff
      }
    `,
    props.viewer
  );

  const testCase = useFragment(
    graphql`
      fragment TestCase_testCase on CodingTestCase {
        id
        points
        public
        public_data {
          input
          output
        }
      }
    `,
    props.testCase
  );

  const [updateTestCase, isUpdating] =
    useMutation<TestCaseUpdateMutation>(graphql`
      mutation TestCaseUpdateMutation($input: UpdateTestCaseInput!) {
        update_test_case(input: $input) {
          ...TestCase_testCase
        }
      }
    `);

  const [isEditing, setIsEditing] = useState(false);

  const [points, setPoints] = useState(testCase.points);
  const [isPublic, setPublic] = useState(testCase.public);

  const handleSave = () => {
    updateTestCase({
      variables: {
        input: {
          id: testCase.id,
          points,
          public: isPublic,
        },
      },
      onCompleted: () => {
        setIsEditing(false);
      },
      onError: (err) => {
        alert(err);
      },
    });
  };

  return (
    <div>
      {isEditing ? (
        <div>
          <Form.Group className="mb-3">
            <Form.Label>Points:</Form.Label>
            <Form.Control
              type="text"
              value={points}
              disabled={isUpdating}
              onChange={(e) => setPoints(e.target.value)}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="releaseCheckbox">
            <Form.Check
              type="checkbox"
              label="Public Test Case"
              checked={isPublic}
              onChange={(e) => setPublic(e.target.checked)}
            />
          </Form.Group>
        </div>
      ) : (
        <div>
          <h6>Case XYZ: {testCase.points} Points</h6>
          {testCase.public_data ? (
            <div>
              Input:
              <pre>{testCase.public_data.input}</pre>
              Expected output:
              <pre>{testCase.public_data.input}</pre>
            </div>
          ) : (
            <div>(Inputs hidden)</div>
          )}
        </div>
      )}
      <p>
        {isStaff &&
          (isEditing ? (
            <Button onClick={handleSave} disabled={isUpdating}>
              {isUpdating ? (
                <Spinner
                  as="span"
                  animation="border"
                  size="sm"
                  role="status"
                  aria-hidden="true"
                />
              ) : (
                "Save"
              )}
            </Button>
          ) : (
            <Button onClick={() => setIsEditing(true)}>Edit</Button>
          ))}
      </p>
    </div>
  );
}
