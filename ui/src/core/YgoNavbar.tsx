import {Container, Nav, Navbar} from "react-bootstrap";
import {Link} from "react-router-dom";
import {useAuth} from "../auth/AuthProvider";
import {useNavigate} from "react-router";
import UserNavbarBadge from "../auth/UserNavbarBadge";
import React from "react";
import ThemeSwitcher from "./ThemeSwitcher";
import {useTheme} from "./context/ColorThemeProvider";

function YgoNavbar() {
    const navigation = useNavigate();
    const {token, setToken} = useAuth();
    const {isDarkMode} = useTheme();

    const logout = () => {
        setToken(null)
        navigation("/login")
    }

    let userInformation = <></>
    if (token) {
        userInformation = <div className={"flex"}>
            <UserNavbarBadge/>
            <Navbar.Collapse className="ml-3 justify-content-end">
                <Nav.Link onClick={logout}>Logout</Nav.Link>
            </Navbar.Collapse>
        </div>
    }

    // noinspection TypeScriptValidateTypes
    return <>
        <Navbar expand="lg" className={"border-bottom border-body"} bg={isDarkMode ? "dark" : "light"} data-bs-theme={isDarkMode ? "dark" : "light"}>
            <Container>
                <Navbar.Brand as={Link} to="/">
                    <img
                        alt=""
                        src="/logo.png"
                        width="30"
                        height="30"
                        className="d-inline-block align-top"
                    />{' '}
                    YgoDraft
                </Navbar.Brand>
                <Nav className="me-auto">
                    <Nav.Link as={Link} to="/">Home</Nav.Link>
                    {token ? <Nav.Link as={Link} to="/randomdeck">Mode: Random</Nav.Link> : <></>}
                    {token ? <Nav.Link as={Link} to="/draftdeck">Mode: Draft</Nav.Link> : <></>}
                </Nav>
                <Nav>
                    <ThemeSwitcher/>
                    {!token ? <Nav.Link as={Link} to="/login">Login</Nav.Link> : <></>}
                    {token ?  userInformation : <></>}
                </Nav>
            </Container>
        </Navbar>
    </>
}

export default YgoNavbar