import SingleCardViewer from "../../deck/SingleCardViewer";
import {Card, SortCards} from "../../api/CardModel";
import React from "react";
import Button from "../../core/Button";
import classNames from "classnames";

export type MultiCardDraftAreaProps = {
    name: string
    draftRound: number
    maxRound: number
    cards: Card[]
    draftAction: (card: Card) => void
}

function MultiCardDraftArea(props: MultiCardDraftAreaProps) {
    let cards = SortCards(props.cards)

    let cardsViewBody = cards.map((card: Card, index: number) => {
            let draftButton = <Button variant={"primary"} className={"mt-2"}
                                      onClick={() => props.draftAction(card)}>Draft</Button>
            return <span key={`card-viewer-card-${card.id}-${index}`}><SingleCardViewer card={card}
                                                                               bottomElement={draftButton}/></span>
        }
    );

    return <div className={"text-dark dark:text-light"}>
        <div
            className={classNames("flex-grow-1 flex items-center p-2 justify-between", "bg-light-1 dark:bg-dark-1", "rounded-tl rounded-tr", "border border-dark-2 dark:border-light-2")}>
            <span className={classNames("p-1", "font-bold text-2xl")}>
               {props.name}
           </span>
            <span className={classNames("p-1", "font-bold text-2xl")}>
               Round: {props.draftRound} / {props.maxRound}
           </span>
        </div>
        <div className={classNames("p-2 grid grid-cols-10 gap-1", "bg-dark-3", "border-l border-b border-r border-dark-2 dark:border-light-2")}>{cardsViewBody}</div>
    </div>
}

export default MultiCardDraftArea