// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import { useLazyLoadQuery } from "react-relay";
import { LinkContainer } from "react-router-bootstrap";
import { generatePath, Navigate, Route, Routes } from "react-router-dom";

import Home from "./Home";
import Problem from "./Problem";
import Problems from "./Problems";
import Submission from "./Submission";
import User from "./User";

export default function App(): React.Node {
  const { viewer } = useLazyLoadQuery(
    graphql`
      query AppQuery {
        viewer {
          id
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
        <Route
          path="/whoami"
          element={
            <Navigate to={generatePath("/user/:id", { id: viewer.id })} />
          }
        />
        <Route path="/user/:id" element={<User />} />
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
                <LinkContainer
                  to={generatePath("/user/:id", { id: viewer.id })}
                >
                  <Nav.Link>Account</Nav.Link>
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
