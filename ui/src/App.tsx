import React from 'react';
import './App.css';
import YgoNavbar from "./core/YgoNavbar";
import {Route, Routes} from "react-router-dom";
import Home from "./home/Home";
import {Container} from "react-bootstrap";
import DeckRandomGeneratorPage from "./deck/DeckRandomGeneratorPage";
import DeckDraftGeneratorPage from "./deck/DeckDraftGeneratorPage";

function App() {
    return (<>
            <YgoNavbar/>
            <Container>
                <Routes>
                    <Route path='/' element={<Home/>}/>
                    <Route path="/randomdeck" element={<DeckRandomGeneratorPage />}/>
                    <Route path="/draftdeck" element={<DeckDraftGeneratorPage />}/>
                </Routes>
            </Container>
        </>
    );
}

export default App;
