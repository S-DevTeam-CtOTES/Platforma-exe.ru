import { useState, useEffect } from "react";
import blueMoon from "@/Shared/assets/icons/blueMoon.svg";
import blueSun from "@/Shared/assets/icons/blueSun.svg";
import "./ToggleTheme.scss";

const ToggleTheme = () => {
  const [theme, setTheme] = useState<string>(
    () => localStorage.getItem("selectedTheme") || ""
  );

  useEffect(() => {
    document.documentElement.setAttribute("data-theme", theme);

    const darkModeMediaQuery = window.matchMedia(
      "(prefers-color-scheme: dark)"
    );

    // @Обработчик изменения темы ОС
    const handleDarkModeChange = (event) => {
      const newTheme = event.matches ? "dark" : "light";
      setTheme(newTheme);
      localStorage.setItem("selectedTheme", newTheme);
    };

    darkModeMediaQuery.addListener(handleDarkModeChange);

    return () => {
      darkModeMediaQuery.removeListener(handleDarkModeChange);
    };
  }, [theme]);

  const toggleTheme = () => {
    const newTheme = theme === "dark" ? "light" : "dark";
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
