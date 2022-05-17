import React from 'react';
import './App.css';
import YgoNavbar from "./core/YgoNavbar";
import {Route, Routes} from "react-router-dom";
import Home from "./home/Home";
import {Container} from "react-bootstrap";
import DeckRandomGeneratorPage from "./deck/DeckRandomGeneratorPage";
import DeckDraftWizard from "./draft/DeckDraftWizard";

function App() {
    return (<>
            <YgoNavbar/>
            <Container>
                <Routes>
                    <Route path='/' element={<Home/>}/>
                    <Route path="/randomdeck" element={<DeckRandomGeneratorPage />}/>
                    <Route path="/draftdeck" element={<DeckDraftWizard />}/>
                </Routes>
            </Container>
        </>
    );
}

export default App;
