import {PUBLIC_URL} from "../index";
import {Card} from "../api/CardModel";
import {Image} from "react-bootstrap";

export type CardViewerProps = {
    card: Card
}

function CardViewer(props: CardViewerProps) {
    return <>{getCardAsImage(props.card)}</>
}

function getCardAsImage(data: Card): JSX.Element {
    return <Image fluid={true}
        src={`${PUBLIC_URL}/images/cards/` + data.id + "/small.png"}
        alt="new"
    />
}

export default CardViewer