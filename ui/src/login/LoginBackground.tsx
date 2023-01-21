import {
    Card,
    SortCards,
} from "../api/CardModel"
import SingleCardViewer from "../deck/SingleCardViewer";
import "./LoginBackground.css"

export type LoginBackgroundProps = {
    cards: Card[]
}

function LoginBackground(props: LoginBackgroundProps) {
    let myInt = 50000
    let cards = SortCards(props.cards)

    let cardsViewBody = cards.map((card: Card) =>
        <span key={myInt++}><SingleCardViewer card={card} onlyImage={true}/></span>
    );

    return <>
        {/*todo use this /*<div className={"mySpecialBackground p-2 grid grid-cols-10 gap-1 bg-dark"}>{cardsViewBody}</div>*/}
        <div className={"test mySpecialBackground w-100 h-100 p-2 grid gap-3 bg-black"}>
            {cardsViewBody}
        </div>
    </>
}

export default LoginBackground