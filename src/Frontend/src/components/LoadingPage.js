// import React, { useState } from 'react';
import './styles.css'
import loadingAnimation from '../assets/loading animation.gif'

const LoadingPage = () => {
  return (
    <div className='LoadingPageContainer'>
      <img src={loadingAnimation} alt="loading" className="loadingAnimation"/>
      <p className='LoadingMessage'>Sedang mencari jawaban</p>
    </div>
  );
};

export default LoadingPage;
