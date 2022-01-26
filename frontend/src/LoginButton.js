// @flow

import * as React from "react";
import { useEffect, useRef } from "react";

import { client_id } from "./conf";

export default function LoginButton(): React.Node {
  const buttonRef = useRef();

  useEffect(() => {
    const domain =
      window.location.hostname === "localhost"
        ? "http://localhost:8080"
        : `https://${window.location.hostname}`;
    window.google.accounts.id.initialize({
      client_id,
      auto_select: true,
      login_uri: new URL("/login", domain).toString(),
      ux_mode: "redirect",
    });
    window.google.accounts.id.renderButton(buttonRef.current, {
      theme: "outline",
      size: "large",
    });
  }, []);

  return <div ref={buttonRef}></div>;
}
