// @flow
import type { ProblemEditorSaveDraftMutation } from "./__generated__/ProblemEditorSaveDraftMutation.graphql";
import type { ProblemEditor_problem$key } from "./__generated__/ProblemEditor_problem.graphql";
import type { ProblemEditorSubmitMutation } from "./__generated__/ProblemEditorSubmitMutation.graphql";
import type { ProblemEditorUpdateSkeletonMutation } from "./__generated__/ProblemEditorUpdateSkeletonMutation.graphql";
import type { ProblemEditor_viewer$key } from "./__generated__/ProblemEditor_viewer.graphql";

import * as React from "react";
import { useCallback, useEffect, useState } from "react";
import graphql from "babel-plugin-relay/macro";
import { useFragment, useMutation } from "react-relay/hooks";
import AceEditor from "react-ace";
import { Button, Stack } from "react-bootstrap";
import { useDebouncedCallback } from "./useDebounced";
import LoadingButton from "./LoadingButton";

import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/mode-java";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/theme-github";

type Props = {
  viewer: ProblemEditor_viewer$key,
  problem: ProblemEditor_problem$key,
  onSubmit: () => mixed,
};

export default function ProblemEditor(props: Props): React.Node {
  const { is_staff: isStaff } = useFragment(
    graphql`
      fragment ProblemEditor_viewer on User {
        is_staff
      }
    `,
    props.viewer
  );

  const {
    id,
    skeleton,
    my_draft: myDraft,
  } = useFragment(
    graphql`
      fragment ProblemEditor_problem on CodingProblem {
        id
        skeleton
        my_draft {
          code
        }
      }
    `,
    props.problem
  );

  const [saveDraft, isSaving] = useMutation<ProblemEditorSaveDraftMutation>(
    graphql`
      mutation ProblemEditorSaveDraftMutation($input: CodingDraftInput!) {
        save_draft(input: $input) {
          code
        }
      }
    `
  );

  const [updateSkeleton, isUpdatingSkeleton] =
    useMutation<ProblemEditorUpdateSkeletonMutation>(
      graphql`
        mutation ProblemEditorUpdateSkeletonMutation(
          $input: UpdateCodingProblemInput!
        ) {
          update_problem(input: $input) {
            skeleton
          }
        }
      `
    );

  const handleUpdateSkeleton = () => {
    updateSkeleton({
      variables: {
        input: { id, problem: { skeleton: studentCode } },
      },
      onError: (e) => {
        alert(e);
      },
    });
  };

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
        props.onSubmit();
      },
      onError: () => {
        alert("submission failed!");
      },
    });
  };

  const [autoSaveEnabled, setAutoSaveEnabled] = useState(true);
  const [studentCode, setStudentCode] = useState(myDraft?.code ?? skeleton);
  const [savedStudentCode, setSavedStudentCode] = useState("");

  const isSaved = !isSaving && studentCode === savedStudentCode;

  const immediatelyForceSave = useCallback(
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
  );

  const forceSave = useDebouncedCallback(immediatelyForceSave, 1500);

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
        {isStaff && (
          <LoadingButton
            isUpdating={isUpdatingSkeleton}
            onClick={handleUpdateSkeleton}
            variant="danger"
            size="sm"
          >
            Set as skeleton
          </LoadingButton>
        )}
        <Button
          onClick={() => setStudentCode(skeleton)}
          variant="danger"
          size="sm"
        >
          Reset to skeleton
        </Button>
        <LoadingButton
          isUpdating={isSaving}
          onClick={() => immediatelyForceSave(studentCode)}
          disabled={isSaved}
          variant="secondary"
          size="sm"
        >
          {isSaved ? "Draft Saved" : "Save Draft"}
        </LoadingButton>
        <LoadingButton isUpdating={isSubmitting} onClick={submit} size="sm">
          Submit
        </LoadingButton>
      </Stack>
    </Stack>
  );
}
