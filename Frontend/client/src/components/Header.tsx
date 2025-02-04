import React from "react";
import "./Header.css";

const Header: React.FC = () => {

  return (
    <header>
      <nav className="navbar">
        <h1 className="navbar-logo">score calculator</h1>
        <div className="navbar-user">
          <button className="sign-up">Sign-up</button>
        </div>
      </nav>
    </header>
  );
};

export default Header;