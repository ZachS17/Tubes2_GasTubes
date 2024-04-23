// import React, { useState } from 'react';
import './styles.css'
import videobg from '../assets/background animation.mp4'

const ResultPage = () => {
    // const [initialpage, setInitialPage] = useState('');
    // const [destinationPage, setDestinationPage] = useState ('');

    // const handleInitialPageChange = (event) => {
    //     setInitialPage(event.target.value)
    // }

    // const handleDestinationPageChange = (event) => {
    //     setDestinationPage(event.target.value)
    // }

  return (
    <div className="ResultPageContainer">
        <video autoPlay muted loop id="videobg">
        <source src={videobg} type="video/mp4" />
      </video>
        <header className='ResultPageHeader'>
            <p className='ResultPageTitle'>
                Result:
            </p>
        </header>
        <div className='ResultContainer'></div>
        <p className='ExecTimeText'>
          Execution time : 
        </p>
    </div>
  );
};

export default ResultPage;
