// @flow

import * as React from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay";
import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import LoginButton from "./LoginButton";
import { BrowserRouter } from "react-router-dom";

export default function App(): React.Node {
  const { viewer } = useLazyLoadQuery(
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
      <Navbar className="mb-4 border-bottom" expand="lg">
        <Navbar.Brand className="fs-4" href="/">
          CS 170 Online Judge
        </Navbar.Brand>
        {viewer && (
          <>
            <Navbar.Toggle aria-controls="navbar-nav" />
            <Navbar.Collapse
              className="me-auto justify-content-end"
              id="navbar-nav"
            >
              <Nav>
                <Nav.Link href="#home" active>
                  Home
                </Nav.Link>
                <Nav.Link href="#features">Problems</Nav.Link>
                <Nav.Link href="#pricing">Submissions</Nav.Link>
                <Nav.Link href="/logout">Log Out</Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </>
        )}
      </Navbar>
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
            {viewer?.name && <>Welcome, {viewer?.name}!</>}
          </p>
          {!viewer && <LoginButton />}
        </Container>
      </div>
    </Container>
  );
}
