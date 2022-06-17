import {Container, Nav, Navbar} from "react-bootstrap";
import {Link} from "react-router-dom";

function YgoNavbar() {
    // noinspection TypeScriptValidateTypes
    return <>
        <Navbar expand="lg" bg="dark" variant="dark">
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
                    <Nav.Link as={Link} to="/randomdeck">Mode: Random</Nav.Link>
                    <Nav.Link as={Link} to="/draftdeck">Mode: Draft</Nav.Link>
                </Nav>
                <Nav>
                    <Nav.Link as={Link} to="/test1">Placeholder</Nav.Link>
                    <Nav.Link as={Link} to="/login">Login</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    </>
}

export default YgoNavbar