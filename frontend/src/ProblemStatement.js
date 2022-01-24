// @flow

import type { ProblemStatement_viewer$key } from "./__generated__/ProblemStatement_viewer.graphql";
import type { ProblemStatement_problem$key } from "./__generated__/ProblemStatement_problem.graphql";
import type { ProblemStatementUpdateMutation } from "./__generated__/ProblemStatementUpdateMutation.graphql";

import * as React from "react";
import { useState } from "react";
import graphql from "babel-plugin-relay/macro";
import { useFragment, useMutation } from "react-relay/hooks";
import { Button, Col, Form, Row } from "react-bootstrap";
import ReactMarkdown from "react-markdown";
import LoadingButton from "./LoadingButton";

type Props = {
  viewer: ProblemStatement_viewer$key,
  problem: ProblemStatement_problem$key,
};

export default function ProblemStatement(props: Props): React.Node {
  const { is_staff: isStaff } = useFragment(
    graphql`
      fragment ProblemStatement_viewer on User {
        is_staff
      }
    `,
    props.viewer
  );

  const problem = useFragment(
    graphql`
      fragment ProblemStatement_problem on CodingProblem {
        id
        name
        statement
        released
      }
    `,
    props.problem
  );

  const [updateProblem, isUpdating] =
    useMutation<ProblemStatementUpdateMutation>(graphql`
      mutation ProblemStatementUpdateMutation(
        $input: UpdateCodingProblemInput!
      ) {
        update_problem(input: $input) {
          ...ProblemStatement_problem
        }
      }
    `);

  const [isEditing, setIsEditing] = useState(false);

  const [name, setName] = useState(problem.name);
  const [statement, setStatement] = useState(problem.statement);
  const [released, setReleased] = useState(problem.released);

  const handleSave = () => {
    updateProblem({
      variables: {
        input: {
          id: problem.id,
          problem: {
            name,
            statement,
            released,
          },
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
    <>
      {isEditing ? (
        <Form.Group className="mb-3">
          <Form.Label>Problem Title</Form.Label>
          <Form.Control
            type="text"
            value={name}
            disabled={isUpdating}
            onChange={(e) => setName(e.target.value)}
          />
        </Form.Group>
      ) : (
        <h3>{problem.name}</h3>
      )}
      {isEditing ? (
        <Form.Group className="mb-3">
          <Form.Label>Problem Statement</Form.Label>
          <Row>
            <Col>
              <Form.Control
                as="textarea"
                value={statement}
                disabled={isUpdating}
                onChange={(e) => setStatement(e.target.value)}
                style={{ minHeight: "100%" }}
              />
            </Col>
            <Col>
              <ReactMarkdown>{statement}</ReactMarkdown>
            </Col>
          </Row>
        </Form.Group>
      ) : (
        <ReactMarkdown>{problem.statement}</ReactMarkdown>
      )}
      {isEditing && (
        <Form.Group className="mb-3" controlId="releaseCheckbox">
          <Form.Check
            type="checkbox"
            label="Release Assignment"
            checked={released}
            onChange={(e) => setReleased(e.target.checked)}
          />
        </Form.Group>
      )}
      <p>
        {isStaff &&
          (isEditing ? (
            <LoadingButton onClick={handleSave} isUpdating={isUpdating}>
              Save
            </LoadingButton>
          ) : (
            <Button onClick={() => setIsEditing(true)} size="sm">
              Edit
            </Button>
          ))}
      </p>
    </>
  );
}
