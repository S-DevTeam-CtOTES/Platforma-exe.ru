import React, { useState, useEffect, createContext, useContext } from "react";

const ThemeContext = createContext({
  theme: "light",
  setTheme: (theme: string) => {},
});

export const useTheme = () => useContext(ThemeContext);

export const ThemeProvider: React.FC = ({ children }) => {
  const [theme, setTheme] = useState(
    () => localStorage.getItem("selectedTheme") || "light"
  );

  useEffect(() => {
    const currentTheme = localStorage.getItem("selectedTheme") || "light";
    document.querySelector("body")?.setAttribute("data-theme", currentTheme);
    setTheme(currentTheme);
  }, []);

  const handleThemeChange = (selectedTheme: string) => {
    setTheme(selectedTheme);
    localStorage.setItem("selectedTheme", selectedTheme);
    document.querySelector("body")?.setAttribute("data-theme", selectedTheme);
  };

  return (
    <ThemeContext.Provider value={{ theme, setTheme: handleThemeChange }}>
      {children}
    </ThemeContext.Provider>
  );
};
