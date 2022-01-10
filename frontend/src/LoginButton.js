// @flow

import * as React from "react";
import { useEffect, useRef } from "react";
import { client_id } from "./conf";

export default function LoginButton(): React.Node {
  const buttonRef = useRef();

  const handleCredentialResponse = (resp) => {
    console.log(resp);
  };

  useEffect(() => {
    const domain =
      window.location.hostname === "localhost"
        ? "http://localhost:8080"
        : `https://${window.location.hostname}`;
    window.google.accounts.id.initialize({
      client_id,
      //   callback: handleCredentialResponse,
      auto_select: true,
      login_uri: new URL("/login", domain).toString(),
      ux_mode: "redirect",
    });
    window.google.accounts.id.renderButton(buttonRef.current, {
      theme: "outline",
      size: "large",
    });
    window.google.accounts.id.prompt();
  });

  return <div ref={buttonRef}></div>;
}
