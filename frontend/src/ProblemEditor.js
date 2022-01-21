// @flow
import type { ProblemSaveDraftMutation } from "./__generated__/ProblemSaveDraftMutation.graphql";
import type { ProblemEditor_problem$key } from "./__generated__/ProblemEditor_problem.graphql";

import * as React from "react";
import { useCallback, useEffect, useState } from "react";
import graphql from "babel-plugin-relay/macro";
import { useFragment, useMutation } from "react-relay/hooks";
import { useParams } from "react-router-dom";
import AceEditor from "react-ace";
import { Button, Stack } from "react-bootstrap";
import { useDebounced, useDebouncedCallback } from "./useDebounced";

import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/mode-java";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/theme-github";
import type { ProblemEditorSubmitMutation } from "./__generated__/ProblemEditorSubmitMutation.graphql";

type Props = {
  problem: ProblemEditor_problem$key,
};

export default function ProblemEditor(props: Props): React.Node {
  const { id } = useParams();

  const { my_draft: myDraft } = useFragment(
    graphql`
      fragment ProblemEditor_problem on CodingProblem {
        name
        statement
        my_draft {
          code
        }
      }
    `,
    props.problem
  );

  const [saveDraft, isSaving] = useMutation<ProblemSaveDraftMutation>(
    graphql`
      mutation ProblemEditorSaveDraftMutation($input: CodingDraftInput!) {
        save_draft(input: $input) {
          code
        }
      }
    `
  );

  const [submitMutation, isSubmitting] =
    useMutation<ProblemEditorSubmitMutation>(
      graphql`
        mutation ProblemEditorSubmitMutation($input: CodingSubmissionInput!) {
          create_submission(input: $input) {
            coding_problem {
              ...ProblemSubmissions_problem
            }
          }
        }
      `
    );

  const submit = () => {
    submitMutation({
      variables: { input: { problem_id: id, code: studentCode } },
      onCompleted: () => {
        alert("submitted!");
      },
      onError: () => {
        alert("submission failed!");
      },
    });
  };

  const [autoSaveEnabled, setAutoSaveEnabled] = useState(true);
  const [studentCode, setStudentCode] = useState(myDraft?.code ?? "");
  const [savedStudentCode, setSavedStudentCode] = useState("");

  const isSaved = !isSaving && studentCode === savedStudentCode;

  const forceSave = useDebouncedCallback(
    useCallback(
      (code: string) => {
        saveDraft({
          variables: {
            input: { problem_id: id, code },
          },
          onCompleted: ({ save_draft }) => {
            setSavedStudentCode(save_draft.code);
            setAutoSaveEnabled(true);
          },
          onError: () => {
            setAutoSaveEnabled(false);
          },
        });
      },
      [id, saveDraft]
    ),
    1500
  );

  useEffect(() => {
    if (autoSaveEnabled && studentCode !== savedStudentCode) {
      forceSave(studentCode);
    }
  }, [autoSaveEnabled, studentCode, savedStudentCode, forceSave]);

  return (
    <Stack gap={2}>
      <AceEditor
        mode="python" // TODO: link syntax highlighting to submit language dropdown
        name="code-input"
        onChange={setStudentCode}
        value={studentCode}
      />
      <Stack direction="horizontal" gap={1}>
        <Button
          size="sm"
          variant="secondary"
          onClick={() => forceSave(studentCode)}
          disabled={isSaved}
        >
          {isSaved ? "Draft Saved" : isSaving ? "Saving..." : "Save Draft"}
        </Button>
        <Button
          size="sm"
          variant="primary"
          disabled={isSubmitting}
          onClick={submit}
        >
          {isSubmitting ? "Submitting" : "Submit"}
        </Button>
      </Stack>
    </Stack>
  );
}
