// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay";
import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Home from "./Home";
import { Routes, Route, Navigate } from "react-router-dom";
import Problems from "./Problems";
import { LinkContainer } from "react-router-bootstrap";
import Problem from "./Problem";
import Submission from "./Submission";

export default function App(): React.Node {
  const { viewer } = useLazyLoadQuery(
    graphql`
      query AppQuery {
        viewer {
          __typename
        }
        coding_problem(id: "5") {
          __typename
        }
      }
    `,
    {}
  );

  let children: React.Node;
  if (viewer == null) {
    // not logged in
    children = (
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    );
  } else {
    // logged in
    children = (
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/problems" element={<Problems />} />
        <Route path="/problem/:id" element={<Problem />} />
        <Route path="/submission/:id" element={<Submission />} />
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    );
  }

  return (
    <Container className="py-4">
      <Navbar className="mb-4 border-bottom" expand="lg">
        <LinkContainer to="/">
          <Navbar.Brand className="fs-4">CS 170 Online Judge</Navbar.Brand>
        </LinkContainer>
        {viewer && (
          <>
            <Navbar.Toggle aria-controls="navbar-nav" />
            <Navbar.Collapse
              className="me-auto justify-content-end"
              id="navbar-nav"
            >
              <Nav>
                <LinkContainer to="/">
                  <Nav.Link>Home</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/problems/">
                  <Nav.Link>Problems</Nav.Link>
                </LinkContainer>
                <Nav.Link href="/logout">Log Out</Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </>
        )}
      </Navbar>
      <React.Suspense fallback="">{children}</React.Suspense>
    </Container>
  );
}
