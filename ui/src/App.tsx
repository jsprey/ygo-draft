import React from 'react';
import './App.css';
import YgoNavbar from "./core/YgoNavbar";
import {Route, Routes} from "react-router-dom";
import Home from "./home/Home";
import {Container} from "react-bootstrap";
import DeckRandomGeneratorPage from "./deck/DeckRandomGeneratorPage";
import DeckDraftWizard from "./draft/DeckDraftWizard";
import LoginPage from "./login/LoginPage";

function App() {
    return (<>
            <Routes>
                <Route path='/' element={<>
                    <Container>
                        <YgoNavbar/>
                        <Home/>
                    </Container>
                </>}/>
                <Route path='/login' element={<LoginPage/>}/>
                <Route path="/randomdeck" element={<>
                    <Container>
                        <YgoNavbar/>
                        <DeckRandomGeneratorPage/>
                    </Container>
                </>}/>
                <Route path="/draftdeck" element={<>
                    <Container>
                        <YgoNavbar/>
                        <DeckDraftWizard/>
                    </Container>
                </>}/>
            </Routes>
        </>
    );
}

export default App;
