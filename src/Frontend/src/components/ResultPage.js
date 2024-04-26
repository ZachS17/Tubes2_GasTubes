// import React, { useState } from 'react';
import './styles.css'

const ResultPage = ({result}) => {
    // const [initialpage, setInitialPage] = useState('');
    // const [destinationPage, setDestinationPage] = useState ('');

    // const handleInitialPageChange = (event) => {
    //     setInitialPage(event.target.value)
    // }

    // const handleDestinationPageChange = (event) => {
    //     setDestinationPage(event.target.value)
    // }

  return (
      <div className='ResultContainer'>
        <p className='ResultPageTitle'>Result:</p>
        <div className='ResultBox'>
            <ol>
                {result.path.map((pageURL, index) => (
                <ul>
                    <a href={pageURL} target="_blank" rel="noopener noreferrer">{index+1}. {pageURL}</a>
                </ul>
                ))}
            </ol>
            <p>Articles Visited: {result.numArticlesVisited}</p>
            <p>Articles Checked: {result.numArticlesChecked}</p>
            <p>Execution Time: {result.executionTime} ms</p>
        </div>
  </div>
  );
};

export default ResultPage;
