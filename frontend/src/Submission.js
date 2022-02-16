// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import { useMemo } from "react";
import { Col, Container, Row, Table } from "react-bootstrap";
import { useLazyLoadQuery } from "react-relay/hooks";
import { generatePath, Link, Navigate, useParams } from "react-router-dom";

import SubmissionQuery from "./__generated__/SubmissionQuery.graphql";
import { useRefreshingQuery } from "./useRefreshingQuery";

export default function Submission(): React.Node {
  const { id } = useParams();

  const { coding_submission } = useLazyLoadQuery(
    graphql`
      query SubmissionQuery($id: ID!) {
        coding_submission(id: $id) {
          create_time
          code
          points
          status
          coding_problem {
            id
            name
          }
          author {
            name
          }
          results {
            case_results {
              case_name
              result
              points
            }
          }
        }
      }
    `,
    { id }
  );

  useRefreshingQuery(
    SubmissionQuery,
    useMemo(() => ({ id }), [id])
  );

  if (!coding_submission) {
    return <Navigate to="404" />;
  }

  return (
    <Container>
      <Row>
        <Col>
          <h2>
            {coding_submission.author.name}'s submission at{" "}
            {Intl.DateTimeFormat("en-US", {
              hour: "numeric",
              minute: "numeric",
              day: "numeric",
              month: "numeric",
              year: "numeric",
            }).format(new Date(coding_submission.create_time))}
          </h2>
          <p className="lead">
            {coding_submission.points == null ? (
              <>Submission is {coding_submission.status}.</>
            ) : (
              <>Submission received {coding_submission.points} points.</>
            )}
          </p>
          {coding_submission.results.case_results.length > 0 && (
            <Table hover>
              <thead>
                <tr>
                  <th>Submission</th>
                  <th>Status</th>
                  <th>Score</th>
                </tr>
              </thead>
              <tbody>
                {coding_submission.results.case_results.map(
                  ({ case_name, result, points }) => (
                    <tr key={case_name}>
                      <td>{case_name} </td>
                      <td>{result}</td>
                      <td>{points} points</td>
                    </tr>
                  )
                )}
              </tbody>
            </Table>
          )}
          <Link
            to={generatePath("/problem/:id", {
              id: coding_submission.coding_problem.id,
            })}
          >
            (Back to {coding_submission.coding_problem.name})
          </Link>
        </Col>
        <Col>
          <h3>Submitted Code:</h3>
          <pre>{coding_submission.code}</pre>
        </Col>
      </Row>
    </Container>
  );
}
