// import React, { useState } from 'react';
import './styles.css'
import videobg from '../assets/background animation.mp4'

const LoadingPage = () => {
    // const [initialpage, setInitialPage] = useState('');
    // const [destinationPage, setDestinationPage] = useState ('');

    // const handleInitialPageChange = (event) => {
    //     setInitialPage(event.target.value)
    // }

    // const handleDestinationPageChange = (event) => {
    //     setDestinationPage(event.target.value)
    // }

  return (
    <div className="LoadingPageContainer">
        <video autoPlay muted loop id="videobg">
        <source src={videobg} type="video/mp4" />
      </video>
        <header className='LoadingPageHeader'>
            <p className='LoadingPageTitle'>
                Loading...
            </p>
            <p>
            Finding the solution...
            </p>
        </header>
    </div>
  );
};

export default LoadingPage;
