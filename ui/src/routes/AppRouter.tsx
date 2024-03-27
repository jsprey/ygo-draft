import {useAuth} from "../auth/AuthProvider";
import YgoNavbar from "../core/YgoNavbar";
import YgoBackground from "../login/YgoBackground";
import React from "react";
import {Container} from "react-bootstrap";
import Home from "../home/Home";
import DeckRandomGeneratorPage from "../deck/DeckRandomGeneratorPage";
import DeckDraftWizard from "../draft-local/DeckDraftWizard";
import {Route, Routes} from "react-router-dom";
import LoginPage from "../auth/LoginPage";
import {ProtectedRoute} from "./ProtectedRoute";
import {useTheme} from "../core/context/ColorThemeProvider";
import UserPage from "../users/UserPage";
import AdminPage from "../users/AdminPage";
import {Navigate} from "react-router";
import ChallengeDraftPage from "../draft-challenge/ChallengeDraftPage";

const AppRouter = () => {
    const {token} = useAuth();
    var {isDarkMode} = useTheme();

    // Define public routes accessible to all users
    const routesForPublic: JSX.Element = <>
        <Route path={"/"} element={withBackground(withNavbar(withContainer(<Home/>, isDarkMode)))}/>
    </>

    // Define routes accessible only to authenticated users
    const routesForAuthenticatedOnly: JSX.Element = <>
        <Route element={<ProtectedRoute/>}>
            <Route path={"/randomdeck"} element={withBackground(withNavbar(withContainer(<DeckRandomGeneratorPage/>, isDarkMode)))}/>
            <Route path={"/draftdeck"} element={withBackground(withNavbar(withContainer(<DeckDraftWizard/>, isDarkMode)))}/>
            <Route path={"/user"} element={withBackground(withNavbar(withContainer(<UserPage/>, isDarkMode)))}/>
            <Route path={"/challenge"} element={withBackground(withNavbar(withContainer(<ChallengeDraftPage/>, isDarkMode)))}/>
        </Route>
    </>

    // Define routes accessible only to authenticated users
    const routesForAdminOnly: JSX.Element = <>
        <Route element={<ProtectedRoute/>}>
            <Route path={"/admin"} element={withBackground(withNavbar(withContainer(<AdminPage/>, isDarkMode)))}/>
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
        {routesForAdminOnly}
        {<Route path="*" element={<Navigate to="/" replace />} />}
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

function withContainer(element: JSX.Element, isDarkMode: boolean) {
    return <>
        <Container className={isDarkMode ? "bg-dark" : "bg-light"}>
            {element}
        </Container>
    </>
}

export default AppRouter;