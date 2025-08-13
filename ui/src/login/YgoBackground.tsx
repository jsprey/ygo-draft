import {Card,} from "../api/CardModel"
import SingleCardViewer from "../deck/SingleCardViewer";
import "./YgoBackground.css"
import {useRandomCards} from "../api/hooks/cards/useCards";
import {CardFilter} from "../api/CardFilter";
import {Alert, Spinner} from "react-bootstrap";
import React from "react";

function YgoBackground() {
    const {data, isLoading, error} = useRandomCards("login", 180, {} as CardFilter, {
        refetchOnWindowFocus: false
    })

    let content
    if (isLoading) {
        content = <Spinner animation="border" role="status">
            <span className="visually-hidden">Loading...</span>
        </Spinner>
    } else if (error) {
        content = <Alert variant={"danger"}>Failed to load background images!</Alert>
    } else if (data) {
        let myInt = 50000

        let cardsViewBody = data.cards.map((card: Card) =>
            <span key={myInt++}><SingleCardViewer className={"card"} card={card} onlyImage={true} readonly={true}/></span>
        );
        content = <div className={"blur-sm loginBackgroundContainer mySpecialBackground flex flex-wrap gap-3 bg-black -z-50"}>
            {cardsViewBody}
        </div>
    }

    return <>{content}</>
}

export default YgoBackground