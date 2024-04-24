import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import logo from './assets/wikipedia logo art.png';
import videobg from './assets/background animation.mp4'
// import SearchBar from './components/SearchBar.js';
// import SearchButton from './components/SearchButton.js';
// import LoadingPage from './components/LoadingPage.js';
// import ResultPage from './components/ResultPage.js';
import MainPage from './components/MainPage';

function App() {
  return (
    <Router>
      <div className="app-container">
        <Routes>
          <Route path="/*" element={<Home/>} />
        </Routes>
      </div>
    </Router>
  );
}

const Home = () => {
  return (
    <div className="App">
      <video autoPlay muted loop id="videobg">
        <source src={videobg} type="video/mp4" />
      </video>
      <header className="App-header">
        <img src={logo} style={{width:'350px', height:'200px',margin:'10px'}} className="App-logo" alt="logo" />
        <p className="Title">
          Wikirace
        </p>
      </header>
      <MainPage/>
    </div>
  );
}

export default App;
