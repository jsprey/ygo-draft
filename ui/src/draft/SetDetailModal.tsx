import {Alert, Modal, Spinner} from "react-bootstrap";
import {useSetCards} from "../api/hooks/useSets";
import React from "react";
import MultiCardViewer from "../deck/MultiCardViewer";

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
        content = <Spinner animation="border" role="status">
            <span className="visually-hidden">Loading...</span>
        </Spinner>
    } else if (error) {
        content = <Alert variant={"danger"}>Failed to load cards from set!</Alert>
    } else if (data) {
        content = <MultiCardViewer name={data.set.set_name} showDetails={false} cards={data.cards}/>
    }

    return <Modal show={props.isShowing}
                  onHide={handleClose}
                  size={"xl"}
                  contentClassName={"dark:text-white"}>
        <Modal.Body className={"bg-ygo-light dark:bg-ygo-dark dark:text-white"}>
            {content}
        </Modal.Body>
    </Modal>
}

export default SetDetailModal