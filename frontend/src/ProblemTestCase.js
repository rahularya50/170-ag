// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import { Button, ButtonGroup, Form, Modal, Stack } from "react-bootstrap";
import { useFragment, useMutation } from "react-relay/hooks";

import type { ProblemTestCase_testCase$key } from "./__generated__/ProblemTestCase_testCase.graphql";
import type { ProblemTestCase_viewer$key } from "./__generated__/ProblemTestCase_viewer.graphql";
import type { ProblemTestCaseInputQuery } from "./__generated__/ProblemTestCaseInputQuery.graphql";
import type { ProblemTestCaseOutputQuery } from "./__generated__/ProblemTestCaseOutputQuery.graphql";
import type { ProblemTestCaseUpdateMutation } from "./__generated__/ProblemTestCaseUpdateMutation.graphql";
import downloadText from "./downloadText";
import LoadingButton from "./LoadingButton";
import ProblemTestDataEditor from "./ProblemTestDataEditor";
import { useQueryFetcher } from "./useQueryFetcher";

type Props = {
  name: string,
  viewer: ProblemTestCase_viewer$key,
  testCase: ProblemTestCase_testCase$key,
};

export default function TestCase(props: Props): React.Node {
  const { is_staff: isStaff } = useFragment(
    graphql`
      fragment ProblemTestCase_viewer on User {
        is_staff
      }
    `,
    props.viewer
  );

  const testCase = useFragment(
    graphql`
      fragment ProblemTestCase_testCase on CodingTestCase {
        id
        points
        visibility
        expanded_data {
          input
          output
        }
        ...ProblemTestDataEditor_testCase
      }
    `,
    props.testCase
  );

  const [updateTestCase, isUpdating] =
    useMutation<ProblemTestCaseUpdateMutation>(graphql`
      mutation ProblemTestCaseUpdateMutation($input: UpdateTestCaseInput!) {
        update_test_case(input: $input) {
          ...ProblemTestCase_testCase
        }
      }
    `);

  const [deleteTestCase, isDeleting] =
    useMutation<ProblemTestCaseUpdateMutation>(graphql`
      mutation ProblemTestCaseDeleteMutation($input: DeleteTestCaseInput!) {
        delete_test_case(input: $input) {
          ...ProblemTestCases_problem
        }
      }
    `);

  const [isEditing, setIsEditing] = useState(false);

  const [points, setPoints] = useState(testCase.points);
  const [visibility, setVisibility] = useState(testCase.visibility);

  const [showDataEditor, setShowDataEditor] = useState(false);
  const [input, setInput] = useState(null);
  const [output, setOutput] = useState(null);

  const handleSave = () => {
    updateTestCase({
      variables: {
        input: {
          id: testCase.id,
          points,
          visibility,
          input,
          output,
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

  const handleDelete = () => {
    deleteTestCase({
      variables: {
        input: {
          id: testCase.id,
        },
      },
      onError: (err) => {
        alert(err);
      },
    });
  };

  const [downloadInput, isDownloadingInput] =
    useQueryFetcher<ProblemTestCaseInputQuery>(
      graphql`
        query ProblemTestCaseInputQuery($id: ID!) {
          test_case(id: $id) {
            data {
              input
            }
          }
        }
      `,
      { id: testCase.id },
      (resp) => {
        downloadText(resp.test_case.data.input, "input.txt");
      },
      alert
    );

  const [downloadOutput, isDownloadingOutput] =
    useQueryFetcher<ProblemTestCaseOutputQuery>(
      graphql`
        query ProblemTestCaseOutputQuery($id: ID!) {
          test_case(id: $id) {
            data {
              output
            }
          }
        }
      `,
      { id: testCase.id },
      (resp) => {
        downloadText(resp.test_case.data.output, "output.txt");
      },
      alert
    );

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
            <Form.Select
              label="Test Case Visibility"
              value={visibility}
              onChange={(e) => setVisibility(e.target.value)}
            >
              <option value="PRIVATE">Private</option>
              <option value="COLLAPSED">Collapsed</option>
              <option value="EXPANDED">Expanded</option>
            </Form.Select>
          </Form.Group>
        </div>
      ) : (
        <div>
          <h6>
            Case {props.name}: {testCase.points} Points
          </h6>
          <div className="mb-2">
            {testCase.expanded_data ? (
              <div>
                Input:
                <pre>{testCase.expanded_data.input}</pre>
                Expected output:
                <pre>{testCase.expanded_data.output}</pre>
              </div>
            ) : testCase.visibility === "PRIVATE" ? (
              <div>(Inputs hidden)</div>
            ) : (
              <ButtonGroup>
                <LoadingButton
                  size="sm"
                  variant="secondary"
                  onClick={downloadInput}
                  isUpdating={isDownloadingInput}
                >
                  Download Input
                </LoadingButton>
                <LoadingButton
                  size="sm"
                  variant="secondary"
                  onClick={downloadOutput}
                  isUpdating={isDownloadingOutput}
                >
                  Download Output
                </LoadingButton>
              </ButtonGroup>
            )}
          </div>
        </div>
      )}
      <p>
        {isStaff &&
          (isEditing ? (
            <Stack direction="horizontal" gap={1}>
              <LoadingButton
                isUpdating={isUpdating}
                onClick={handleSave}
                size="sm"
              >
                Save
              </LoadingButton>
              <Button onClick={() => setShowDataEditor(true)} size="sm">
                Edit Data
              </Button>
              <LoadingButton
                isUpdating={isDeleting}
                onClick={handleDelete}
                size="sm"
                variant="danger"
              >
                Delete
              </LoadingButton>
            </Stack>
          ) : (
            <Button onClick={() => setIsEditing(true)} size="sm">
              Edit
            </Button>
          ))}
      </p>
      <Modal
        show={showDataEditor}
        onHide={() => setShowDataEditor(false)}
        size="lg"
      >
        <Modal.Body>
          {showDataEditor && (
            <React.Suspense fallback={null}>
              <ProblemTestDataEditor
                testCase={testCase}
                input={input}
                onInputChange={setInput}
                output={output}
                onOutputChange={setOutput}
              />
            </React.Suspense>
          )}
        </Modal.Body>
      </Modal>
    </div>
  );
}
