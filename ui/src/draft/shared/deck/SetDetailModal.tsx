import {useSetCards} from "../../../api/hooks/cards/useSets";
import React from "react";
import MultiCardViewer from "./MultiCardViewer";
import Spinner from "../../../core/Spinner";
import Alert from "../../../core/Alert";
import Modal from "../../../core/Modal";

export type SetDetailModalProps = {
    setCode: string
    isShowing: boolean
    setShow: React.Dispatch<React.SetStateAction<boolean>>
}

function SetDetailModal(props: SetDetailModalProps) {
    const {data, isLoading, error} = useSetCards(props.setCode)
    const handleClose = () => props.setShow(false);

    let content
    if (isLoading) {
        content = <Spinner/>
    } else if (error) {
        content = <Alert variant={'danger'}>Failed to load cards from set!</Alert>
    } else if (data) {
        content = <MultiCardViewer name={data.set.set_name} showDetails={false} cards={data.cards}/>
    }

    return <Modal show={props.isShowing}
                  setShow={props.setShow}
                         onHide={handleClose}>
        <div className={"bg-ygo-light dark:bg-ygo-dark dark:text-white"}>
            {content}
        </div>
    </Modal>
}

export default SetDetailModal