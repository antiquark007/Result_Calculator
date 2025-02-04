import React from 'react';
import { useLocation } from 'react-router-dom';
import Header from '../components/Header';

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
      <Header/>
      <h2>CGPA Calculator</h2>
      <p>Branch: {state.branch}</p>
      <p>Semester: {state.semester}</p>
      <p>Year: {state.year}</p>
    </div>
  );
};

export default CGPA;