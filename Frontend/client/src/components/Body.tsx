import React from "react";

import './Body.css';

const Body: React.FC = () => {
  return (
    <div className="body">
      <div className="selectors">
        <select>
          <option value="selectors">Branch</option>
          <option value="option1">Option 1</option>
          <option value="option2">Option 2</option>
          <option value="option3">Option 3</option>
        </select>
        <select>
          <option value="selectors">Semesters</option>
          <option value="option1">Option 1</option>
          <option value="option2">Option 2</option>
          <option value="option3">Option 3</option>
        </select>
        <select>
          <option value="selectors">Graduating-year</option>
          <option value="option1">Option 1</option>
          <option value="option2">Option 2</option>
          <option value="option3">Option 3</option>
        </select>
      </div>
      <div className="smt-btn">
        <button>CGPA</button>
        <button>SGPA</button>
      </div>
    </div>
  );
};

export default Body