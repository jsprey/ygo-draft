import {Container, Nav, Navbar} from "react-bootstrap";
import {Link} from "react-router-dom";
import {useAuth} from "../auth/AuthProvider";
import {useNavigate} from "react-router";
import UserNavbarBadge from "../users/UserNavbarBadge";
import React from "react";
import ThemeSwitcher from "./ThemeSwitcher";
import {useTheme} from "./context/ColorThemeProvider";
import {useCurrentUser} from "../api/hooks/users/useUser";
import SvgIconButton from "./SvgIconButton";

const CogWheelIcon = <SvgIconButton size={18} classNames={"fill-neutral-600 dark:fill-neutral-50"}>
    <path
        d="M9.405 1.05c-.413-1.4-2.397-1.4-2.81 0l-.1.34a1.464 1.464 0 0 1-2.105.872l-.31-.17c-1.283-.698-2.686.705-1.987 1.987l.169.311c.446.82.023 1.841-.872 2.105l-.34.1c-1.4.413-1.4 2.397 0 2.81l.34.1a1.464 1.464 0 0 1 .872 2.105l-.17.31c-.698 1.283.705 2.686 1.987 1.987l.311-.169a1.464 1.464 0 0 1 2.105.872l.1.34c.413 1.4 2.397 1.4 2.81 0l.1-.34a1.464 1.464 0 0 1 2.105-.872l.31.17c1.283.698 2.686-.705 1.987-1.987l-.169-.311a1.464 1.464 0 0 1 .872-2.105l.34-.1c1.4-.413 1.4-2.397 0-2.81l-.34-.1a1.464 1.464 0 0 1-.872-2.105l.17-.31c.698-1.283-.705-2.686-1.987-1.987l-.311.169a1.464 1.464 0 0 1-2.105-.872zM8 10.93a2.929 2.929 0 1 1 0-5.86 2.929 2.929 0 0 1 0 5.858z"/>
</SvgIconButton>

function YgoNavbar() {
    const navigation = useNavigate();
    const {token, setToken} = useAuth();
    const user = useCurrentUser()
    const {isDarkMode} = useTheme();

    const logout = () => {
        setToken(null)
        navigation("/login")
    }

    let userInformation = <></>
    if (token) {
        userInformation = <div className={"flex"}>
            <UserNavbarBadge/>
            <Navbar.Collapse className="justify-content-end">
                <Nav.Link onClick={logout}>Logout</Nav.Link>
            </Navbar.Collapse>
        </div>
    }

    // noinspection TypeScriptValidateTypes
    return <>
        <Navbar expand="lg" className={"border-bottom border-body navBarContainer"} bg={isDarkMode ? "dark" : "light"}
                data-bs-theme={isDarkMode ? "dark" : "light"}>
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
                    {user && user.data?.is_admin ? <Nav.Link as={Link} to="/admin" className={"align-self-center"}>{CogWheelIcon}</Nav.Link> : <></>}
                    {!token ? <Nav.Link as={Link} to="/login">Login</Nav.Link> : <></>}
                    {token ? userInformation : <></>}
                </Nav>
            </Container>
        </Navbar>
    </>
}

export default YgoNavbar