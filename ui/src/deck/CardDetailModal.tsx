import {PUBLIC_URL} from "../index";
import {Card} from "../api/CardModel";
import Modal from "../core/Modal";
import classNames from "classnames";
import {ReactNode} from "react"

export type CardDetailModalProps = {
    card: Card
    isShowing: boolean
    setShow: React.Dispatch<React.SetStateAction<boolean>>
}

function CardDetailModal(props: CardDetailModalProps) {
    const handleClose = () => props.setShow(false);

    const badgeCN = classNames("mr-1 p-2 rounded-lg bg-secondary-light dark:bg-dark-1")

    const card = props.card
    return <Modal show={props.isShowing}
                  setShow={props.setShow}
                  onHide={handleClose}>
        <div className={"flex"}>
            <div className={"w-[28%]"}>
                {getCardAsBigImage(card)}
            </div>
            <div className={"w-[72%]"}>
                <div className="font-bold text-3xl mt-2 mb-2 ml-2 capitalize">{card.name}</div>
                <div className={"flex mb-2 ml-2"}>
                    {card.level !== 0 ? <div className={classNames("flex", badgeCN)}>{getStars(props.card.level)}</div> : null}
                    <div className={classNames(badgeCN)}>{card.type}</div>
                    <div className={classNames(badgeCN)}>{card.race}</div>
                    {card.attribute ? <div className={classNames(badgeCN)}>{card.attribute}</div> : null}
                    {card.atk !== 0 ? <div className={classNames(badgeCN, "!text-dark !bg-danger-light")}>ATK: {card.atk}</div> : null}
                    {card.def !== 0 ? <div className={classNames(badgeCN, "!text-dark !bg-success-light")}>DEF: {card.def}</div> : null}
                </div>
                <p className={"p-2 mb-0 mt-2 ml-2 font-semibold bg-light-3 dark:bg-dark-3"}>Description:</p>
                <div className={"p-2 mb-1 ml-2 bg-light-1 dark:bg-dark-1"}>{card.desc}</div>
                <p className={"p-2 mb-0 mt-2 ml-2 font-semibold bg-light-3 dark:bg-dark-3"}>Included in
                    the following Sets:</p>
                <div className={"p-2 mb-1 ml-2 bg-light-1 dark:bg-dark-1 dark:text-white"}>
                    {card.sets.length === 0 ? <div key={"empty"}>This card is not available in any set.</div> :
                        <SetList card={card}/>}
                </div>
            </div>
        </div>
    </Modal>
}

function getCardAsBigImage(data: Card): JSX.Element {
    return <img
        src={`${PUBLIC_URL}/images/cards/` + data.id + "/big.png"}
        alt="card"
        className={"rounded-lg"}
    />
}

function getStars(numberOfStars:number) : ReactNode {
    const starsList: ReactNode[] = []
    for (let i = 0; i < numberOfStars; i++) {
        starsList.push(<img
            key={`level-star-${i}`}
            src={"/images/card_star.png"}
            alt={"level star of card"}
            className={"w-[20px]"}
        />)
    }

    return <div className={"flex items-center"}>
        {starsList}
    </div>
}

type NewlineTextProps = {
    card: Card
}

function SetList(props: NewlineTextProps) {
    const text = props.card.sets;
    return <div className={"flex flex-col"}>
        {text.split(',').map(str => <div key={str}>● {str}</div>)}
    </div>
}

export default CardDetailModal