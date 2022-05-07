import HeaderImage from '../images/header.jpg';
import {Link} from "react-router-dom";

var sectionStyle = {
    backgroundImage: `url(${HeaderImage})`,
    backgroundRepeat: "no-repeat",
    backgroundAttachment: "fixed",
    backgroundPosition: "center top"
}

function Jumbotron() {
    return <>
        <div className="container-fluid text-light bg-dark p-3 mt-2 mb-2">
        <div style={sectionStyle} className="container-fluid text-light p-0 ">
            <div className="container bg-dark bg-opacity-75 p-5">
                <h1 className="display-4">Welcome to YGO Draft</h1>
                <hr/>
                <p>Create your first randomized deck!</p>
                <Link to="/deck" className="btn btn-primary">Create Deck!</Link>
            </div>
            </div>
        </div>
    </>
}

export default Jumbotron