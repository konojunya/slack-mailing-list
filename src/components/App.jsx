import React from "react"
import jsCookie from "js-cookie"

// components
import Login from "./Login"
import Main from "./Main"

export default () => {
  const authed = jsCookie.get("slack-mailing-list-authed");
  return authed === "true" ? <Main /> : <Login />;
}
