import React, { useState } from 'react';
import './Subject.css';

interface SubjectProps {
  subjectName: string;
  maxTheory: number;
  maxPractical?: number; // Optional for subjects without practical
  onMarksChange: (theory: number, practical: number) => void;
}

const Subject: React.FC<SubjectProps> = ({ 
  subjectName, 
  maxTheory = 100, 
  maxPractical = 0, 
  onMarksChange 
}) => {
  const [theoryMarks, setTheoryMarks] = useState(0);
  const [practicalMarks, setPracticalMarks] = useState(0);

  const handleTheoryChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = Number(e.target.value);
    setTheoryMarks(value);
    onMarksChange(value, practicalMarks);
  };

  const handlePracticalChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = Number(e.target.value);
    setPracticalMarks(value);
    onMarksChange(theoryMarks, value);
  };

  return (
    <div className="subject-container">
      <h3>{subjectName}</h3>
      <div className="marks-container">
        <div className="marks-section">
          <label>Theory Marks: {theoryMarks}</label>
          <input
            type="range"
            min="0"
            max={maxTheory}
            value={theoryMarks}
            onChange={handleTheoryChange}
            className="slider"
          />
        </div>
        {maxPractical > 0 && (
          <div className="marks-section">
            <label>Practical Marks: {practicalMarks}</label>
            <input
              type="range"
              min="0"
              max={maxPractical}
              value={practicalMarks}
              onChange={handlePracticalChange}
              className="slider"
            />
          </div>
        )}
      </div>
      <div className="total-marks">
        Total: {theoryMarks + practicalMarks}/{maxTheory + maxPractical}
      </div>
    </div>
  );
};

export default Subject;