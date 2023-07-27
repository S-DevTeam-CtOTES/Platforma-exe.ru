import { Routes, Route } from "react-router-dom";
import { Main } from "./ui";
import { Reviews } from "@/Widgets";

export const Routing = () => {
  return (
    <Routes>
      {/* <Route path="/" element={<Main />} /> */}
      <Route path="/" element={<Reviews />} />
    </Routes>
  );
};
