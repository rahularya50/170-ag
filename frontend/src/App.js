// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay";
import Container from "react-bootstrap/Container";

export default function App(): React.Node {
  const data = useLazyLoadQuery(
    graphql`
      query AppQuery($ids: [ID!]!) {
        user(id: "1") {
          name
          age
        }
        nodes(ids: $ids) {
          id
          ... on User {
            age
          }
        }
      }
    `,
    { ids: ["1", "2"] }
  );
  return (
    <Container className="py-4">
      <header class="pb-3 mb-4 border-bottom">
        <a
          href="/"
          class="d-flex align-items-center text-dark text-decoration-none"
        >
          <span class="fs-4">CS 170 Online Judge</span>
        </a>
      </header>
      <div class="p-5 mb-4 bg-light rounded-3">
        <Container className="py-5" fluid>
          <h1 class="display-5 fw-bold">Welcome to the CS 170 Online Judge!</h1>
          <p class="col-md-8 fs-4">
            This is some placeholder text to be replaced later! Probably the TAs
            will have opinions on what to put here?
          </p>
          <button class="btn btn-primary btn-lg" type="button">
            Log In
          </button>
        </Container>
      </div>
    </Container>
  );
}
