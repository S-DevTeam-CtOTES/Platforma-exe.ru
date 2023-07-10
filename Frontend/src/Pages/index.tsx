import { Routes, Route } from "react-router-dom";
import { Main } from "./ui";

export const Routing = () => {
  return (
    <Routes>
      <Route path="/" element={<Main />} />
    </Routes>
  );
};
