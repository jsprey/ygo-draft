import SingleCardViewer from "./SingleCardViewer";
import {
    Card,
    FilterByType, SortCards,
} from "../api/CardModel"
import {CardType} from "../api/CardType";

export type MultiCardViewerProps = {
    name: string
    showDetails: boolean
    cards: Card[]
    singleCardElement ?: JSX.Element
}

function MultiCardViewer(props: MultiCardViewerProps) {
    let myInt = 50000
    let cards = SortCards(props.cards)

    const cardsTrapCardsCount = FilterByType(cards, [CardType.TrapCard]).length
    const cardsSpellCardsCount = FilterByType(cards, [CardType.SpellCard]).length
    const cardsMonsterCardsCount = cards.length - cardsTrapCardsCount - cardsSpellCardsCount
    let cardsViewBody = cards.map((card: Card) =>
        <span key={myInt++}><SingleCardViewer card={card}/></span>
    );

    return <>
        <span className={"fw-bold font-monospace text-xl dark:text-white"}>{props.name}</span>
        {props.showDetails ? <div>
            <span className={"mr-2 font-monospace fw-light dark:text-white"}>{cardsMonsterCardsCount} Monster Cards |</span>
            <span className={"mr-2 font-monospace fw-light dark:text-white"}>{cardsSpellCardsCount} Spell Cards |</span>
            <span className={"mr-2 font-monospace fw-light dark:text-white"}>{cardsTrapCardsCount} Trap Cards</span>
        </div>: <></>}
        <div className={"shadow-md p-2 grid grid-cols-10 gap-1 bg-zinc-700 mt-2 mb-4"}>{cardsViewBody}</div>
    </>
}

export default MultiCardViewer