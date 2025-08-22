import {useAuth} from "../core/auth/AuthProvider";
import YgoNavbar from "../core/YgoNavbar";
import YgoBackground from "../login/YgoBackground";
import React from "react";
import Home from "../home/Home";
import DeckRandomGeneratorPage from "../draft/shared/deck/DeckRandomGeneratorPage";
import LocaleDraftWizard from "../draft/local/LocaleDraftWizard";
import {createBrowserRouter, Route, Routes} from "react-router-dom";
import LoginPage from "../core/auth/LoginPage";
import {ProtectedRoute} from "./ProtectedRoute";
import UserPage from "../users/UserPage";
import AdminPage from "../users/AdminPage";
import {Navigate, RouterProvider} from "react-router";
import OnlineChallengePage from "../draft/online/challengePage/OnlineChallengePage";
import DraftOverviewPage from "../draft/online/draftPage/DraftOverviewPage";
import DraftMyRoundDeckPage from "../draft/online/roundPages/RoundPage";
import classNames from "classnames";
import Footer from "../core/Footer";

export const LoginPath = "/login"
export const AdminPath = "/admin"
export const HomePath = "/"
export const RandomDeckPath = "/randomdeck"
export const DraftDeckPath = "/draftdeck"
export const UserPath = "/user"
export const ChallengeUserPath = "/challenge"

export function DraftPagePath(draftID: string | number): string {
    return `/draft/${draftID}`
}

// 3️⃣ Router singleton created
const router = createBrowserRouter([
    {path: "*", element: <Root/>},
]);

function Root() {
    const {token} = useAuth();

    // Define public routes accessible to all users
    const routesForPublic: JSX.Element = <>
        <Route path={HomePath} element={withAll(<Home/>)}/>
    </>

    // Define routes accessible only to authenticated users
    const routesForAuthenticatedOnly: JSX.Element = <>
        <Route element={<ProtectedRoute/>}>
            <Route path={RandomDeckPath} element={withAll(<DeckRandomGeneratorPage/>)}/>
            <Route path={DraftDeckPath} element={withAll(<LocaleDraftWizard/>)}/>
            <Route path={UserPath} element={withAll(<UserPage/>)}/>
            <Route path={"/challenge"} element={withAll(<OnlineChallengePage/>)}/>
            <Route path={"/draft/:id"} element={withAll(<DraftOverviewPage/>)}/>
            <Route path={"/draft/:id/draftDeck"} element={withAll(<DraftMyRoundDeckPage/>)}/>
            <Route path={"/draft/:id/:roundID"} element={withAll(<DraftMyRoundDeckPage/>)}/>
        </Route>
    </>

    // Define routes accessible only to authenticated users
    const routesForAdminOnly: JSX.Element = <>
        <Route element={<ProtectedRoute/>}>
            <Route path={AdminPath} element={withAll(<AdminPage/>)}/>
        </Route>
    </>

    // Define routes accessible only to non-authenticated users
    const routesForNotAuthenticatedOnly: JSX.Element = <>
        <Route path={LoginPath} element={withBackground(withNavbar(<LoginPage/>))}/>
    </>

    // Provide the router configuration using RouterProvider
    return <Routes>
        {routesForPublic}
        {routesForAuthenticatedOnly}
        {!token ? routesForNotAuthenticatedOnly : <></>}
        {routesForAdminOnly}
        {<Route path="*" element={<Navigate to={HomePath} replace/>}/>}
    </Routes>
}

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
        <div className={classNames("container mx-auto bg-light dark:bg-dark", "border-t-2 border-border")}>
            <div className={"p-4"}>
                {element}
            </div>
            <Footer/>
        </div>
    </>
}

function withAll(element: JSX.Element) {
    return withBackground(withNavbar(withContainer(element)))
}

function AppRouter() {
    return <RouterProvider router={router}/>;
}

export default AppRouter;