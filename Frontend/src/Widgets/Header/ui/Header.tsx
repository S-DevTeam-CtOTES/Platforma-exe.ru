import { App } from "@/Shared";
import './Header.scss';
import { ToggleTheme } from "@/Features";

const Header = () => {
  return (
    <header className="header">
      <div className="container">
        <nav className="header__nav">
          <div className="header__nav-wrapper">
            <div className="header__nav-wrapper-burger">
              <div className="header__nav-wrapper-burger-line"></div>
              <div className="header__nav-wrapper-burger-line"></div>
            </div>
            <div className="header__nav-wrapper-title">{App.Name}</div>
          </div>

          <ToggleTheme/>
        </nav>
      </div>
    </header>
  );
};

export { Header };
