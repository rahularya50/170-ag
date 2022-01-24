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
import TestCase from "./TestCase";

export default function Problem(): React.Node {
  const { id } = useParams();

  const { coding_problem, viewer } = useLazyLoadQuery(
    graphql`
      query ProblemQuery($id: ID!) {
        viewer {
          ...ProblemStatement_viewer
          ...TestCase_viewer
        }
        coding_problem(id: $id) {
          test_cases {
            id
            ...TestCase_testCase
          }
          ...ProblemStatement_problem
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

  <TestCase
    key={coding_problem.test_cases[0].id}
    viewer={viewer}
    testCase={coding_problem.test_cases[0]}
  />;

  return (
    <Container>
      <Row>
        <Col>
          <ProblemStatement viewer={viewer} problem={coding_problem} />
          <h3>Test Cases</h3>
          {coding_problem.test_cases.map((testCase) => (
            <TestCase key={testCase.id} viewer={viewer} testCase={testCase} />
          ))}
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
