// @flow

import * as React from "react";
import { Button, Spinner } from "react-bootstrap";

type Props = {
  onClick: () => mixed,
  disabled?: ?boolean,
  isUpdating: boolean,
  children: string,
  size?: ?("sm" | "lg"),
  variant?: ?("primary" | "secondary" | "danger"),
};

export default function LoadingButton(props: Props): React.Node {
  return (
    <Button
      onClick={props.onClick}
      disabled={props.isUpdating || (props.disabled ?? false)}
      size={props.size}
      variant={props.variant}
    >
      {props.isUpdating ? (
        <Spinner
          as="span"
          animation="border"
          size="sm"
          role="status"
          aria-hidden="true"
        />
      ) : (
        props.children
      )}
    </Button>
  );
}
