import {Card, useRandomCard} from "./hooks/useRandomCard";
import CardViewer from "./CardViewer";
import {Deck} from "./hooks/useRandomCards";

export type DeckViewerProps = {
    deck: Deck
}

function DeckViewer(props: DeckViewerProps) {

    let myRandomID = 50000
    let body = props.deck.cards.map((number) =>
        <span key={myRandomID++}><CardViewer id={number}/></span>
    );

    return <><p className={"text-3xl"}>myDeck</p>
        <div className={"grid grid-cols-7 gap-4"}>{body}</div>
    </>
}

export default DeckViewer