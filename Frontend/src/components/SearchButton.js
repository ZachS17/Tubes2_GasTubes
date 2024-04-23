import React from 'react';
import "./styles.css"

const SearchButton = () => {
  // Function to handle search button click
  const handleSearch = () => {
    // Implement search functionality here
    console.log('Search button clicked');
  };

  return (
    <button onClick={handleSearch}>Search</button>
  );
};

export default SearchButton;
