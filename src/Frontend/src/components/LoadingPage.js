// import React, { useState } from 'react';
import './styles.css'
import loadingAnimation from '../assets/loading animation.gif'

const LoadingPage = () => {
  return (
    <div>
      <img src={loadingAnimation} alt="loading" className="loadingAnimation"/>
      <p>Sedang mencari jawaban, sabar yaa...</p>
    </div>
  );
};

export default LoadingPage;
