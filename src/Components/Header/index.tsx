import "./Header.scss";
import logo from "@imgs/react.svg"

const Header = () => {
  return (
    <header className="header">
      <img src={logo} alt="logotype" />
      <h1>I'm Header</h1>
    </header>
  );
};

export default Header;
