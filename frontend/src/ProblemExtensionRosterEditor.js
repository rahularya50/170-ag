// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import AceEditor from "react-ace";
import { Col, Container, Row } from "react-bootstrap";
import { useFragment, useLazyLoadQuery, useMutation } from "react-relay/hooks";

import type { ProblemExtensionRosterEditor_problem$key } from "./__generated__/ProblemExtensionRosterEditor_problem.graphql";
import type { ProblemExtensionRosterEditorMutation } from "./__generated__/ProblemExtensionRosterEditorMutation.graphql";
import LoadingButton from "./LoadingButton";

type Props = {
  problem: ProblemExtensionRosterEditor_problem$key,
};

export default function ProblemExtensionRosterEditor(props: Props): React.Node {
  const { id } = useFragment(
    graphql`
      fragment ProblemExtensionRosterEditor_problem on CodingProblem {
        id
      }
    `,
    props.problem
  );

  const { coding_problem } = useLazyLoadQuery(
    graphql`
      query ProblemExtensionRosterEditorQuery($id: ID!) {
        coding_problem(id: $id) {
          extension_roster
        }
      }
    `,
    { id },
    { fetchPolicy: "network-only" }
  );

  const [updateRoster, isUpdatingRoster] =
    useMutation<ProblemExtensionRosterEditorMutation>(graphql`
      mutation ProblemExtensionRosterEditorMutation($input: ExtensionsInput!) {
        set_problem_extensions(input: $input) {
          my_deadline
          extension_roster
        }
      }
    `);

  const saveRoster = () => {
    updateRoster({
      variables: {
        input: {
          problem_id: id,
          roster,
        },
      },
      onError: (err) => {
        alert(err);
      },
    });
  };

  const [roster, setRoster] = useState(coding_problem.extension_roster);

  return (
    <Container>
      <Row>
        <Col>
          <AceEditor
            name="case-input"
            width={500}
            onChange={setRoster}
            value={roster}
          />
          <LoadingButton
            onClick={saveRoster}
            size="sm"
            isUpdating={isUpdatingRoster}
          >
            Save
          </LoadingButton>
        </Col>
      </Row>
    </Container>
  );
}
