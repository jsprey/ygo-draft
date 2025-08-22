import SingleCardViewer from "./SingleCardViewer";
import {Card, FilterByType, SortCards,} from "../../../api/CardModel"
import {CardType} from "../../../api/CardType";
import classNames from "classnames";

export type MultiCardViewerProps = {
    name: string
    showDetails: boolean
    cards: Card[]
    singleCardElement?: JSX.Element
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

    const normalEffectMonsterBadge = classNames("rounded-lg p-1", "text-light border-2 border-border", "bg-cardcolors-effect")
    const spellCardBadge = classNames("rounded-lg p-1", "text-light border-2 border-border", "bg-cardcolors-spell")
    const trapCardsBadge = classNames("rounded-lg p-1", "text-light border-2 border-border", "bg-cardcolors-trap")

    return <div className={classNames("flex flex-col dark:text-light")}>
        <div
            className={classNames("flex-grow-1 flex items-center p-2", "bg-light-1 dark:bg-dark-1", "rounded-tl rounded-tr", "border border-border")}>
           <span className={classNames("p-1 mr-2", "font-bold text-2xl")}>
               {props.name}
           </span>
            {props.showDetails ? <div>
            <span className={classNames(normalEffectMonsterBadge, "mr-2 dark:text-white")}>
                {cardsMonsterCardsCount} Monster Cards
            </span>
                <span className={classNames(spellCardBadge, "mr-2 dark:text-white")}>
                {cardsSpellCardsCount} Spell Cards
            </span>
                <span className={classNames(trapCardsBadge, "mr-2 dark:text-white")}>
                {cardsTrapCardsCount} Trap Cards
            </span>
            </div> : <></>}
        </div>
        <div
            className={classNames("p-2 grid grid-cols-10 gap-1", "bg-dark-3", "border-l border-b border-r border-border")}>
            {cardsViewBody}
        </div>
    </div>
}

export default MultiCardViewer