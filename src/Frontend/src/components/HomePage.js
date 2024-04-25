import React from 'react'
import './styles.css'
import { useState } from 'react';
import logo from '../assets/wikipedia logo art.png'
import videobg from '../assets/background animation.mp4'
import LoadingPage from './LoadingPage';
import ResultPage from './ResultPage';

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
            {isLoading ? (
                <LoadingPage/>
                ) : result ? (
                <ResultPage result={result} />
                    ):
            <div className='SubApp'>
                <video autoPlay muted loop id="videobg">
                <source src={videobg} type="video/mp4" />
                </video>
                <header className="App-header">
                <img src={logo} style={{width:'350px', height:'200px',margin:'10px'}} className="App-logo" alt="logo" />
                <p className="Title">
                    Wikirace
                </p>
                </header>
                <div className="Container">
                    <div className='SearchBarContainer'>
                        <div className='SearchBarLabel'>
                            <input
                                type="text"
                                id="searchInput"
                                onChange={(e) => setInitialPage(e.target.value)}
                                placeholder="Initial Page"
                                value = {initialPage}
                            />
                        </div>
                        <div className='SearchBarLabel'>
                            <input
                                type="text"
                                id="searchInput"
                                placeholder="Destination Page"
                                value={destinationPage}
                                onChange={(e) => setDestinationPage(e.target.value)}
                            />
                        </div>
                    </div>
                    <div className='AlgorithmButton'>
                        <button onClick = {handleBFSClick}>BFS</button>
                        <button onClick = {handleIDSClick}>IDS</button>
                    </div>
                    <div className="SearchButton">
                        <button onClick={handleSearchClick}>Search</button>
                    </div>
                </div>
            </div>
        }
        </div>
    );
}

export default HomePage