import React from 'react';
import './App.css';
import {BrowserRouter, Route, Routes} from 'react-router-dom';

import Navbar from './components/Navbar';
import Login from './components/Login';
import Home from './components/Home';

function App() {
  return (
  <div className="App">
    <BrowserRouter>
        <Navbar/>
        <Routes>
            <Route path="/" element={<Home/>}/>
            <Route path="/login" element={<Login/>}/>
        </Routes>
    </BrowserRouter>
  </div>
  );
}

export default App;
