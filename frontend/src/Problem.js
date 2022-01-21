// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Navigate, useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import { Col, Container, Row } from "react-bootstrap";
import ProblemEditor from "./ProblemEditor";
import ProblemSubmissions from "./ProblemSubmissions";

export default function Problems(): React.Node {
  const { id } = useParams();

  const { coding_problem } = useLazyLoadQuery(
    graphql`
      query ProblemQuery($id: ID!) {
        coding_problem(id: $id) {
          name
          statement
          ...ProblemSubmissions_problem
          ...ProblemEditor_problem
        }
      }
    `,
    { id }
  );

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
          <ProblemSubmissions problem={coding_problem} />
        </Col>
        <Col>
          <ProblemEditor problem={coding_problem} />
        </Col>
      </Row>
    </Container>
  );
}
