import React, { useState } from 'react';
import './styles.css'

const SearchBar = () => {
    const [initialpage, setInitialPage] = useState('');
    const [destinationPage, setDestinationPage] = useState ('');

    const handleInitialPageChange = (event) => {
        setInitialPage(event.target.value)
    }

    const handleDestinationPageChange = (event) => {
        setDestinationPage(event.target.value)
    }

  return (
    <div className="SearchBarContainer">
      <div className='SearchBarLabel'>
        <label htmlFor="searchInput" className='InputLabel'>Initial Page:   </label>
        <input
          type="text"
          id="searchInput"
          placeholder="Enter your search query"
          value = {initialpage}
          onChange={handleInitialPageChange}
        />
      </div>
      <div className='SearchBarLabel'>
        <label htmlFor="searchInput" className='InputLabel'>Destination Page:   </label>
        <input
          type="text"
          id="searchInput"
          placeholder="Enter your search query"
          value={destinationPage}
          onChange={handleDestinationPageChange}
        />
      </div>
    </div>
  );
};

export default SearchBar;
