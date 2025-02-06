import React from 'react';
import { useLocation } from 'react-router-dom';
import Header from '../components/Header/Header';
import Subject from '../components/result/Subject';

interface LocationState {
  branch: string;
  semester: string;
  year: string;
}

const CGPA: React.FC = () => {
  const location = useLocation();
  const state = location.state as LocationState;

  return (
    <div>
      <Header />
      <Subject
        subjectName="Mathematics"
        maxTheory={70}
        maxPractical={30}
        onMarksChange={(theory, practical) => {
          console.log(`Theory: ${theory}, Practical: ${practical}`);
        }}
      />
      <h2>CGPA Calculator</h2>
      <p>Branch: {state.branch}</p>
      <p>Semester: {state.semester}</p>
      <p>Year: {state.year}</p>
    </div>
  );
};

export default CGPA;