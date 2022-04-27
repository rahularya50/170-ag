// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { Col, Container, Row } from "react-bootstrap";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Navigate, useParams } from "react-router-dom";

export default function Problem(): React.Node {
  const { id } = useParams();

  const { user } = useLazyLoadQuery(
    graphql`
      query UserQuery($id: ID!) {
        user(id: $id) {
          email
        }
      }
    `,
    { id }
  );

  if (!user) {
    return <Navigate to="404" />;
  }

  return (
    <Container>
      <Row>
        <Col>
          Your email is <a href={`mailto:${user.email}`}>{user.email}</a>.
        </Col>
      </Row>
    </Container>
  );
}
