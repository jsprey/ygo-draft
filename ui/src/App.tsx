import React from 'react';
import './App.css';
import YgoNavbar from "./core/YgoNavbar";
import {Route, Routes} from "react-router-dom";
import Home from "./home/Home";
import {Alert, Container, Spinner} from "react-bootstrap";
import DeckRandomGeneratorPage from "./deck/DeckRandomGeneratorPage";
import DeckDraftWizard from "./draft/DeckDraftWizard";
import LoginPage from "./login/LoginPage";
import {useRandomCards} from "./api/hooks/useCards";
import {CardFilter} from "./api/CardFilter";
import LoginBackground from "./login/LoginBackground";

function App() {
    const {data, isLoading, error} = useRandomCards("login", 90, {} as CardFilter, {
        refetchOnWindowFocus: false
    })

    let content
    if (isLoading) {
        content = <Spinner animation="border" role="status">
            <span className="visually-hidden">Loading...</span>
        </Spinner>
    } else if (error) {
        content = <Alert variant={"danger"}>Failed to load background images!</Alert>
    } else if (data) {
        content = <>
            <YgoNavbar/>
            <LoginBackground cards={data.cards}/>
        </>
    }

    return (<>
            <Routes>
                <Route path='/' element={<>
                    {content}
                    <Container className={"bg-light"}>
                        <Home/>
                    </Container>
                </>}/>
                <Route path='/login' element={<>
                    {content}
                    <Container>
                        <LoginPage/>
                    </Container>
                </>}/>
                <Route path="/randomdeck" element={<>
                    {content}
                    <Container className={"bg-light"}>
                        <DeckRandomGeneratorPage/>
                    </Container>
                </>}/>
                <Route path="/draftdeck" element={<>
                    {content}
                    <Container className={"bg-light"}>
                        <DeckDraftWizard/>
                    </Container>
                </>}/>
            </Routes>
        </>
    );
}

export default App;
