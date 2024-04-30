import Jumbotron from "../core/Jumbotron";
import React from "react";
import classNames from "classnames";

function Home() {
    return <div>
        <Jumbotron/>
        <div className={classNames("p-3 mt rounded-2 shadow-md", "bg-light-1 dark:bg-dark-1", "text-dark dark:text-light")}>
            <div className={"text-lg font-bold"}>
                What is the draft mode exactly?
            </div>
            <li>
                The draft mode is a fun mode between at least two YGO
                players. Both players draft a deck round by round
                by selecting one of multiple cards each round. When both player reach the required 40 cards for their
                main decks they can download their decks and duel against each other.
            </li>

            <div className={"text-lg font-bold mt-2"}>
                What is the random mode exactly?
            </div>
            <li>
                The random mode is a fun mode between at least two YGO players. Both players generate a completely
                random deck and use it to duel against each other.
            </li>

            <div className={"text-lg font-bold mt-2 "}>
                How to Play?
            </div>
            <li>
                At first it is necessary to have a device capable of playing any YGO game that supports <i>.ydk</i> deck
                files.
            </li>
            <li>
                Then it is time to create a random deck or to draft a deck round by round.
            </li>
            <li>
                In the respective overviews of the deck generator or draft generator it is require to export the deck
                via the provided button. The current deck is downloaded as <i>.ydk</i> file.
            </li>
            <li>
                This file needs to be placed into the deck folder of your YGO game of choice.
            </li>
        </div>
    </div>
}

export default Home