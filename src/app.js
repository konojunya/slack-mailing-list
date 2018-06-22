import React from "react"
import { hydrate } from "react-dom";

import App from "./components/Index"

hydrate(
  <App />,
  document.getElementById("app")
)