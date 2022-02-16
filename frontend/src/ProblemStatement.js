// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import { Button, Col, Form, Row } from "react-bootstrap";
import ReactMarkdown from "react-markdown";
import { useFragment, useMutation } from "react-relay/hooks";

import type { ProblemStatement_problem$key } from "./__generated__/ProblemStatement_problem.graphql";
import type { ProblemStatement_viewer$key } from "./__generated__/ProblemStatement_viewer.graphql";
import type { ProblemStatementUpdateMutation } from "./__generated__/ProblemStatementUpdateMutation.graphql";
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
        deadline
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
  const [deadline, setDeadline] = useState(problem.deadline.slice(0, -1));
  const [statement, setStatement] = useState(problem.statement);
  const [released, setReleased] = useState(problem.released);

  const handleSave = () => {
    updateProblem({
      variables: {
        input: {
          id: problem.id,
          problem: {
            name,
            deadline: new Date(deadline + "Z").toISOString(),
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
        <>
          <Form.Group className="mb-3">
            <Form.Label>Problem Title</Form.Label>
            <Form.Control
              type="text"
              value={name}
              disabled={isUpdating}
              onChange={(e) => setName(e.target.value)}
            />
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Label>Deadline (UTC)</Form.Label>
            <Form.Control
              type="datetime-local"
              value={deadline}
              disabled={isUpdating}
              onChange={(e) => setDeadline(e.target.value)}
              step={1}
            />
          </Form.Group>
        </>
      ) : (
        <div className="mb-3">
          <h3>{problem.name}</h3>
          <i>
            Due by{" "}
            {Intl.DateTimeFormat("en-US", {
              hour: "numeric",
              minute: "numeric",
              second: "numeric",
              day: "numeric",
              month: "numeric",
              year: "numeric",
            }).format(new Date(problem.deadline))}
          </i>
        </div>
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
