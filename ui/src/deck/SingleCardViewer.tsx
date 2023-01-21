import {PUBLIC_URL} from "../index";
import {Card} from "../api/CardModel";
import {Image} from "react-bootstrap";
import {useState} from "react";
import CardDetailModal from "./CardDetailModal";
import "./SingleCardViewer.css"

export type SingleCardViewerProps = {
    card: Card
    bottomElement ?: JSX.Element
    onlyImage ?: boolean
}

function SingleCardViewer(props: SingleCardViewerProps) {
    const [isShowingDetailView, setIsShowingDetailView] = useState(false);
    const handleShowDetailModal = () => setIsShowingDetailView(true);

    return <>
        <div className={"justify-center place-content-center flex flex-wrap"}>
            <div
                className={"hover:outline-none hover:ring hover:ring-sky-600 active:outline-none active:ring active:ring-sky-400"}
                onClick={handleShowDetailModal}>
                {getCardAsImage(props.card, props.onlyImage)}
            </div>
            {props.bottomElement ? props.bottomElement : <></>}
        </div>
        <CardDetailModal card={props.card} setShow={setIsShowingDetailView} isShowing={isShowingDetailView}/>
    </>
}

function getCardAsImage(data: Card, onlyImage: boolean | undefined): JSX.Element {
    if (onlyImage) {
        return <Image fluid={true}
                      className={"rounded-3"}
                      src={`${PUBLIC_URL}/images/cards/` + data.id + "/small.png"}
                      alt="new"
        />
    }

    return <Image fluid={true}
           src={`${PUBLIC_URL}/images/cards/` + data.id + "/small.png"}
           alt="new"
    />
}

export default SingleCardViewer