import {PUBLIC_URL} from "../index";
import {Card} from "../api/CardModel";
import {Badge, Image, Modal} from "react-bootstrap";

export type CardDetailModalProps = {
    card: Card
    isShowing: boolean
    setShow: React.Dispatch<React.SetStateAction<boolean>>
}

function CardDetailModal(props: CardDetailModalProps) {
    const handleClose = () => props.setShow(false);

    const card = props.card
    return <Modal show={props.isShowing}
                  onHide={handleClose}
                  size={"xl"}
                  contentClassName={"border border-white"}>
        <Modal.Body className={"bg-dark rounded-md text-white "}>
            <div className="flex flex-row">
                <div>
                    {getCardAsBigImage(card)}
                </div>
                <div className={"flex-grow-1"}>
                    <div className="text-3xl mb-1 ml-2 capitalize">{card.name}</div>
                    <div className={"mb-3 ml-2"}>
                        {card.level !== 0 ? <Badge className={"mr-1"}>Level: {card.level}</Badge> : <></>}
                        <Badge className={"mr-1"}>{card.type}</Badge>
                        <Badge className={"mr-1"}>{card.race}</Badge>
                        {card.attribute ? <Badge className={"mr-1"}>{card.attribute}</Badge> : <></>}
                        {card.atk !== 0 ? <Badge className={"mr-1"} bg={"danger"}>{card.atk}</Badge> : <></>}
                        {card.def !== 0 ? <Badge className={"mr-1"} bg={"success"}>{card.def}</Badge> : <></>}
                    </div>
                    <p className={"p-2 mb-0 mt-2 ml-2 bg-gray-800"}>Description:</p>
                    <div className={"p-2 mb-1 ml-2 bg-gray-700"}>{card.desc}</div>
                    <p className={"p-2 mb-0 mt-2 ml-2 bg-gray-800"}>Included in the following Sets:</p>
                    <div className={"p-2 mb-1 ml-2 bg-gray-700"}>{card.sets}</div>
                </div>
            </div>
        </Modal.Body>
    </Modal>
}

function getCardAsBigImage(data: Card): JSX.Element {
    return <Image fluid={true}
                  src={`${PUBLIC_URL}/images/cards/` + data.id + "/big.png"}
                  alt="new"
    />
}

export default CardDetailModal