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
      <div>
      <ol>
          {result.path.map((pageURL, index) => (
          <li key={index}>
              <a href={pageURL} target="_blank" rel="noopener noreferrer">{pageURL}</a>
          </li>
          ))}
      </ol>
      <p>Articles Visited: {result.numArticlesVisited}</p>
      <p>Articles Checked: {result.numArticlesChecked}</p>
      <p>Execution Time: {result.executionTime} ms</p>
  </div>
  );
};

export default ResultPage;
