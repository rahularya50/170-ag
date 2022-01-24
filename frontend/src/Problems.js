// @flow

import * as React from "react";
import { useState } from "react";
import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay/hooks";
import { Form } from "react-bootstrap";
import ProblemTable from "./ProblemTable";

export default function Problems(): React.Node {
  const [includeUnreleased, setIncludeUnreleased] = useState(false);

  const query = useLazyLoadQuery(
    graphql`
      query ProblemsQuery($include_unreleased: Boolean!) {
        viewer {
          is_staff
        }
        ...ProblemTable
      }
    `,
    { include_unreleased: includeUnreleased }
  );

  return (
    <>
      <h3>Problems</h3>
      {query.viewer?.is_staff && (
        <Form.Group className="mb-3" controlId="unreleasedCheckbox">
          <Form.Check
            type="checkbox"
            label="Include unreleased"
            checked={includeUnreleased}
            onChange={(e) => setIncludeUnreleased(e.target.checked)}
          />
        </Form.Group>
      )}
      <React.Suspense fallback={null}>
        <ProblemTable query={query} />
      </React.Suspense>
    </>
  );
}
