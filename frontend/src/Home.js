// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import Container from "react-bootstrap/Container";
import LoginButton from "./LoginButton";
import { useLazyLoadQuery } from "react-relay/hooks";

export default function Home(): React.Node {
  const { viewer } = useLazyLoadQuery(
    graphql`
      query HomeQuery {
        viewer {
          name
        }
      }
    `,
    {}
  );
  return (
    <div className="p-5 mb-4 bg-light rounded-3">
      <Container className="py-5" fluid>
        <h1 className="display-5 fw-bold">
          Welcome to the CS 170 Online Judge!
        </h1>
        <p className="col-md-8 fs-4">
          This is some placeholder text to be replaced later! Probably the TAs
          will have opinions on what to put here?
        </p>
        <p className="col-md-8 fs-4">
          {viewer?.name != null && <>Welcome, {viewer.name}!</>}
        </p>
        {!viewer && <LoginButton />}
      </Container>
    </div>
  );
}
