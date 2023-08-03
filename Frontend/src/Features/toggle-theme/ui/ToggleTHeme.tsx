import { useState, useEffect } from "react";

import blueMoon from "@/Shared/assets/icons/blueMoon.svg";
import blueSun from "@/Shared/assets/icons/blueSun.svg";
import "./ToggleTheme.scss";

const ToggleTheme = () => {
  const [toggle, setToggle] = useState<boolean>(false);
  const [selectedTheme, setSelectedTheme] = useState(
    () => localStorage.getItem("selectedTheme") || "light"
  );

  useEffect(() => {
    const currentTheme = localStorage.getItem("selectedTheme") || "light";
    document.querySelector("body").setAttribute("data-theme", currentTheme);
    setSelectedTheme(currentTheme);
  }, []);

  const setDarkMode = () => {
    document.querySelector("body").setAttribute("data-theme", "dark");
    setSelectedTheme("dark");
    localStorage.setItem("selectedTheme", "dark");
  };

  const setLightMode = () => {
    document.querySelector("body").setAttribute("data-theme", "light");
    setSelectedTheme("light");
    localStorage.setItem("selectedTheme", "light");
  };

  const toggleTheme = (e) => {
    if (e.target.checked) setDarkMode();
    else setLightMode();
  };

  return (
    <label className="theme__toggle">
      <input
        type="checkbox"
        id="darkmode-toggle"
        className="theme__toggle-input"
        onClick={toggleTheme}
      />
      <span className="theme__toggle-slider theme__toggle-round">
        <div className="theme__toggle-slider-wrap ">
          <div className="theme__toggle-slider-wrap-img">
            <img
              src={selectedTheme === "dark" ? blueMoon : blueSun}
              alt="toggle"
            />
          </div>
        </div>
      </span>
    </label>
  );
};

export default ToggleTheme;
