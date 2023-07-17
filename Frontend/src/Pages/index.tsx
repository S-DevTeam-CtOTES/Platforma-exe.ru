import { Routes, Route } from "react-router-dom";
import { Error404, Main } from "./ui";

export const Routing = () => {
  return (
    <Routes>
      <Route path="/" element={<Main />} />
      <Route path="/*" element={<Error404 />} />
    </Routes>
  );
};
