import React from "react";
import { Route, Routes } from "react-router-dom";
import Home from './pages/Home';
import About from './pages/About';
import NotFound from './pages/NotFound';
import Body from './components/Body';
import CGPA from './pages/CGPA';
import SGPA from './pages/SGPA';

const App: React.FC = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/body" element={<Body />} />
        <Route path="/cgpa" element={<CGPA />} />
        <Route path="/sgpa" element={<SGPA />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </div>
  );
};

export default App;