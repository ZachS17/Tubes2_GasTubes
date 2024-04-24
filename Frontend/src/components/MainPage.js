import React from 'react'
import './styles.css'
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import loadingAnimation from '../assets/loading animation.gif'

const MainPage = () => {
    const [initialPage, setInitialPage] = useState('');
    const [destinationPage, setDestinationPage] = useState ('');
    const [result, setResult] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    // const [loadingMessage, setLoadingMessage] = useState('');
    // const [showMessage, setShowMessage] = useState(true);

    // const handleInitialPageChange = (event) => {
    //     setInitialPage(event.target.value)
    // }

    // const handleDestinationPageChange = (event) => {
    //     setDestinationPage(event.target.value)
    // }

    const navigate = useNavigate();

    const handleBack = () => {
        navigate('/');
    }

    const handleBFSClick = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        try {
            // ubah jadi underscore untuk yang > 1 kata
            const spacedInitialPage = initialPage.replace(/ /g, '_');
            const spacedDestinationPage = destinationPage.replace(/ /g, '_');
            // bentuk url untuk algoritma
            const fullInitialPageURL = `https://en.wikipedia.org/wiki/${spacedInitialPage}`;
            const fullDestinationPageURL = `https://en.wikipedia.org/wiki/${spacedDestinationPage}`;
    
            const response = await fetch(`http://localhost:8080/shortestpath?algorithm=bfs&initial=${encodeURIComponent(fullInitialPageURL)}&destination=${encodeURIComponent(fullDestinationPageURL)}`);
            const data = await response.json();

            // const nodes = data.path.map((url, index) => {
            //     const pageName = url.split('/').pop(); // ambil nama aja
            //     return {
            //         id: index,
            //         label: pageName,
            //         shape: 'star',
            //         color: index === 0 ? 'red' : index === data.path.length - 1 ? 'rgba(133, 225, 4, 0.946)' : undefined // Set color for start and end points
            //     };
            // });
            // const edges = [];
            // for (let i = 0; i < nodes.length - 1; i++) {
            //     edges.push({ id: `edge${i}`, from: i, to: i + 1, arrows: "to" });
            // }
            // const graph = { nodes, edges };
            // setGraphData(graph);

            setResult(data);
            setIsLoading(false);
        } catch (error) {
            console.error('Error fetching data:', error);
            setIsLoading(false);
        }
    }

    return (
        <div className="Container">
            <div className='SearchBarContainer'>
                <div className='SearchBarLabel'>
                    <label htmlFor="searchInput" className='InputLabel'>Initial Page:   </label>
                    <input
                        type="text"
                        id="searchInput"
                        onChange={(e) => setInitialPage(e.target.value)}
                        placeholder="Enter your search query"
                        value = {initialPage}
                    />
                </div>
                <div className='SearchBarLabel'>
                    <label htmlFor="searchInput" className='InputLabel'>Destination Page:   </label>
                    <input
                        type="text"
                        id="searchInput"
                        placeholder="Enter your search query"
                        value={destinationPage}
                        onChange={(e) => setDestinationPage(e.target.value)}
                    />
                </div>
            </div>
            <div className="SearchButton">
                <button onClick={handleBFSClick}>Search</button>
            </div>
            <div className='ResultContainer'>
                {isLoading ? (
                        <div>
                            <img src={loadingAnimation} alt="loading" className="loadingAnimation"/>
                            <p>Sedang mencari jawaban, sabar yaa...</p>
                        </div>
                    ) : result ? (
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
                    ):null}
            </div>
            <div className="BackButton">
                <button onClick={handleBack}>Back</button>
            </div>
        </div>
    );
}

export default MainPage