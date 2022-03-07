// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import { Button, Modal } from "react-bootstrap";
import { useFragment } from "react-relay/hooks";

import type { ProblemRoster_problem$key } from "./__generated__/ProblemRoster_problem.graphql";
import type { ProblemRosterScoresQuery } from "./__generated__/ProblemRosterScoresQuery.graphql";
import downloadText from "./downloadText";
import LoadingButton from "./LoadingButton";
import ProblemExtensionRosterEditor from "./ProblemExtensionRosterEditor";
import { useQueryFetcher } from "./useQueryFetcher";

type Props = {
  problem: ProblemRoster_problem$key,
};

export default function ProblemRoster(props: Props): React.Node {
  const problem = useFragment(
    graphql`
      fragment ProblemRoster_problem on CodingProblem {
        id
        ...ProblemExtensionRosterEditor_problem
      }
    `,
    props.problem
  );

  const [showRosterEditor, setShowRosterEditor] = useState(false);

  const [downloadScores, isDownloadingScores] =
    useQueryFetcher<ProblemRosterScoresQuery>(
      graphql`
        query ProblemRosterScoresQuery($id: ID!) {
          coding_problem(id: $id) {
            score_roster
          }
        }
      `,
      { id: problem.id },
      (resp) => {
        downloadText(resp.coding_problem.score_roster, "scores.csv");
      },
      alert
    );

  return (
    <div className="mb-3">
      <Button onClick={() => setShowRosterEditor(true)} size="sm">
        Edit Roster
      </Button>{" "}
      <LoadingButton
        isUpdating={isDownloadingScores}
        onClick={downloadScores}
        size="sm"
      >
        Export Scores
      </LoadingButton>
      <Modal
        show={showRosterEditor}
        onHide={() => setShowRosterEditor(false)}
        size="lg"
      >
        <Modal.Body>
          {showRosterEditor && (
            <React.Suspense fallback={null}>
              <ProblemExtensionRosterEditor problem={problem} />
            </React.Suspense>
          )}
        </Modal.Body>
      </Modal>
    </div>
  );
}
