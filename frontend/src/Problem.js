// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useState } from "react";
import { Col, Container, Row, Tab, Tabs } from "react-bootstrap";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Link, Navigate, useParams } from "react-router-dom";

import ProblemEditor from "./ProblemEditor";
import ProblemRoster from "./ProblemRoster";
import ProblemStatement from "./ProblemStatement";
import ProblemSubmissions from "./ProblemSubmissions";
import ProblemTestCases from "./ProblemTestCases";
import ProblemValidation from "./ProblemValidation";

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
          ...ProblemRoster_problem
          ...ProblemSubmissions_problem
          ...ProblemEditor_problem
          ...ProblemValidation_problem
        }
      }
    `,
    { id }
  );

  const [tab, setTab] = useState("problem");

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
              {viewer.is_staff && <ProblemRoster problem={coding_problem} />}
            </Tab>
            <Tab eventKey="validation" title="Validation">
              <ProblemValidation problem={coding_problem} />
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
            onValidate={() => setTab("validation")}
            onSubmit={() => setTab("submissions")}
          />
        </Col>
      </Row>
    </Container>
  );
}
