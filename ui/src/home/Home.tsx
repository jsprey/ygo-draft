import Jumbotron from "../core/Jumbotron";
import React from "react";

function Home() {
    return <>
        <Jumbotron/>
        <div className={"p-3 mt-0 bg-blue-200 rounded-2 shadow-md"}>
            <div className={"text-lg fw-bold"}>What is the draft mode exactly?</div>
            <li>The draft mode is a fun mode between at least two YGO players. Both players draft a deck round by round
                by selecting one of multiple cards each round. When both player reach the required 40 cards for their
                main decks they can download their decks and duel against each other.
            </li>
            <div className={"text-lg fw-bold mt-2"}>What is the random mode exactly?</div>
            <li>The random mode is a fun mode between at least two YGO players. Both players generate a completely
                random deck and use it to duel against each other.
            </li>
            <div className={"text-lg fw-bold mt-2"}>How to Play?</div>
            <li>At first it is necessary to have a device capable of playing any YGO game that supports <i>.ydk</i> deck
                files.
            </li>
            <li>Then it is time to create a random deck or to draft a deck round by round.</li>
            <li>In the respective overviews of the deck generator or draft generator it is require to export the deck
                via the provided button. The current deck is downloaded as <i>.ydk</i> file.
            </li>
            <li>This file needs to be placed into the deck folder of your YGO game of choice.</li>
        </div>
    </>
}

export default Home