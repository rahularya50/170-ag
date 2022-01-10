// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay";
import Container from "react-bootstrap/Container";
import LoginButton from "./LoginButton";

export default function App(): React.Node {
  const data = useLazyLoadQuery(
    graphql`
      query AppQuery {
        viewer {
          name
        }
      }
    `,
    {}
  );
  return (
    <Container className="py-4">
      <header className="pb-3 mb-4 border-bottom">
        <a
          href="/"
          className="d-flex align-items-center text-dark text-decoration-none"
        >
          <span className="fs-4">CS 170 Online Judge</span>
        </a>
      </header>
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
            {data.viewer?.name ? (
              <>Hello, {data.viewer?.name}!</>
            ) : (
              "Hello, mysterious stranger!"
            )}
          </p>
          <LoginButton />
        </Container>
      </div>
    </Container>
  );
}
