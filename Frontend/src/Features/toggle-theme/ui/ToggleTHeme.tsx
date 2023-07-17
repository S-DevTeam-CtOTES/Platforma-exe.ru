import { useState } from "react";

import blueMoon from "@/Shared/assets/icons/blueMoon.svg";
import blueSun from "@/Shared/assets/icons/blueSun.svg";
import "./ToggleTheme.scss";

const ToggleTheme = () => {
  const [toggle, setToggle] = useState<boolean>(false);

  const store = () => {
    setToggle(!toggle);
  };

  return (
    <label className="theme__toggle">
      <input type="checkbox" className="theme__toggle-input" onClick={store} />
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

export { ToggleTheme };
