import {PUBLIC_URL} from "../index";
import {Card} from "../api/CardModel";
import {Image} from "react-bootstrap";
import {useState} from "react";
import CardDetailModal from "./CardDetailModal";

export type CardViewerProps = {
    card: Card
}

function CardViewer(props: CardViewerProps) {
    const [isShowingDetailView, setIsShowingDetailView] = useState(false);
    const handleShowDetailModal = () => setIsShowingDetailView(true);

    return <>
        {/*<div className={"p-1 hover:bg-sky-600 active:bg-sky-400"} onClick={handleShowDetailModal}>*/}
        <div className={"hover:outline-none hover:ring hover:ring-sky-600 active:outline-none active:ring active:ring-sky-400"} onClick={handleShowDetailModal}>
            {getCardAsImage(props.card)}
        </div>
        <CardDetailModal card={props.card} setShow={setIsShowingDetailView} isShowing={isShowingDetailView}/>
    </>
}

function getCardAsImage(data: Card): JSX.Element {
    return <Image fluid={true}
                  src={`${PUBLIC_URL}/images/cards/` + data.id + "/small.png"}
                  alt="new"
    />
}

export default CardViewer