import HeaderImage from '../images/header.jpg';
import {useNavigate} from "react-router";
import Button from "./Button";
import {DraftDeckPath} from "../routes/AppRouter";

const sectionStyle = {
    backgroundImage: `url(${HeaderImage})`,
    backgroundRepeat: "no-repeat",
    backgroundAttachment: "fixed",
    backgroundPosition: "center top",
    backgroundSize: 'cover',
}

function Jumbotron() {
    var navigateFunction = useNavigate();

    return <>
        <div className="mb-2 text-light">
            <div style={sectionStyle} className="p-0">
                <div className="container bg-dark bg-opacity-75 px-10 py-20">
                    <p className={"text-3xl font-bold"}>Welcome to YGO Draft</p>
                    <hr/>
                    <p className={"text-light-3"}>Create your first randomized deck!</p>
                    <Button className={"mt-2"} variant={"primary"} onClick={() => {navigateFunction(DraftDeckPath)}}>Create Deck!</Button>
                </div>
            </div>
        </div>
    </>
}

export default Jumbotron