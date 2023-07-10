// import React from "react";
import { withProviders } from "./providers/index";

import { Routing } from "@/Pages";

const App = () => {
  return (
    <Routing/>
  );
};

export  default withProviders(App);
