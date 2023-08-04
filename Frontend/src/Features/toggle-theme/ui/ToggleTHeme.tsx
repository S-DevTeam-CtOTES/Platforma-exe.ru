import { useState, useEffect } from "react";
import blueMoon from "@/Shared/assets/icons/blueMoon.svg";
import blueSun from "@/Shared/assets/icons/blueSun.svg";
import "./ToggleTheme.scss";

const ToggleTheme = () => {
  const [theme, setTheme] = useState<string>("");

  useEffect(() => {
    document.querySelector("body").setAttribute("data-theme", theme);

    setTheme(() => localStorage.getItem("selectedTheme") || "");
    const darkModeMediaQuery = window.matchMedia(
      "(prefers-color-scheme: dark)"
    );

    // @Обработчик изменения темы ОС
    const handleDarkModeChange = (event) => {
      const newTheme = event.matches ? "dark" : "light";
      setTheme(newTheme);
      localStorage.setItem("selectedTheme", newTheme);
    };

    darkModeMediaQuery.addEventListener("change", handleDarkModeChange);

    return () => {
      darkModeMediaQuery.addEventListener("change", handleDarkModeChange);
    };
  }, [theme]);

  const toggleTheme = () => {
    const newTheme = theme === "light" ? "dark" : "light";
    setTheme(newTheme);
    localStorage.setItem("selectedTheme", newTheme);
  };

  return (
    <label className="theme__toggle">
      <input
        type="checkbox"
        id="darkmode-toggle"
        className="theme__toggle-input"
        checked={theme === "dark"}
        onChange={toggleTheme}
      />
      <span className="theme__toggle-slider theme__toggle-round">
        <div className="theme__toggle-slider-wrap ">
          <div className="theme__toggle-slider-wrap-img">
            <img src={theme === "dark" ? blueMoon : blueSun} alt="toggle" />
          </div>
        </div>
      </span>
    </label>
  );
};

export default ToggleTheme;
