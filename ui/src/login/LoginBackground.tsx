import {
    Card,
} from "../api/CardModel"
import SingleCardViewer from "../deck/SingleCardViewer";
import "./LoginBackground.css"

export type LoginBackgroundProps = {
    cards: Card[]
}

function LoginBackground(props: LoginBackgroundProps) {
    let myInt = 50000

    let cardsViewBody = props.cards.map((card: Card) =>
        <span key={myInt++}><SingleCardViewer card={card} onlyImage={true}/></span>
    );

    return <>
        <div className={"blur-sm loginBackgroundContainer mySpecialBackground p-2 grid gap-3 bg-black -z-50"}>
            {cardsViewBody}
        </div>
    </>
}

export default LoginBackground