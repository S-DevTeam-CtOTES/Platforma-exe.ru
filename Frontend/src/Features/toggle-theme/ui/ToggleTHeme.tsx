import { useState, useEffect } from "react";

import blueMoon from "@/Shared/assets/icons/blueMoon.svg";
import blueSun from "@/Shared/assets/icons/blueSun.svg";
import "./ToggleTheme.scss";

const ToggleTheme = () => {
  const [toggle, setToggle] = useState<boolean>(
    () => localStorage.getItem("selectedTheme") === "dark"
  );

  useEffect(() => {
    const currentTheme = localStorage.getItem("selectedTheme") || "light";
    document.querySelector("body").setAttribute("data-theme", currentTheme);
  }, []);

  const setDarkMode = () => {
    document.querySelector("body").setAttribute("data-theme", "dark");
    localStorage.setItem("selectedTheme", "dark");
    setToggle(true);
  };

  const setLightMode = () => {
    document.querySelector("body").setAttribute("data-theme", "light");
    localStorage.setItem("selectedTheme", "light");
    setToggle(false);
  };

  const toggleTheme = () => {
    if (toggle) setLightMode();
    else setDarkMode();
  };

  return (
    <label className="theme__toggle">
      <input
        type="checkbox"
        id="darkmode-toggle"
        className="theme__toggle-input"
        checked={toggle}
        onChange={toggleTheme}
      />
      <span className="theme__toggle-slider theme__toggle-round">
        <div className="theme__toggle-slider-wrap ">
          <div className="theme__toggle-slider-wrap-img">
            <img src={toggle ? blueMoon : blueSun} alt="toggle" />
          </div>
        </div>
      </span>
    </label>
  );
};

export default ToggleTheme;
