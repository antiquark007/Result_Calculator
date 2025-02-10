import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Body.css';

const Body: React.FC = () => {
  const navigate = useNavigate();
  const [branch, setBranch] = useState('');
  const [semester, setSemester] = useState('');
  const [year, setYear] = useState('');

  const handleCGPA = () => {
    navigate('/cgpa', { state: { branch, semester, year } });
  };

  const handleSGPA = () => {
    navigate('/sgpa', { state: { branch, semester, year } });
  };

  return (
    <div className="body">
      <div className="selectors">
        <select value={branch} onChange={(e) => setBranch(e.target.value)}>
          <option value="">Branch</option>
          <option value="CSE">CSE</option>
          <option value="ECE">ECE</option>
          <option value="ME">ME</option>
        </select>
        <select value={semester} onChange={(e) => setSemester(e.target.value)}>
          <option value="">Semesters</option>
          <option value="1">Semester 1</option>
          <option value="2">Semester 2</option>
          <option value="3">Semester 3</option>
        </select>
        <select value={year} onChange={(e) => setYear(e.target.value)}>
          <option value="">Graduating-year</option>
          <option value="2024">2024</option>
          <option value="2025">2025</option>
          <option value="2026">2026</option>
        </select>
      </div>
      <div className="smt-btn">
        <button onClick={handleCGPA}>CGPA</button>
        <button onClick={handleSGPA}>SGPA</button>
      </div>
    </div>
  );
};

export default Body;