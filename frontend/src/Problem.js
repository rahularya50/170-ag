// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import { Button, Col, Container, Modal, Row, Tab, Tabs } from "react-bootstrap";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Link, Navigate, useParams } from "react-router-dom";

import ProblemEditor from "./ProblemEditor";
import ProblemExtensionRosterEditor from "./ProblemExtensionRosterEditor";
import ProblemStatement from "./ProblemStatement";
import ProblemSubmissions from "./ProblemSubmissions";
import ProblemTestCases from "./ProblemTestCases";

export default function Problem(): React.Node {
  const { id } = useParams();

  const { coding_problem, viewer } = useLazyLoadQuery(
    graphql`
      query ProblemQuery($id: ID!) {
        viewer {
          is_staff
          ...ProblemStatement_viewer
          ...ProblemTestCases_viewer
          ...ProblemEditor_viewer
        }
        coding_problem(id: $id) {
          ...ProblemStatement_problem
          ...ProblemTestCases_problem
          ...ProblemExtensionRosterEditor_problem
          ...ProblemSubmissions_problem
          ...ProblemEditor_problem
        }
      }
    `,
    { id }
  );

  const [tab, setTab] = useState("problem");
  const [showRosterEditor, setShowRosterEditor] = useState(false);

  if (!viewer) {
    return <Navigate to="404" />;
  }

  return (
    <Container>
      <Row>
        <Col>
          <Tabs activeKey={tab} onSelect={setTab} className="mb-3">
            <Tab eventKey="problem" title="Problem">
              <ProblemStatement viewer={viewer} problem={coding_problem} />
              <ProblemTestCases viewer={viewer} problem={coding_problem} />
              {viewer.is_staff && (
                <Button
                  className="mb-3"
                  onClick={() => setShowRosterEditor(true)}
                  size="sm"
                >
                  Edit Roster
                </Button>
              )}
            </Tab>
            <Tab eventKey="submissions" title="Submissions">
              <ProblemSubmissions problem={coding_problem} />
            </Tab>
          </Tabs>
          <Link to="/problems/">(Back to all problems)</Link>
        </Col>
        <Col>
          <ProblemEditor
            viewer={viewer}
            problem={coding_problem}
            onSubmit={() => setTab("submissions")}
          />
        </Col>
      </Row>
      <Modal
        show={showRosterEditor}
        onHide={() => setShowRosterEditor(false)}
        size="lg"
      >
        <Modal.Body>
          {showRosterEditor && (
            <React.Suspense fallback={null}>
              <ProblemExtensionRosterEditor problem={coding_problem} />
            </React.Suspense>
          )}
        </Modal.Body>
      </Modal>
    </Container>
  );
}
