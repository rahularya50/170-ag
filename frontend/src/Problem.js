// @flow
import type { ProblemSaveDraftMutation } from "./__generated__/ProblemSaveDraftMutation.graphql";

import * as React from "react";
import { useCallback, useEffect, useState } from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery, useMutation } from "react-relay/hooks";
import { Navigate, useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import AceEditor from "react-ace";
import { Button, Col, Container, Row, Stack } from "react-bootstrap";

import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/mode-java";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/theme-github";

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

  const [saveDraft, isSaving] = useMutation<ProblemSaveDraftMutation>(
    graphql`
      mutation ProblemSaveDraftMutation($input: CodingDraftInput!) {
        save_draft(input: $input) {
          code
        }
      }
    `
  );

  const [autoSaveEnabled, setAutoSaveEnabled] = useState(true);
  const [studentCode, setStudentCode] = useState("");
  const [savedStudentCode, setSavedStudentCode] = useState("");

  const isSaved = !isSaving && studentCode === savedStudentCode;

  const forceSave = useCallback(() => {
    if (!isSaving) {
      saveDraft({
        variables: {
          input: { problem_id: id, code: studentCode },
        },
        onCompleted: ({ save_draft }) => {
          setSavedStudentCode(save_draft.code);
          setAutoSaveEnabled(true);
        },
        onError: () => {
          setAutoSaveEnabled(false);
        },
      });
    }
  }, [id, isSaving, saveDraft, studentCode]);

  useEffect(() => {
    if (autoSaveEnabled && !isSaved) {
      forceSave();
    }
  }, [autoSaveEnabled, isSaved, forceSave]);

  if (!coding_problem) {
    return <Navigate to="404" />;
  }

  return (
    <Container>
      <Row>
        <Col>
          <h3>{coding_problem.name}</h3>
          <p>{coding_problem.statement}</p>{" "}
          <Link to="/problems/">(Back to all problems)</Link>
        </Col>
        <Col>
          <Stack gap={2}>
            <AceEditor
              mode="python" // TODO: link syntax highlighting to submit language dropdown
              name="code-input"
              onChange={setStudentCode}
              value={studentCode}
            />
            <Stack direction="horizontal" gap={1}>
              <Button size="sm" variant="secondary" onClick={forceSave}>
                {isSaved
                  ? "Draft Saved"
                  : isSaving
                  ? "Saving..."
                  : "Save Draft"}
              </Button>
              <Button size="sm" variant="primary">
                Submit
              </Button>
            </Stack>
          </Stack>
        </Col>
      </Row>
    </Container>
  );
}
