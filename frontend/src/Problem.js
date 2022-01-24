// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Navigate, useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import { Col, Container, Row } from "react-bootstrap";
import ProblemEditor from "./ProblemEditor";
import ProblemSubmissions from "./ProblemSubmissions";
import ProblemStatement from "./ProblemStatement";
import ProblemTestCases from "./ProblemTestCases";

export default function Problem(): React.Node {
  const { id } = useParams();

  const { coding_problem, viewer } = useLazyLoadQuery(
    graphql`
      query ProblemQuery($id: ID!) {
        viewer {
          ...ProblemStatement_viewer
          ...ProblemTestCases_viewer
        }
        coding_problem(id: $id) {
          ...ProblemStatement_problem
          ...ProblemTestCases_problem
          ...ProblemSubmissions_problem
          ...ProblemEditor_problem
        }
      }
    `,
    { id }
  );

  if (!coding_problem || !viewer) {
    return <Navigate to="404" />;
  }

  return (
    <Container>
      <Row>
        <Col>
          <ProblemStatement viewer={viewer} problem={coding_problem} />
          <ProblemTestCases viewer={viewer} problem={coding_problem} />
          <Link to="/problems/">(Back to all problems)</Link>
          <ProblemSubmissions problem={coding_problem} />
        </Col>
        <Col>
          <ProblemEditor problem={coding_problem} />
        </Col>
      </Row>
    </Container>
  );
}
