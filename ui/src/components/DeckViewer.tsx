import {Card, useRandomCard} from "./hooks/useRandomCard";
import CardViewer from "./CardViewer";

export type DeckViewerProps = {
    deck: string[]
}

function DeckViewer(props: DeckViewerProps) {

    let myRandomID = 50000
    let body = props.deck.map((number) =>
        <span key={myRandomID++}><CardViewer id={number}/></span>
    );

    return <><p className={"text-3xl"}>myDeck</p>
        <div className={"grid grid-cols-7 gap-4"}>{body}</div>
    </>
}

export default DeckViewer