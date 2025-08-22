import {PUBLIC_URL} from "../../../index";
import {Card} from "../../../api/CardModel";
import {useState} from "react";
import CardDetailModal from "./CardDetailModal";
import "./SingleCardViewer.css"

export type SingleCardViewerProps = {
    card: Card
    bottomElement ?: JSX.Element
    onlyImage ?: boolean
    readonly ?: boolean
    className ?: string
}

type CardViewDivProps = {
    className: string
    onClick?: () => void
}

function SingleCardViewer(props: SingleCardViewerProps) {
    const [isShowingDetailView, setIsShowingDetailView] = useState(false);
    const handleShowDetailModal = () => setIsShowingDetailView(true);

    const cardViewProps: CardViewDivProps = {
        className: ""
    }

    if (!props.readonly) {
        cardViewProps.className = "hover:ring hover:ring-primary active:ring active:ring-primary-active"
        cardViewProps.onClick = handleShowDetailModal
    }

    return <>
        <div className={`justify-center place-content-center flex flex-wrap ${props.className}`}>
            <div {...cardViewProps}>
                {getCardAsImage(props.card, props.onlyImage)}
            </div>
            {props.bottomElement ? props.bottomElement : <></>}
        </div>
        <CardDetailModal card={props.card} setShow={setIsShowingDetailView} isShowing={isShowingDetailView}/>
    </>
}

function getCardAsImage(data: Card, onlyImage: boolean | undefined): JSX.Element {
    if (onlyImage) {
        return <img
                      className={"rounded-3"}
                      src={`${PUBLIC_URL}/images/cards/` + data.id + "/small.png"}
                      alt="new"
                      draggable={false}
        />
    }

    return <img src={`${PUBLIC_URL}/images/cards/` + data.id + "/small.png"}
           alt="new"
    />
}

export default SingleCardViewer