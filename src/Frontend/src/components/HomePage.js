import React from 'react'
import './styles.css'
import { useState } from 'react';
import logo from '../assets/wikipedia logo art.png'
import videobg from '../assets/background animation.mp4'
import LoadingPage from './LoadingPage';
import ResultPage from './ResultPage';
import searchBar from '../assets/search bar.png'

const HomePage = () => {
    const [initialPage, setInitialPage] = useState('');
    const [destinationPage, setDestinationPage] = useState ('');
    const [result, setResult] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    const [algorithm, setAlgorithm] = useState('');

    const handleBFSClick = () => {
        setAlgorithm("BFS");
    }

    const handleIDSClick = () => {
        setAlgorithm("IDS");
    }

    const handleSearchClick = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        try {
            // ubah jadi underscore untuk yang > 1 kata
            const spacedInitialPage = initialPage.replace(/ /g, '_');
            const spacedDestinationPage = destinationPage.replace(/ /g, '_');
            // bentuk url untuk algoritma
            const fullInitialPageURL = `https://en.wikipedia.org/wiki/${spacedInitialPage}`;
            const fullDestinationPageURL = `https://en.wikipedia.org/wiki/${spacedDestinationPage}`;
            // minta respon
            const response = await fetch(`http://localhost:8080/wikirace?&algorithm=${encodeURIComponent(algorithm)}&initial=${encodeURIComponent(fullInitialPageURL)}&destination=${encodeURIComponent(fullDestinationPageURL)}`);
            const data = await response.json();

            // set state
            setResult(data);
            setIsLoading(false);
        } catch (error) {
            console.error('Error fetching data:', error);
            setIsLoading(false);
        }
    }

    return (
        <div className="App">
                            <video autoPlay muted loop id="videobg">
                <source src={videobg} type="video/mp4" />
            </video>
            {isLoading ? (
                <LoadingPage/>
                ) : result ? (
                <ResultPage result={result} />
                    ):
            <div className='SubApp'>
                <header className="App-header">
                <img src={logo} style={{width:'350px', height:'200px',margin:'10px'}} className="App-logo" alt="logo" />
                <p className="Title">
                    Wikirace
                </p>
                </header>
                <div className="Container">
                    <div className="InputContainer">
                        <div className="InputLabel" id='InitialPage'>
                            <input
                                type="text"
                                id="SearchInput"
                                onChange={(e) => setInitialPage(e.target.value)}
                                placeholder="Initial Page"
                                value = {initialPage}
                                
                            />
                        </div>
                        <div className='InputLabel' id='DestinationPage'>
                            <input
                                type="text"
                                id="SearchInput"
                                placeholder="Destination Page"
                                value={destinationPage}
                                onChange={(e) => setDestinationPage(e.target.value)}
                            />
                        </div>
                    </div>
                    <div className='AlgorithmButtonContainer'>
                        <button onClick = {handleBFSClick} className='AlgorithmButton' id='BFSButton'>BFS</button>
                        <button onClick = {handleIDSClick} className='AlgorithmButton' id='IDSButton'>IDS</button>
                    </div>
                    <div className="SearchContainer">
                        <button onClick={handleSearchClick} id='SearchBarButton'>
                            Search
                            <img src={searchBar} alt='SearchBar' id='SearchButton'></img>
                        </button>
                    </div>
                </div>
            </div>
        }
        </div>
    );
}

export default HomePage