import React from 'react';
import './App.css';
import YgoNavbar from "./core/YgoNavbar";
import {Route, Routes} from "react-router-dom";
import Home from "./home/Home";
import {Container} from "react-bootstrap";
import DeckGeneratorPage from "./deck/DeckGeneratorPage";

function App() {
    return (<>
            <YgoNavbar/>
            <Container>
                <Routes>
                    <Route path='/' element={<Home/>}/>
                    <Route path="/deck" element={<DeckGeneratorPage />}/>
                </Routes>
            </Container>
        </>
    );
}

export default App;
