import SingleCardViewer from "../deck/SingleCardViewer";
import {Card, SortCards} from "../api/CardModel";
import {Button} from "react-bootstrap";
import React from "react";

export type MultiCardDraftAreaProps = {
    name: string
    draftRound: number
    maxRound: number
    cards: Card[]
    draftAction: (card: Card) => void
}

function MultiCardDraftArea(props: MultiCardDraftAreaProps) {
    let myInt = 50000
    let cards = SortCards(props.cards)

    let cardsViewBody = cards.map((card: Card) => {
            let draftButton = <Button className={"mt-1"}
                                      onClick={() => props.draftAction(card)}>Draft</Button>
            return <span key={myInt++}><SingleCardViewer card={card} bottomElement={draftButton}/></span>
        }
    );

    return <>
        <span className={"fw-bold font-monospace text-xl dark:text-white"}>{props.name}</span>
        <div>
            <span className={"mr-2 font-monospace fw-light dark:text-white"}>Round: {props.draftRound} / {props.maxRound}</span>
        </div>
        <div className={"p-2 grid grid-cols-10 gap-1 bg-ygo-card-viewer mt-2 mb-4"}>{cardsViewBody}</div>
    </>
}

export default MultiCardDraftArea