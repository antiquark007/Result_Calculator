import React from "react";
import { Routes, Route } from "react-router-dom";
import Router from "./router/Router";

const App: React.FC = () => {
  return (
    <>
      <Routes>
        {Router.map((route) => (
          <Route key={route.path} path={route.path} element={route.element} />
        ))}
      </Routes>
    </>
  );
};

export default App;