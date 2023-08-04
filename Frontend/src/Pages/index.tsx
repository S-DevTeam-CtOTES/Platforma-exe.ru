import { Routes, Route } from "react-router-dom";
import { Error404, Main, Research } from "./ui";

export const Routing = () => {
  return (
    <Routes>
      <Route path="/" element={<Main />} />
      <Route path="/*" element={<Error404 />} />
      <Route path="/Research" element={<Research/>} />
    </Routes>
  );
};
