import {useAuth} from "../auth/AuthProvider";
import YgoNavbar from "../core/YgoNavbar";
import YgoBackground from "../login/YgoBackground";
import React from "react";
import {Container} from "react-bootstrap";
import Home from "../home/Home";
import DeckRandomGeneratorPage from "../deck/DeckRandomGeneratorPage";
import DeckDraftWizard from "../draft/DeckDraftWizard";
import {Route, Routes} from "react-router-dom";
import LoginPage from "../login/LoginPage";
import {ProtectedRoute} from "./ProtectedRoute";

const AppRouter = () => {
    const {token} = useAuth();

    // Define public routes accessible to all users
    const routesForPublic: JSX.Element = <>
        <Route path={"/"} element={withBackground(withNavbar(withContainer(<Home/>)))}/>
    </>

    // Define routes accessible only to authenticated users
    const routesForAuthenticatedOnly: JSX.Element = <>
        <Route element={<ProtectedRoute/>}>
            <Route path={"/randomdeck"} element={withBackground(withNavbar(withContainer(<DeckRandomGeneratorPage/>)))}/>
            <Route path={"/draftdeck"} element={withBackground(withNavbar(withContainer(<DeckDraftWizard/>)))}/>
        </Route>
    </>

    // Define routes accessible only to non-authenticated users
    const routesForNotAuthenticatedOnly: JSX.Element = <>
        <Route path={"/login"} element={withBackground(withNavbar(<LoginPage/>))}/>
    </>

    // Provide the router configuration using RouterProvider
    return <Routes>
        {routesForPublic}
        {routesForAuthenticatedOnly}
        {!token ? routesForNotAuthenticatedOnly : <></>}
    </Routes>
};

function withNavbar(element: JSX.Element) {
    return <>
        <YgoNavbar/>
        {element}
    </>
}

function withBackground(element: JSX.Element) {
    return <>
        <YgoBackground/>
        {element}
    </>
}

function withContainer(element: JSX.Element) {
    return <>
        <Container className={"bg-dark"}>
            {element}
        </Container>
    </>
}

export default AppRouter;