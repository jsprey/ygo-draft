import {Container} from "react-bootstrap";
import Jumbotron from "../core/Jumbotron";

function Home() {
    return <>
        <Jumbotron />
        <div className={"text-white"}>
            <h1 className={"pt-5 "}>What is YGO-Draft?</h1>
            <p> This is the description of YGO Draft.</p>
        </div>
    </>
}

export default Home