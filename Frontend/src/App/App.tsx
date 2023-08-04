import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { ThemeProvider } from "./providers/Theme-provider";

import { Error404, Main, Research } from "@/Pages/ui";
import { Routing } from "@/Pages";

const App = () => {
  return (
       <Router>
      <ThemeProvider>
        <Routing />
      </ThemeProvider>
    </Router>
  );
};

export default App;
