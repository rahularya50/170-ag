// @flow

import graphql from "babel-plugin-relay/macro";
import * as React from "react";
import AceEditor from "react-ace";
import { Col, Container, Row } from "react-bootstrap";
import { useFragment, useLazyLoadQuery } from "react-relay/hooks";

import type { ProblemTestDataEditor_testCase$key } from "./__generated__/ProblemTestDataEditor_testCase.graphql";

type Props = {
  testCase: ProblemTestDataEditor_testCase$key,
  input: ?string,
  output: ?string,
  onInputChange: (string) => mixed,
  onOutputChange: (string) => mixed,
};

export default function ProblemTestDataEditor(props: Props): React.Node {
  const { id } = useFragment(
    graphql`
      fragment ProblemTestDataEditor_testCase on CodingTestCase {
        id
      }
    `,
    props.testCase
  );

  const { test_case } = useLazyLoadQuery(
    graphql`
      query ProblemTestDataEditorQuery($id: ID!) {
        test_case(id: $id) {
          data {
            input
            output
          }
        }
      }
    `,
    { id },
    { fetchPolicy: "network-only" }
  );

  return (
    <Container>
      <Row>
        <Col md={6}>
          <AceEditor
            name="case-input"
            width={300}
            onChange={props.onInputChange}
            value={props.input ?? test_case.data.input}
          />
        </Col>
        <Col md={6}>
          <AceEditor
            name="case-output"
            width={300}
            onChange={props.onOutputChange}
            value={props.output ?? test_case.data.output}
          />
        </Col>
      </Row>
    </Container>
  );
}
